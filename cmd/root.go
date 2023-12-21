package cmd

import (
	"github.com/boltdbgui/common/logger"
	"github.com/boltdbgui/server"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Boltdb Server",
	Long:  `Start the boltdb browser server`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.NewLogger("DEBUG")
		server.StartServer()
	},
}

var input = new(struct {
	dbPath string
})

func init() {
	rootCmd.Flags().StringVarP(&input.dbPath, "db-path", "d", "", "path to the bolt db file")
	err := rootCmd.MarkFlagRequired("db-path")
	if err != nil {
		panic(errors.Wrap(err, "error while setting required flag db-path"))
	}
}

func Execute() error {
	return rootCmd.Execute()
}
