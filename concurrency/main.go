package main

import (
	"fmt"
	"os"

	"example.com/concurrency/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   os.Args[0],
		Short: "Server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Server")
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	}

	serverCmd := cmd.NewServerCommand()
	serverCmd.PersistentFlags().StringP("port", "p", "8080", "port")
	serverCmd.PersistentFlags().StringP("server", "s", "127.0.0.1", "server")
	rootCmd.AddCommand(serverCmd)

	crawlerCmd := cmd.NewCrawlerCommand()
	crawlerCmd.PersistentFlags().StringP("url", "u", "", "url to crawl")
	crawlerCmd.PersistentFlags().IntP("concurrency", "c", 10, "concurrency factor")
	crawlerCmd.PersistentFlags().IntP("timeout", "t", 30, "timeout in sec")
	rootCmd.AddCommand(crawlerCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
