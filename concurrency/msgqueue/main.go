package main

import (
	"fmt"
	"time"

	"github.com/lu-moreira/shouldgo/concurrency/msgqueue/worker"
)

func PrintPayload(payload map[string]string) {
	fmt.Println(payload)
}

func main() {
	defer worker.Wait()

	for i := 0; i < 10; i++ {
		j := worker.Job{
			Action: PrintPayload,
			Payload: map[string]string{
				"time": time.Now().String(),
			},
		}
		j.Fire()
	}
}
