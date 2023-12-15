package banfunc_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/masakurapa/banfunc"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, banfunc.Analyzer, "a")
}
