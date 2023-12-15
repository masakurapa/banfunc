package banfunc

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type banFunc struct {
	ban string

	banMap map[fnc]struct{}
}

type fnc struct {
	name string
}

func New() *analysis.Analyzer {
	bf := &banFunc{}
	a := &analysis.Analyzer{
		Name: "banfunc",
		Doc: `banfunc is a linter that reports the call of a banned function.

Example:
	banfunc -ban Println ./...
	banfunc -ban Println,Print,Printf ./...`,
		Run:      bf.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	a.Flags.StringVar(&bf.ban, "ban", "", "banned function names(multiple comma separated)")
	return a
}

func (bf *banFunc) initFuncMap() {
	ss := strings.Split(bf.ban, ",")
	bf.banMap = make(map[fnc]struct{})
	for _, s := range ss {
		bf.banMap[fnc{name: strings.TrimSpace(s)}] = struct{}{}
	}
}

func (bf *banFunc) run(pass *analysis.Pass) (interface{}, error) {
	bf.initFuncMap()

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

		f := fnc{name: selector.Sel.Name}
		if _, ok = bf.banMap[f]; ok {
			pass.Report(analysis.Diagnostic{
				Pos:     call.Pos(),
				End:     call.End(),
				Message: fmt.Sprintf("%s is banned!", f.name),
			})
		}
	})

	return nil, nil
}
