package package1

import (
	"context"
	"fmt"
	"os"
)

func f1() {
	context.Background()
}

func f2() {
	// some
	// comments
	fmt.Fprintln(os.Stdout, "message")
}
