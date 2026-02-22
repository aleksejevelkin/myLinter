package plugin

import (
	"github.com/aleksejevelkin/myLinter/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", New)
}

type Settings struct{}

type Plugin struct {
	settings Settings
}

func New(settings any) (register.LinterPlugin, error) {
	// Настройки сейчас не используются, но оставляем поддержку формата.
	_, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return &Plugin{settings: Settings{}}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	// Мы используем только AST, без types info.
	return register.LoadModeSyntax
}
