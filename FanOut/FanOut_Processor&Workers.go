package main

import (
	"fmt"
	"time"
)

// URL: https://medium.com/geekculture/golang-concurrency-patterns-fan-in-fan-out-1ee43c6830c4

var workersPool int = 3
var taskQuantity int = 10

func generate(data string) <-chan string {
	channel := make(chan string)
	i := 1
	go func() {
		for {
			channel <- fmt.Sprint(data, "-", i)
			time.Sleep(3000)
			i++
		}
	}()

	return channel
}

type Processor struct {
	jobChannel chan string
	done       chan *Worker
	workers    []*Worker
}
type Worker struct {
	name string
}

func (w *Worker) processJob(data string, done chan *Worker) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, w.name)
		time.Sleep(300)
		done <- w
	}()

}

func GetProcessor() *Processor {
	p := &Processor{
		jobChannel: make(chan string),
		workers:    make([]*Worker, workersPool),
		done:       make(chan *Worker),
	}
	for i := 0; i < workersPool; i++ {
		w := &Worker{name: fmt.Sprintf("<Worker - %d>", i+1)}
		p.workers[i] = w
	}
	p.startProcess()
	return p
}

func (p *Processor) startProcess() {
	go func() {
		for {
			select {
			default:
				if len(p.workers) > 0 {
					w := p.workers[0]
					p.workers = p.workers[1:]
					w.processJob(<-p.jobChannel, p.done)
				}
			case w := <-p.done:
				p.workers = append(p.workers, w)
			}
		}
	}()
}

func (p *Processor) postJob(jobs <-chan string) {
	p.jobChannel <- <-jobs
}

func main() {
	source := generate("data string")
	process := GetProcessor()

	for i := 0; i < taskQuantity; i++ {
		process.postJob(source)
	}

}
