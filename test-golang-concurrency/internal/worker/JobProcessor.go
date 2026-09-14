package worker

import (
	"context"
	"fmt"
)

type JobProcessor struct {
}

type IJobProcessor interface {
	StartWorker(ctx context.Context, jobs <-chan int)
	TriggerJob(jobs chan<- int, jobId int)
}

func NewJobProcessor() IJobProcessor {
	return &JobProcessor{}
}

func (jp *JobProcessor) StartWorker(ctx context.Context, jobs <-chan int) {
	fmt.Println("Starting worker...")
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("job channel has been closed")
				return
			}
			jp.process(job)
		case <-ctx.Done():
			fmt.Printf("worker stopped: %v\n", ctx.Err())
			return
		}
	}
}

func (jp *JobProcessor) TriggerJob(jobs chan<- int, jobId int) {
	jobs <- jobId
}

func (jp *JobProcessor) process(job int) {
	fmt.Printf("running job %d\n", job)
}
