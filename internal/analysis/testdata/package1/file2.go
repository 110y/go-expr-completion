// package1
// package comments
package package1

import (
	"github.com/110y/go-expr-completion/internal/analysis/internal/visitor"
)

func f1() {
	f := func() {
		visitor.New(0, nil, nil)
	}
}

func f2() {
	f3(func() {
		visitor.New(0, nil, nil)
	})
}

func f3(f func()) {
	f()
}
