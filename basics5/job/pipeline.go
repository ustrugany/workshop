package job

import (
	"context"
	"os"
	"os/signal"
	"runtime/trace"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Pipeline struct {
	workers   []*Worker
	generator *Generator
	logger    *logrus.Logger
	channel   chan Job
}

func NewPipeline(logger *logrus.Logger, concurrency, delay, frequency int) Pipeline {
	pipeline := Pipeline{
		channel: make(chan Job),
		logger:  logger,
	}

	// we have only this number of concurrent workers to process
	for i := 0; i < concurrency; i++ {
		pipeline.workers = append(pipeline.workers, NewWorker(logger, i))
	}

	// job generator
	pipeline.generator = NewGenerator(logger, pipeline.channel)

	return pipeline
}

func (p *Pipeline) Run(ctx context.Context, timeout, count int) {
	// ignore
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)
	defer cancel()

	// ignore
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// ignore

	// ignore
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, os.Interrupt, os.Kill)
	go func() {
		signal := <-gracefulStop
		p.logger.Errorf("caught signal: [%+v]", signal)
		os.Exit(1)
	}()

	// IMPORTANT PART:
	// channel to receive results
	results := make(chan Job)
	// waits for collection of goroutines to finish
	wg := &sync.WaitGroup{}
	p.generator.Generate(ctx, count)

	// start workers
	for _, worker := range p.workers {
		wg.Add(1)
		go worker.Work(ctx, wg, p.channel, results)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				p.logger.Errorf("timeout in %d ms...", timeout)
				return
			case result, ok := <-results:
				if !ok {
					p.logger.Println("result channel closed...")
					return
				}
				p.logger.Printf("result %v", result)
			}
		}
	}()
	wg.Wait()
}
