package package2

import (
	"go/ast"

	"github.com/110y/go-expr-completion/internal/analysis/internal/visitor"
)

func f1() {
	var v ast.Visitor
	v.(*visitor.Visitor)
}

func f2() {
	returnChannelReceiver()
	returnChannelSender()
	returnChannelSenderReceiver()
}

func returnChannelReceiver() <-chan struct{} {
	return nil
}

func returnChannelSender() chan<- struct{} {
	return nil
}

func returnChannelSenderReceiver() chan struct{} {
	return nil
}
