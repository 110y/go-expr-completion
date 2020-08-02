package run

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/110y/go-expr-completion/internal/analysis"
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	var pos int
	flag.IntVar(&pos, "pos", 0, "position of cursor")

	var path string
	flag.StringVar(&path, "file", "", "file")

	flag.Parse()

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %s", err.Error())
		return 1
	}

	if pos == 0 {
		// TODO: log
		fmt.Fprintf(os.Stderr, "must pass pos")
		return 1
	}

	types, err := analysis.GetExprTypeInfo(ctx, f.Name(), pos)
	if err != nil {
		// TODO: log
		fmt.Fprintf(os.Stderr, "failed to execute: %s", err.Error())
		return 1
	}

	if types == nil {
		fmt.Fprint(os.Stderr, "pos not in any Expressions")
		return 1
	}

	j, err := json.Marshal(types)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal results to json: %s", err.Error())
		return 1
	}

	fmt.Fprint(os.Stdout, string(j))

	return 0
}
