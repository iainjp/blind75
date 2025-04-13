package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

// Example:
// {"Time":"2025-04-13T15:32:28.46481+01:00","Action":"start","Package":"github.com/iainjp/blind75/utils"}
type TestResult struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Output  string    `json:"Output"`
	Elapsed int       `json:"Elapsed"`
}

func Report() {
	binary, lookErr := exec.LookPath("go")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"test", "-json", "./..."}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	testCmd := exec.Command(binary, args...)
	testCmd.Stdout = &stdout
	testCmd.Stderr = &stderr

	execErr := testCmd.Run()
	if execErr != nil {
		fmt.Printf("Error running tests. Stderr: %v", stderr.String())
		panic(execErr)
	}

	results := []TestResult{}

	scanner := bufio.NewScanner(&stdout)
	for scanner.Scan() {
		r := TestResult{}
		json.Unmarshal(scanner.Bytes(), &r)
		results = append(results, r)
	}

	fmt.Println(results)
}
