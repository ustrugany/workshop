package cmd

import (
	"net"
	"net/http"

	"example.com/concurrency/page"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newServerHandler() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		logger := logrus.New()
		logger.Formatter = new(logrus.JSONFormatter)
		port, _ := cmd.Flags().GetString("port")
		server, _ := cmd.Flags().GetString("server")
		router := mux.NewRouter()
		router.StrictSlash(true)
		router.
			Methods([]string{http.MethodGet}...).
			Path("/crawl/{url}").
			Name("results").
			Handler(page.NewCrawlHandler(logger))
		logger.Printf("listening on port %s:%s...", server, port)
		err := http.ListenAndServe(net.JoinHostPort(server, port), router)
		if err != nil {
			logger.Fatalf("http server boot failed: %s", err)
		}
	}
}

func NewServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Server command",
		Run:   newServerHandler(),
	}
}
