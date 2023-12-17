package banfunc_test

import (
	"testing"

	"github.com/masakurapa/banfunc"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyze(t *testing.T) {
	for _, tc := range []struct {
		ban string
		dir string
	}{
		{ban: "Println", dir: "single/a"},
		{ban: "fmt.Println", dir: "single/b"},
		{ban: "Println,Sprintf", dir: "multiple/a"},
		{ban: "fmt.Println,fmt.Sprintf", dir: "multiple/b"},
	} {
		t.Run(tc.ban, func(t *testing.T) {
			a := banfunc.New()
			if err := a.Flags.Set("ban", tc.ban); err != nil {
				t.Fatal(err)
			}
			analysistest.Run(t, analysistest.TestData(), a, tc.dir)
		})
	}
}
