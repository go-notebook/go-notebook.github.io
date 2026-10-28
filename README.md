# go-notebook.github.io

The go-notebook landing page, https://go-notebook.github.io/, built with Hugo
and the fleet's [landing kit](https://github.com/go-fleettools/landingkit)
(a Hugo module pinned in `go.mod`).

- **[/playground/](https://go-notebook.github.io/playground/)**: notebooks
  in the browser, Go and Ruby cells, run in WebAssembly with no server. The
  deploy workflow builds it with
  [go-notebook/playground](https://github.com/go-notebook/playground)'s
  `cmd/build`, at the version `playground/go.mod` pins, from the notebooks in
  `playground/notebooks/` (it runs each one first and fails if a cell fails).
  The red "Try it in your browser" button in the hero leads there. While its
  first WebAssembly core downloads, the page shows what go-notebook does: the
  landing kit's loading carousel (playground v0.3.0), never on a cached load.
  Each cell carries a Go or Ruby badge read from its fence (v0.4.0); cells can
  be added, edited, run one at a time, switched between Go, Ruby, Julia and
  Markdown, deleted with undo and moved, with Jupyter's keys (v0.6.0).
- **The hero** is the kit's carousel: the README's notebook
  (`layouts/partials/hero-terminal.html`), a notebook with a Ruby cell and
  what it printed (`hero-terminal-ruby.html`), and the playground.
- **The Ruby card is a real run.** It shows `example/waits.md` and
  `example/output.txt`, read from the files. To refresh the output after a
  bump of `example/go.mod`:

  ```sh
  cd example
  go tool notebook run waits.md > /tmp/out.raw
  go run ./stable < /tmp/out.raw > output.txt
  ```

  `stable` removes the timings and "kept" markers, the only parts that change
  from run to run; the notebook prints rounded numbers, so the output does not
  depend on the architecture. The `ci` workflow runs it again and fails if
  the output differs.

`hugo server` previews the page (Go is needed to resolve the kit).
