package main

import (
	"github.com/masakurapa/banfunc"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(banfunc.New())
}
