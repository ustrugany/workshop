package _go

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ustrugany/classifier/cmd"
	"github.com/ustrugany/classifier/pkg/config"
)

var (
	port       = flag.String("port", ":8080", "HTTP port")
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	//c := &classifier.Classifier{
	//	ServerAddr: *serverAddr,
	//	CaFile: *caFile,
	//	ServerHostOverride: *serverHostOverride,
	//	TLS: *tls,
	//}
	//uh := &handler.UploadHandler{Path: "./input", Classifier:c}
	//http.Handle("/classify", uh)
	//if err := http.ListenAndServe(*port, nil); err != nil {
	//	log.Fatal(err)
	//}
	configs := config.NewConfigs()
	rootCmd := &cobra.Command{
		Use:   os.Args[0],
		Short: "Classifier",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Classifier")
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	}

	//server command
	serverCmd := cmd.NewServerCommand(configs)
	serverCmd.PersistentFlags().StringP("port", "p", "", "port")
	crawlerCmd := cmd.NewCrawlerCommand(configs)
	crawlerCmd.PersistentFlags().StringP("url", "u", "", "url to crawl")
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(crawlerCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
