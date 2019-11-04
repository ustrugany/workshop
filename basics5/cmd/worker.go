package cmd

import (
	"context"

	"example.com/concurrency/job"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newWorkerHandler(ctx context.Context) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		logger := logrus.New()
		logger.Formatter = new(logrus.JSONFormatter)
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		jobs, _ := cmd.Flags().GetInt("jobs")
		timeout, _ := cmd.Flags().GetInt("timeout")
		delay, _ := cmd.Flags().GetInt("delay")
		frequency, _ := cmd.Flags().GetInt("frequency")

		pipeline := job.NewPipeline(logger, concurrency, delay, frequency)
		pipeline.Run(ctx, timeout, jobs)
	}
}

func NewWorkerCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "worker command",
		Run:   newWorkerHandler(ctx),
	}
}
