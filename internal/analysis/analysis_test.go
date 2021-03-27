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
		"map": {
			path: "package2/file1.go",
			pos:  56,
			expected: &analysis.TypeInfo{
				StartPos: 56,
				EndPos:   64,
				Values: []*analysis.Value{
					{
						Name: "s",
						Type: "string",
					},
					{
						Name: "ok",
						Type: "bool",
					},
				},
			},
		},
		"map field": {
			path: "package2/file1.go",
			pos:  163,
			expected: &analysis.TypeInfo{
				StartPos: 163,
				EndPos:   180,
				Values: []*analysis.Value{
					{
						Name: "i",
						Type: "interface{}",
					},
					{
						Name: "ok",
						Type: "bool",
					},
				},
			},
		},
		"nested map field": {
			path: "package2/file1.go",
			pos:  163,
			expected: &analysis.TypeInfo{
				StartPos: 163,
				EndPos:   180,
				Values: []*analysis.Value{
					{
						Name: "i",
						Type: "interface{}",
					},
					{
						Name: "ok",
						Type: "bool",
					},
				},
			},
		},
		"type assertion": {
			path: "package2/file2.go",
			pos:  146,
			expected: &analysis.TypeInfo{
				StartPos: 146,
				EndPos:   166,
				Values: []*analysis.Value{
					{
						Name: "v",
						Type: "*github.com/110y/go-expr-completion/internal/analysis/internal/visitor.Visitor",
					},
					{
						Name: "ok",
						Type: "bool",
					},
				},
			},
		},
		"receiver channel": {
			path: "package2/file2.go",
			pos:  203,
			expected: &analysis.TypeInfo{
				StartPos: 183,
				EndPos:   206,
				Values: []*analysis.Value{
					{
						Name: "ch",
						Type: "<-chan struct{}",
					},
				},
			},
		},
		"sender channel": {
			path: "package2/file2.go",
			pos:  208,
			expected: &analysis.TypeInfo{
				StartPos: 208,
				EndPos:   229,
				Values: []*analysis.Value{
					{
						Name: "ch",
						Type: "chan<- struct{}",
					},
				},
			},
		},
		"send recv channel": {
			path: "package2/file2.go",
			pos:  231,
			expected: &analysis.TypeInfo{
				StartPos: 231,
				EndPos:   260,
				Values: []*analysis.Value{
					{
						Name: "ch",
						Type: "chan struct{}",
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
