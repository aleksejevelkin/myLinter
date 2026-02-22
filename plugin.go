package loglint

import (
	//"github.com/aleksejevelkin/myLinter/analyzer"
	"golang.org/x/tools/go/analysis"
	"testSelectel/analyzer"
)

func New(settings any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}
