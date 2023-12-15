package banfunc_test

import (
	"testing"

	"github.com/masakurapa/banfunc"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	a := banfunc.New()
	if err := a.Flags.Set("ban", "Println,Sprintf"); err != nil {
		t.Fatal(err)
	}
	analysistest.Run(t, analysistest.TestData(), a, "a")
}
