<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse: `memfs` Library
**URL** [multiverse-os.org](https://multiverse-os.org)

A pure-go dependency free memory filesystem (memfs) that abstracts a file system
ontop of go datatypes and allows developers to interact through an API modeled
after the Go File standard library.

Development plans include mounting the memory file system so it is accessible
and visible from terminal and provide easy hooking of functionality on
write/read and other operations for further layers of abstraction.

#### Usage 
```go
package main

import(
    "fmt"
    "os"

    "github.com/absfs/memfs"
)

func main() {
    fs, _ := memfs.NewFS() // remember kids don't ignore errors

    // Opens a file with read/write permissions in the current directory
    f, _ := fs.Create("/example.txt")

    f.Write([]byte("Hello, world!"))
    f.Close()

    fs.Remove("example.txt")
}
```

