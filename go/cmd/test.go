package cmd

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests for exercises",
	Run: func(cmd *cobra.Command, args []string) {
		binary, lookErr := exec.LookPath("go")
		if lookErr != nil {
			panic(lookErr)
		}

		cmdArgs := []string{"go", "tool", "gotestsum"}

		env := os.Environ()

		execErr := syscall.Exec(binary, cmdArgs, env)
		if execErr != nil {
			panic(execErr)
		}
	},
}
