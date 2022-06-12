package cmd

import (
	"fmt"
	"github.com/razielsd/rzgrpcmock/internal/generator"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate mock api",
	Long:  `generate mock api`,
	Run: genServiceExecute,
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func genServiceExecute(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		_ = cmd.Usage()
		return
	}
	builder := generator.NewBuilder()
	err := builder.Run(args[0], args[1])
	message := ""
	if err == nil {
		message = fmt.Sprintf("Succesefully saved: %s\n", "")
	}
	printCli(message, err)
}