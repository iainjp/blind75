package app

import (
	"os"
	"os/exec"
	"syscall"
)

func Report() {
	binary, lookErr := exec.LookPath("go")
	if lookErr != nil {
		panic(lookErr)
	}

	cmdArgs := []string{"go", "test", "-json", "./..."}

	env := os.Environ()

	execErr := syscall.Exec(binary, cmdArgs, env)
	if execErr != nil {
		panic(execErr)
	}
}
