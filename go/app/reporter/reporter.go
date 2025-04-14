package app

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type ResultViewModel struct {
	Package     string
	Elapsed     string
	BytesPerOp  string
	AllocsPerOp string
	NsPerOp     string
	TestResult  string
}

func buildPackageLink(packageName string) string {
	relativePath := "go/" + strings.Split(packageName, "blind75/")[1]
	name := strings.Split(packageName, "exercises/")[1]
	return fmt.Sprintf("[%v](%v)", name, relativePath)
}

func buildTestResultString(testSummary TestSummary) string {
	var emoji string
	if testSummary.TestFail == 0 {
		emoji = ":white_check_mark:"
	} else {
		emoji = ":x:"
	}

	return fmt.Sprintf("%v/%v %v", testSummary.TestPass, testSummary.TestCount, emoji)
}

func buildElapsedString(testSummary TestSummary) string {
	return fmt.Sprintf("%.3f", testSummary.Elapsed)
}

func BuildViewModel(testSummary TestSummary) ResultViewModel {
	packageLink := buildPackageLink(testSummary.Package)
	testResultString := buildTestResultString(testSummary)
	elapsedString := buildElapsedString(testSummary)

	return ResultViewModel{
		Package:     packageLink,
		Elapsed:     elapsedString,
		BytesPerOp:  testSummary.BytesPerOp,
		AllocsPerOp: testSummary.AllocsPerOp,
		NsPerOp:     testSummary.NsPerOp,
		TestResult:  testResultString,
	}
}

func Report() {
	summaries := RunTests()

	viewModels := []ResultViewModel{}
	for _, s := range summaries {
		viewModels = append(viewModels, BuildViewModel(s))
	}

	// path from execution root (go./ dir) to README.md (top of repo)
	outFile := filepath.Join("..", "README.md")
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	var tmplFile = "_README.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, viewModels)
	if err != nil {
		panic(err)
	}
}
