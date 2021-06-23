 [![go-doc](https://godoc.org/github.com/holoplot/go-bmap?status.svg)](https://godoc.org/github.com/holoplot/go-bmap)

# bmap sparse file writer implementation, written in Go

go-bmap is a pure Go implementation of a reader of the [bmap file format](https://github.com/intel/bmap-tools).
It can be used to efficiently write sparse file images to disk when the input files are transported by mechanisms
or file-systems that are unaware of holes in block allocations.

For more information on the file format and use cases, please refer to the reference implementation linked to above.

# Installation

Install the package like this:

```
go get github.com/holoplot/go-bmap/pkg/bmap
```

And then use it in your source code.

```
import "github.com/holoplot/go-bmap/pkg/bmap"
```

# Example

For a standalone example of this package, check out the code in [cmd/bmaptool](./cmd/bmaptool/main.go).

```
go run ./cmd/bmaptool/main.go -bmap test/data/input.bmap -input test/data/input.bz2 -output out
diff out test/data/output.bin
```

# License

MIT
