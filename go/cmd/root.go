package cmd

import (
	"fmt"
	"os"

	cmd "github.com/iainjp/blind75/cmd/report"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "blind75",
	Short: "blind75 solutions, by iainjp",
	Long: `A suite of solutions, tests and benchmarks for the
                blind75 suite of Leetcode exercises`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RootCmd")
	},
}

func init() {
	addCmds()
}

func addCmds() {
	RootCmd.AddCommand(
		cmd.ReportCmd,
	)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
