package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"slices"
	"sort"
	"strings"
	"time"
)

// Examples:
// {"Time":"2025-04-13T15:32:28.46481+01:00","Action":"start","Package":"github.com/iainjp/blind75/utils"}
// {"Time":"2025-04-13T16:35:06.05039+01:00","Action":"pass","Package":"github.com/iainjp/blind75/app/exercises/03-contains-duplicates","Test":"TestContainsDuplicates","Elapsed":0}
// {"Time":"2025-04-13T16:35:06.050743+01:00","Action":"output","Package":"github.com/iainjp/blind75/app/exercises/03-contains-duplicates","Output":"PASS\n"}
// {"Time":"2025-04-13T16:35:06.051109+01:00","Action":"output","Package":"github.com/iainjp/blind75/app/exercises/03-contains-duplicates","Output":"ok  \tgithub.com/iainjp/blind75/app/exercises/03-contains-duplicates\t0.151s\n"}
// {"Time":"2025-04-13T16:35:06.051121+01:00","Action":"pass","Package":"github.com/iainjp/blind75/app/exercises/03-contains-duplicates","Elapsed":0.151}

// Struct containing results of a golang test
type TestResult struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Output  string    `json:"Output"`
	Elapsed float64   `json:"Elapsed"`
	Test    string    `json:"Test"`
}

type TestSummary struct {
	Package     string
	Elapsed     float64
	BytesPerOp  string
	AllocsPerOp string
	NsPerOp     string
	TestCount   int
	TestPass    int
	TestFail    int
}

func runTestsAndBenchmarks() []TestResult {
	binary, lookErr := exec.LookPath("go")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"test", "-bench=.", "-benchmem", "-json", "./..."}

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

	wantedActions := []string{"pass", "fail", "output"}
	scanner := bufio.NewScanner(&stdout)
	for scanner.Scan() {
		tr := TestResult{}
		json.Unmarshal(scanner.Bytes(), &tr)

		if slices.Contains(wantedActions, tr.Action) {
			results = append(results, tr)
		}
	}

	return results
}

func summarizeResults(results []TestResult) []TestSummary {
	groupedByExercise := map[string][]TestResult{}
	for _, result := range results {
		// skip project tests unrelated to exercises
		if !strings.Contains(result.Package, "exercises") {
			continue
		}
		val, ok := groupedByExercise[result.Package]
		if !ok {
			val = []TestResult{}
		}

		groupedByExercise[result.Package] = append(val, result)
	}

	reports := []TestSummary{}

	for pkg, results := range groupedByExercise {
		elapsed := 0.0
		passes := 0
		fails := 0
		var nsPerOp string
		var bytesPerOp string
		var allocsPerOp string

		// grab relevant details from results
		for _, res := range results {
			switch res.Action {
			case "pass":
				if !strings.Contains(res.Test, "Benchmark") && strings.Contains(res.Test, "Case") {
					passes += 1
				}
			case "fail":
				fails += 1
			}

			// identify relevant benchmark results
			if strings.Contains(res.Test, "Benchmark") && strings.Contains(res.Output, "ns/op") {
				split := strings.Split(res.Output, "\t")
				nsPerOp = strings.TrimSpace(split[2])
				bytesPerOp = strings.TrimSpace(split[3])
				allocsPerOp = strings.TrimSpace(split[4])
			}

			// identify result has total elapsed time across tests for single package
			if res.Action == "pass" && res.Test == "" {
				elapsed = res.Elapsed
			}
		}

		report := TestSummary{
			Package:     pkg,
			Elapsed:     elapsed,
			BytesPerOp:  bytesPerOp,
			NsPerOp:     nsPerOp,
			AllocsPerOp: allocsPerOp,
			TestCount:   passes + fails,
			TestPass:    passes,
			TestFail:    fails,
		}

		reports = append(reports, report)
	}

	sortByPackageName := func(i, j int) bool { return reports[i].Package < reports[j].Package }
	sort.Slice(reports, sortByPackageName)

	return reports
}

func RunTests() []TestSummary {
	results := runTestsAndBenchmarks()
	return summarizeResults(results)
}
