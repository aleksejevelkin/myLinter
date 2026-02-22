package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"

	"github.com/aleksejevelkin/myLinter/checkers"

	"golang.org/x/tools/go/analysis"
)

// Analyzer — анализатор с настройками по умолчанию.
var Analyzer = New(Config{})

// New создаёт анализатор с заданной конфигурацией.
func New(cfg Config) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "loglint",
		Doc:  "checks log messages for common issues: lowercase start, English only, no special chars/emojis, no sensitive data",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			return run(pass, cfg)
		},
	}
}

func run(pass *analysis.Pass, cfg Config) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				if ident, ok := selExpr.X.(*ast.Ident); ok {
					if ident.Name == "log" || ident.Name == "slog" || ident.Name == "zap" {
						for _, arg := range callExpr.Args {
							lit, ok := arg.(*ast.BasicLit)
							if !ok || lit.Kind != token.STRING {
								continue
							}

							msg, err := strconv.Unquote(lit.Value)
							if err != nil {
								continue
							}

							if cfg.IsLowercaseEnabled() {
								if err := checkers.CheckLowercase(msg); err != nil {
									d := analysis.Diagnostic{Pos: lit.Pos(), Message: "log message issue: " + err.Error()}
									if newMsg, ok := fixLowercaseStart(msg); ok {
										if fix, ok := buildReplaceWholeLiteralFix(lit.Pos(), lit.Value, newMsg); ok {
											d.SuggestedFixes = []analysis.SuggestedFix{*fix}
										}
									}
									pass.Report(d)
								}
							}

							if cfg.IsEnglishOnlyEnabled() {
								if err := checkers.CheckEnglishOnly(msg); err != nil {
									pass.Reportf(lit.Pos(), "log message issue: %v", err)
								}
							}

							if cfg.IsSpecialCharsEnabled() {
								if err := checkers.CheckSpecialChars(msg); err != nil {
									d := analysis.Diagnostic{Pos: lit.Pos(), Message: "log message issue: " + err.Error()}
									if newMsg, ok := fixSpecial(msg); ok {
										if fix, ok := buildReplaceWholeLiteralFix(lit.Pos(), lit.Value, newMsg); ok {
											d.SuggestedFixes = []analysis.SuggestedFix{*fix}
										}
									}
									pass.Report(d)
								}
							}

							if cfg.IsSensitiveEnabled() {
								if err := checkers.CheckNoSensitiveDataWithKeywords(msg, cfg.SensitiveKeywords); err != nil {
									pass.Reportf(lit.Pos(), "log message issue: %v", err)
								}
							}
						}
					}
				}
			}
			return true
		})
	}

	return nil, nil
}
