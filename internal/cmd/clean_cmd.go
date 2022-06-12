package cmd

import (
	"fmt"
	"github.com/razielsd/rzgrpcmock/internal/template"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean default grpc mock service",
	Long:  `clean default grpc mock service`,
	Run:   cleanTemplateExecute,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func cleanTemplateExecute(cmd *cobra.Command, args []string) {
	path := extractFirstArgOrDie(cmd, args, "Require path to clean default grpc mock service")
	tmlService := template.NewTemplateService()
	err := tmlService.Clean(path)
	message := fmt.Sprintf("Successfully saved: %s\n", path)
	if err != nil {
		message = err.Error()
	}
	printCli(message, err)
}