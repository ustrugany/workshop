package main

import (
	"context"
	"fmt"
	"os"

	"example.com/concurrency/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   os.Args[0],
		Short: "concurrency",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("concurrency")
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	}

	workerCmd := cmd.NewWorkerCommand(context.Background())
	workerCmd.PersistentFlags().IntP("concurrency", "c", 1, "Concurrency factor")
	workerCmd.PersistentFlags().IntP("frequency", "f", 1, "Jobs frequency in millisecond")
	workerCmd.PersistentFlags().IntP("delay", "d", 1, "Worker delay in millisecond")
	workerCmd.PersistentFlags().IntP("jobs", "j", 1000, "Jobs factor")
	workerCmd.PersistentFlags().IntP("timeout", "t", 10, "Timeout in seconds")
	rootCmd.AddCommand(workerCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
