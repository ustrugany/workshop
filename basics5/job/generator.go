package job

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type Job struct {
	Result   string
	WorkerId int
}

type Generator struct {
	logger  *logrus.Logger
	channel chan Job
}

func NewGenerator(logger *logrus.Logger, channel chan Job) *Generator {
	return &Generator{logger: logger, channel: channel}
}
func (g *Generator) Generate(ctx context.Context, count int) {
	send := 0
	go func() {
		defer close(g.channel)
		for {
			select {
			case <-ctx.Done():
				g.logger.Println("generator is cleaning up")
				return
			default:
				if send >= count {
					return
				}
				time.Sleep(time.Duration(50) * time.Millisecond)
				g.logger.Printf("generator sending job %d", send)
				g.channel <- Job{}
				send++
			}
		}
	}()
}
