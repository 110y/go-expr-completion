# go-expr-completion

A tool to complete a left-hand side from given expression for Go.
<br>This tool is aimed to be integrated with text editors.

## Usage


### Example

Let's assume that your `cursor` is on the `F` on line:16 (`fmt.Fprintln`) with this file:


`./internal/analysis/testdata/package1/file1.go`

```go
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
```

The `F` is `130` as byte offsets for the file.<br>
So, let's specify the file path and `-pos 130` argument to `go-expr-completion` like below:

```sh
> go-expr-completion -pos 130 -file ./internal/analysis/testdata/package1/file1.go | jq .

{
  "start_pos": 126,
  "end_pos": 160,
  "values": [
    {
      "name": "i",
      "type": "int"
    },
    {
      "name": "err",
      "type": "error"
    }
  ]
}
```

Finally, `go-expr-completion` returs type information about `fmt.Fprintln` which is specified by your cursor.<br>
By using this type information in text editors, we can complete the left-hand side from the expression like below (Vim):

![01d57234-441d-46ef-bb0d-7b8f336b84b4](https://user-images.githubusercontent.com/2134196/89279213-12ef8100-d682-11ea-8b93-5660b232255d.gif)

## Text Editor Plugins

- Vim
    - https://github.com/110y/vim-go-expr-completion/
- Other Text Editors
    - Help Wanted :pray:
