package analysis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/110y/go-expr-completion/internal/analysis"
)

func TestGetExprTypeInfo(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		path     string
		pos      int
		expected *analysis.TypeInfo
	}{
		"function call which returns single context.Context variable": {
			path: "package1/file1.go",
			pos:  67,
			expected: &analysis.TypeInfo{
				StartPos: 67,
				EndPos:   87,
				Values: []*analysis.Value{
					{
						Name: "ctx",
						Type: "context.Context",
					},
				},
			},
		},
		"function call which returns two variables": {
			path: "package1/file1.go",
			pos:  160,
			expected: &analysis.TypeInfo{
				StartPos: 126,
				EndPos:   160,
				Values: []*analysis.Value{
					{
						Name: "i",
						Type: "int",
					},
					{
						Name: "err",
						Type: "error",
					},
				},
			},
		},
		"pos is in the function literal": {
			path: "package1/file2.go",
			pos:  166,
			expected: &analysis.TypeInfo{
				StartPos: 164,
				EndPos:   188,
				Values: []*analysis.Value{
					{
						Name: "v",
						Type: "*github.com/110y/go-expr-completion/internal/analysis/internal/visitor.Visitor",
					},
				},
			},
		},
		"pos is in the function arguments": {
			path: "package1/file2.go",
			pos:  242,
			expected: &analysis.TypeInfo{
				StartPos: 222,
				EndPos:   246,
				Values: []*analysis.Value{
					{
						Name: "v",
						Type: "*github.com/110y/go-expr-completion/internal/analysis/internal/visitor.Visitor",
					},
				},
			},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			path := fmt.Sprintf("testdata/%s", test.path)
			actual, err := analysis.GetExprTypeInfo(context.Background(), path, test.pos)
			if err != nil {
				t.Fatalf("error: %s\n", err.Error())
			}

			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Errorf("\n(-expected, +actual)\n%s", diff)
			}
		})
	}
}
