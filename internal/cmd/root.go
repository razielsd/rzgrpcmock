package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use:   "rzgrpcmock",
	Short: "rzgrpcmock",
	Long:  `RzGrpcMock service cli`,
	Version: "0.0.1",
}

// Execute main run point.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func extractFirstArg(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("empty args")
	}
	return args[0], nil
}

func extractFirstArgOrDie(cmd *cobra.Command, args []string, errMessage string) string {
	param, err := extractFirstArg(args)
	if err != nil {
		fmt.Println(errMessage)
		_ = cmd.Usage()
		os.Exit(1)
	}
	return param
}

func printCli(data string, err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Println(data)
	}
}