package cmd

import (
	app "github.com/iainjp/blind75/app/reporter"
	"github.com/spf13/cobra"
)

var ReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Print report of exercises",
	Run: func(cmd *cobra.Command, args []string) {
		app.Report()
	},
}
