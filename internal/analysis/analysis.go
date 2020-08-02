package analysis

import (
	"context"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/tools/go/packages"

	"github.com/110y/go-expr-completion/internal/analysis/internal/visitor"
)

var (
	pkgConfigMode = packages.NeedName |
		packages.NeedFiles |
		packages.NeedImports |
		packages.NeedSyntax |
		packages.NeedTypesInfo |
		packages.NeedTypes |
		packages.NeedTypesSizes

	specializedTypeVarNameMap = map[string]string{
		"bool":            "ok",
		"error":           "err",
		"context.Context": "ctx",
	}
)

func GetExprTypeInfo(ctx context.Context, path string, pos int) (*TypeInfo, error) {
	fs := token.NewFileSet()

	fpath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get abs file path: %w", err)
	}

	cfg := &packages.Config{
		Context: ctx,
		Fset:    fs,
		Dir:     filepath.Dir(fpath),
		Mode:    pkgConfigMode,
		Tests:   true,
	}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}

	var pkgIdx int
	var fIdx int
	for i, pkg := range pkgs {
		for j, f := range pkg.GoFiles {
			if fpath == f {
				fIdx = j
				pkgIdx = i
			}
		}
	}

	pkg := pkgs[pkgIdx]
	f := pkg.Syntax[fIdx]

	v := visitor.New(pos, fs, pkg.TypesInfo)
	ast.Walk(v, f)

	t := v.GetType()
	if t == nil {
		return nil, nil
	}

	tf := &TypeInfo{
		StartPos: v.GetStartPos(),
		EndPos:   v.GetEndPos(),
		Values:   createTypeInfo(t),
	}

	return tf, nil
}

func createTypeInfo(t types.Type) []*Value {
	switch t := t.(type) {
	case *types.Tuple:
		result := make([]*Value, t.Len())
		for i := 0; i < t.Len(); i++ {
			result[i] = createTypeInfo(t.At(i).Type())[0]
		}
		return result

	default:
		s := t.String()
		n, ok := specializedTypeVarNameMap[s]
		if ok {
			return []*Value{
				{
					Name: n,
					Type: s,
				},
			}
		}

		n = strings.TrimPrefix(strings.TrimPrefix(s, "[]"), "*")

		splited := strings.Split(n, ".")
		if len(splited) > 1 {
			n = splited[len(splited)-1]
		}

		return []*Value{
			{
				Name: lowerFirstChar(n),
				Type: t.String(),
			},
		}
	}
}

func lowerFirstChar(str string) string {
	for _, s := range str {
		return string(unicode.ToLower(s))
	}

	return ""
}
