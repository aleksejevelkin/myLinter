package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"
	"testSelectel/checkers"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages for common issues: lowercase start, English only, no special chars/emojis, no sensitive data",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {

		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// ищем узел вида x.y
			if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				// ищем узел с идентификатором
				if ident, ok := selExpr.X.(*ast.Ident); ok {
					// проверяем, что идентификатор - это log, slog или zap
					if ident.Name == "log" || ident.Name == "slog" || ident.Name == "zap" {

						for _, arg := range callExpr.Args {
							// проверяем, что аргумент - это строковый литерал
							lit, ok := arg.(*ast.BasicLit)
							if !ok || lit.Kind != token.STRING {
								continue
							}

							// распаковываем строковый литерал, чтобы получить его значение без кавычек
							msg, err := strconv.Unquote(lit.Value)
							if err != nil {
								continue // Не удалось распарсить строку
							}

							if err := checkers.CheckLowercase(msg); err != nil {
								pass.Reportf(lit.Pos(), "log message issue: %v", err)
							}

							if err := checkers.CheckEnglishOnly(msg); err != nil {
								pass.Reportf(lit.Pos(), "log message issue: %v", err)
							}

							if err := checkers.CheckSpecialChars(msg); err != nil {
								pass.Reportf(lit.Pos(), "log message issue: %v", err)
							}

							if err := checkers.CheckNoSensitiveData(msg); err != nil {
								pass.Reportf(lit.Pos(), "log message issue: %v", err)
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
