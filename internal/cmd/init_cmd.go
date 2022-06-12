package cmd

import (
	"fmt"
	"github.com/razielsd/rzgrpcmock/internal/template"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init default grpc mock service",
	Long:  `init default grpc mock service`,
	Run:   initTemplateExecute,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initTemplateExecute(cmd *cobra.Command, args []string) {
	path := extractFirstArgOrDie(cmd, args, "Require path to init default grpc mock service")
	tmlService := template.NewTemplateService()
	err := tmlService.Init(path)
	message := fmt.Sprintf("Successfully saved: %s\n", path)
	if err != nil {
		message = err.Error()
	}
	printCli(message, err)
}