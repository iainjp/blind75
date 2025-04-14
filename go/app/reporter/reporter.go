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
	Elapsed     float64
	BytesPerOp  string
	AllocsPerOp string
	NsPerOp     string
	TestCount   int
	TestPass    int
	TestFail    int
	Emoji       string
}

func buildPackageLink(packageName string) string {
	relativePath := "go/" + strings.Split(packageName, "blind75/")[1]
	name := strings.Split(packageName, "exercises/")[1]
	return fmt.Sprintf("[%v](%v)", name, relativePath)
}

func getEmoji(testSummary TestSummary) string {
	if testSummary.TestFail == 0 {
		return ":white_check_mark:"
	} else {
		return ":x:"
	}
}

func BuildViewModel(testSummary TestSummary) ResultViewModel {
	packageLink := buildPackageLink(testSummary.Package)
	emoji := getEmoji(testSummary)

	return ResultViewModel{
		Package:     packageLink,
		Elapsed:     testSummary.Elapsed,
		BytesPerOp:  testSummary.BytesPerOp,
		AllocsPerOp: testSummary.AllocsPerOp,
		NsPerOp:     testSummary.NsPerOp,
		TestCount:   testSummary.TestCount,
		TestPass:    testSummary.TestPass,
		TestFail:    testSummary.TestFail,
		Emoji:       emoji,
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
