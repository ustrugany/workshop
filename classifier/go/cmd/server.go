package cmd

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/ustrugany/classifier/pkg/config"
	"github.com/ustrugany/classifier/pkg/upload"
)

func newServerCommandHandler(configs config.Configurations) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		logger := logrus.New()
		logger.Formatter = new(logrus.JSONFormatter)

		port, _ := cmd.Flags().GetString("port")
		host, _ := cmd.Flags().GetString("host")

		router := mux.NewRouter()
		router.StrictSlash(true)
		router.
			Methods([]string{http.MethodGet}...).
			Path("/results").
			Name("results").
			Handler(&upload.UploadHandler{})
		router.
			PathPrefix("/static/").
			Handler(http.StripPrefix(
				"/static/",
				http.FileServer(http.Dir("./web/static")),
			))

		logger.Println("listening on port :%s...", port)
		err := http.ListenAndServe(net.JoinHostPort(host, port), router)
		if err != nil {
			logger.Fatalf("http server boot failed: %s", err)
		}
	}
}

func NewServerCommand(configuration config.Configurations) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Server command",
		Run:   newServerCommandHandler(configuration),
	}
}
