package main

import (
	"log"

	pool "github.com/mediaprodcast/commons/pool/minimal"
)

func main() {
	tasks := []*pool.Task{
		pool.NewTask(func() error { return nil }),
		pool.NewTask(func() error { return nil }),
		pool.NewTask(func() error { return nil }),
	}

	concurrency := 2
	p := pool.NewPool(tasks, concurrency)
	p.Run()

	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			log.Println(task.Err)
			numErrors++
		}
		if numErrors >= 10 {
			log.Println("Too many errors.")
			break
		}
	}
}
