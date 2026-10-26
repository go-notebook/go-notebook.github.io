# A notebook is Markdown

This page is a notebook: a Markdown document whose `go` blocks are cells. Edit any cell and press Ctrl+Enter (⌘+Enter) to run the notebook again.

A cell may declare functions and types, and run statements at the top level. What it defines, the cells after it see.

```go
import (
	"sort"
	"strings"
)

type Count struct {
	Word string
	N    int
}

func count(text string) []Count {
	seen := map[string]int{}
	for _, w := range strings.Fields(strings.ToLower(text)) {
		seen[strings.Trim(w, ".,;:!?")]++
	}
	out := make([]Count, 0, len(seen))
	for w, n := range seen {
		out = append(out, Count{w, n})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].N != out[j].N {
			return out[i].N > out[j].N
		}
		return out[i].Word < out[j].Word
	})
	return out
}
```

The last expression of a cell is shown under it; a struct shows its fields by name.

```go
counts := count("To be, or not to be: that is the question.")
counts[0]
```

```go
len(counts)
```
