package banfunc

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var (
	Analyzer = &analysis.Analyzer{
		Name: "banfunc",
		Doc: `banfunc is a linter that reports the call of a banned function.

Example:
	banfunc -func Println ./...
	banfunc -func Println,Print,Printf ./...`,
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	funcs string
)

func init() {
	Analyzer.Flags.StringVar(&funcs, "funcs", "", "banned function names(multiple comma separated)")
}

type f struct {
	name string
}

func makeFuncMap() map[f]struct{} {
	ss := strings.Split(funcs, ",")
	m := make(map[f]struct{})
	for _, s := range ss {
		m[f{name: strings.TrimSpace(s)}] = struct{}{}
	}
	return m
}

func run(pass *analysis.Pass) (interface{}, error) {
	fm := makeFuncMap()
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	ins.Preorder(nodeFilter, func(node ast.Node) {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		k := f{name: selector.Sel.Name}
		if _, ok = fm[k]; ok {
			pass.Report(analysis.Diagnostic{
				Pos:     call.Pos(),
				End:     call.End(),
				Message: fmt.Sprintf("%s is banned!", k.name),
			})
		}
	})

	return nil, nil
}
