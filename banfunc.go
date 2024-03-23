package banfunc

import (
	"fmt"
	"go/ast"
	"strings"
	"sync"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func New() *analysis.Analyzer {
	bf := &banFunc{}
	a := &analysis.Analyzer{
		Name: "banfunc",
		Doc: `banfunc is a Go linter that reports the call of a banned function.

Example:
	banfunc -ban Println ./...
	banfunc -ban Println,Print,Printf ./...`,
		Run:      bf.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	a.Flags.StringVar(&bf.ban, "ban", "", "banned function names(multiple comma separated)")
	return a
}

type banFunc struct {
	ban string

	banMap map[fnc]struct{}

	mux sync.Mutex
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

		var f fnc
		switch fun := call.Fun.(type) {
		case *ast.SelectorExpr:
			// example in this case is "fmt.Println()"
			f = fnc{name: fun.Sel.Name}
			if idt, ok := fun.X.(*ast.Ident); ok {
				f.structName = idt.Name
			}
		case *ast.Ident:
			// example in this case is "Println()"
			f = fnc{name: fun.Name}
		default:
			return
		}

		if bf.isBan(f) {
			pass.Report(analysis.Diagnostic{
				Pos:     call.Pos(),
				End:     call.End(),
				Message: fmt.Sprintf("%s is banned!", f.string()),
			})
		}
	})

	return nil, nil
}

func (bf *banFunc) initFuncMap() {
	ss := strings.Split(bf.ban, ",")
	bf.mux.Lock()
	bf.banMap = make(map[fnc]struct{})
	bf.mux.Unlock()
	for _, s := range ss {

		// split "Println" to ["Println"]
		// or
		// "fmt.Println" to ["fmt", "Println"]
		v := strings.SplitN(strings.TrimSpace(s), ".", 2)
		f := fnc{name: v[0]}
		if len(v) == 2 {
			f = fnc{structName: v[0], name: v[1]}
		}

		bf.mux.Lock()
		bf.banMap[f] = struct{}{}
		bf.mux.Unlock()
	}
}

func (bf *banFunc) isBan(f fnc) bool {
	bf.mux.Lock()
	defer bf.mux.Unlock()
	if _, ok := bf.banMap[f]; ok {
		return true
	}
	if _, ok := bf.banMap[f.nameOnly()]; ok {
		return true
	}
	return false
}

type fnc struct {
	structName string
	name       string
}

func (f *fnc) nameOnly() fnc {
	return fnc{name: f.name}
}

func (f *fnc) string() string {
	if f.structName == "" {
		return f.name
	}
	return fmt.Sprintf("%s.%s", f.structName, f.name)
}
