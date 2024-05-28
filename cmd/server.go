package cmd

import (
	"net/http"

	"github.com/spf13/cobra"
	"ozon/pkg/gql"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "start api server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server()
	},
}

func server() error {
	http.HandleFunc("/graphql", gql.Handler)

	println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	return nil
}
