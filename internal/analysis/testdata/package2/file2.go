package package2

import (
	"go/ast"

	"github.com/110y/go-expr-completion/internal/analysis/internal/visitor"
)

func f1() {
	var v ast.Visitor
	v.(*visitor.Visitor)
}
