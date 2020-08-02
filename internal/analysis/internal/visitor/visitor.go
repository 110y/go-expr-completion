package visitor

import (
	"go/ast"
	"go/token"
	"go/types"
)

var _ ast.Visitor = (*Visitor)(nil)

func New(pos int, fs *token.FileSet, info *types.Info) *Visitor {
	return &Visitor{
		fileset:   fs,
		cursorPos: pos,
		info:      info,
	}
}

type Visitor struct {
	cursorPos int
	fileset   *token.FileSet
	info      *types.Info
	types     types.Type
	startPos  int
	endPos    int

	inFuncDecl     bool
	inFuncDeclBody bool
}

func (v *Visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return nil
	}

	defer func() {
		switch node.(type) {
		case *ast.FuncDecl, *ast.FuncLit:
			v.inFuncDecl = true
			v.inFuncDeclBody = false
		case *ast.BlockStmt:
			if v.inFuncDecl {
				v.inFuncDeclBody = true
			}
		}
	}()

	startPos := v.getPositionOffset(node.Pos())
	endPos := v.getPositionOffset(node.End())

	if v.cursorPos < startPos || v.cursorPos > endPos {
		return nil
	}

	expr, ok := node.(ast.Expr)
	if !ok {
		return v
	}

	if v.inFuncDecl && v.inFuncDeclBody {
		if tv, ok := v.info.Types[expr]; ok {
			v.types = tv.Type
			v.startPos = startPos
			v.endPos = endPos

			v.inFuncDecl = false
			v.inFuncDeclBody = false
		}
	}

	return v
}

func (v *Visitor) GetType() types.Type {
	return v.types
}

func (v *Visitor) GetStartPos() int {
	return v.startPos
}

func (v *Visitor) GetEndPos() int {
	return v.endPos
}

func (v *Visitor) getPositionOffset(pos token.Pos) int {
	return v.fileset.Position(pos).Offset
}
