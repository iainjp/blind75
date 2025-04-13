package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Print report of exercises",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("// NOT IMPLEMENTED YET")
	},
}
