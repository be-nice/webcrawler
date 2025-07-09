package main

import (
	"crawly/pkg"
	"fmt"
)

func main() {
	fmt.Println(pkg.NormalizeURL("http://blog.boot.dev/path/"))
}
