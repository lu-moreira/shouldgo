package myctx

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

type Ctx interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type emptyCtx int

func (emptyCtx) Deadline() (deadline time.Time, ok bool) { return }
func (emptyCtx) Done() <-chan struct{}                   { return nil }
func (emptyCtx) Err() error                              { return nil }
func (emptyCtx) Value(key interface{}) interface{}       { return nil }

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Ctx {
	return background
}

func TODO() Ctx {
	return todo
}

type cancelCtx struct {
	Ctx
	done chan struct{}
	err  error
	m    sync.Mutex
}

func (ctx *cancelCtx) Done() <-chan struct{} { return ctx.done }
func (ctx *cancelCtx) Err() error {
	ctx.m.Lock()
	defer ctx.m.Unlock()
	return ctx.err
}

var Canceled = errors.New("context canceled")
var DeadlineExceeded = timeouterErr{} // errors.New("deadline exceeded")

type CancelFunc func()

func WithCancel(parent Ctx) (Ctx, CancelFunc) {
	ctx := &cancelCtx{
		Ctx:  parent,
		done: make(chan struct{}),
	}

	cancel := func() {
		ctx.cancelWithErr(Canceled)
	}

	go func() {
		select {
		case <-parent.Done():
			ctx.cancelWithErr(parent.Err())
		case <-ctx.Done():
		}
	}()
	return ctx, cancel
}

func (ctx *cancelCtx) cancelWithErr(err error) bool {
	ctx.m.Lock()
	defer ctx.m.Unlock()
	if ctx.err != nil {
		return true
	}
	ctx.err = err
	close(ctx.done)
	return false
}

type deadlineCtx struct {
	*cancelCtx
	deadline time.Time
}

func (ctx *deadlineCtx) Timeout() bool {
	return time.Now().After(ctx.deadline)
}

func (ctx *deadlineCtx) Deadline() (deadline time.Time, ok bool) { return ctx.deadline, true }

func WithDeadline(parent Ctx, deadline time.Time) (Ctx, CancelFunc) {
	cctx, cancel := WithCancel(parent)
	ctx := &deadlineCtx{
		cancelCtx: cctx.(*cancelCtx),
		deadline:  deadline,
	}

	t := time.AfterFunc(time.Until(deadline), func() {
		ctx.cancelWithErr(DeadlineExceeded)
	})

	stop := func() {
		t.Stop()
		cancel()
	}

	return ctx, stop
}

func WithTimeout(parent Ctx, timeout time.Duration) (Ctx, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

type valueCtx struct {
	Ctx
	value, key interface{}
}

func (ctx *valueCtx) Value(key interface{}) interface{} {
	if key == ctx.key {
		return ctx.value
	}
	return ctx.Ctx.Value(key)
}

func WithValue(parent Ctx, key, value interface{}) Ctx {
	if key == nil {
		panic("key is nil")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}

	return &valueCtx{
		Ctx:   parent,
		key:   key,
		value: value,
	}
}

type timeouterErr struct{}

func (timeouterErr) Error() string {
	return "exceeded"
}
func (timeouterErr) Timeout() bool {
	return true
}
