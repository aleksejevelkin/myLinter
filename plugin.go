package plugin

import (
	"github.com/aleksejevelkin/myLinter/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", New)
}

type Settings struct {
	Lowercase         *bool    `json:"lowercase" mapstructure:"lowercase"`
	EnglishOnly       *bool    `json:"englishOnly" mapstructure:"englishOnly"`
	SpecialChars      *bool    `json:"specialChars" mapstructure:"specialChars"`
	Sensitive         *bool    `json:"sensitive" mapstructure:"sensitive"`
	SensitiveKeywords []string `json:"sensitiveKeywords" mapstructure:"sensitiveKeywords"`
}

type Plugin struct {
	cfg analyzer.Config
}

func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	cfg := analyzer.Config{
		Lowercase:         s.Lowercase,
		EnglishOnly:       s.EnglishOnly,
		SpecialChars:      s.SpecialChars,
		Sensitive:         s.Sensitive,
		SensitiveKeywords: s.SensitiveKeywords,
	}

	return &Plugin{cfg: cfg}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.New(p.cfg)}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
