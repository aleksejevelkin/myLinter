package main

import (
	"testSelectel/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {

	singlechecker.Main(analyzer.Analyzer)

}
