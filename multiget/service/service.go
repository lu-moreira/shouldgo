package service

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/lu-moreira/shouldgo/multiget/dto"
	"github.com/lu-moreira/shouldgo/multiget/model"
)

type getAllFunc func(ctx context.Context, request dto.GetUserRequest) (interface{}, error)

type RetrieverResult struct {
	P *dto.GetUserAssignedPermissionsResponse `json:"permissions"`
	U *dto.GetUserResponse                    `json:"user"`
	A *dto.GetUserAtrributesResponse          `json:"attributes"`
}

func GetUserInformation(ctx context.Context, request dto.GetUserRequest) (*RetrieverResult, error) {
	fatalErrors := make(chan error)
	results := make(chan interface{})
	done := make(chan bool)
	defer func() {
		close(done)
		close(results)
		close(fatalErrors)
	}()

	getAll := []getAllFunc{
		func(ctx context.Context, request dto.GetUserRequest) (interface{}, error) {
			time.Sleep(21000 * time.Millisecond)
			permisions := dto.GetUserAssignedPermissionsResponse([]model.UserAssignedPermission{{ID: 100}})
			return &permisions, nil
		},
		func(ctx context.Context, request dto.GetUserRequest) (interface{}, error) {
			time.Sleep(5000 * time.Millisecond)
			user := dto.GetUserResponse(model.User{ID: 102})
			return &user, nil
		},
		func(ctx context.Context, request dto.GetUserRequest) (interface{}, error) {
			time.Sleep(2000 * time.Millisecond)
			attr := dto.GetUserAtrributesResponse([]model.Attribute{{ID: 101}})
			return &attr, nil
		},
	}

	wg := &sync.WaitGroup{}

	for _, getFn := range getAll {
		wg.Add(1)
		go func(fn getAllFunc) {
			defer wg.Done()
			res, err := fn(ctx, request)
			if err != nil {
				fatalErrors <- err
				return
			}
			results <- res
		}(getFn)
	}

	go func() {
		wg.Wait()
		done <- true
	}()

	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())

	rr := &RetrieverResult{}
	for {
		select {
		case fatal := <-fatalErrors:
			return nil, fatal
		case r := <-results:
			switch res := r.(type) {
			case *dto.GetUserAssignedPermissionsResponse:
				rr.P = res
			case *dto.GetUserResponse:
				rr.U = res
			case *dto.GetUserAtrributesResponse:
				rr.A = res
			}
		case <-done:
			return rr, nil
		}
	}
}
