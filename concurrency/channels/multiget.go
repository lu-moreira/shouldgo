package channels

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo
// Realizar diferente pedidos e montar um result unico

func ProduceConsumeOrders() {
	orders := make(chan int)

	// send messages to channel
	makeOrder := func(id int, order chan<- int) {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Preparing order #", id)
		order <- id
	}

	totalOrders := []int{1, 2, 3, 4, 5}

	wg := &sync.WaitGroup{}

	// receiving messages from channel
	for _, v := range totalOrders { // blocking
		wg.Add(1)
		go func(orderID int) {
			defer wg.Done()
			makeOrder(orderID, orders)
		}(v)
	}

	go func() {
		wg.Wait() // blocking
		close(orders)
	}()

	for order := range orders { // blocking
		fmt.Println("Prepared order #", order)
	}
}
