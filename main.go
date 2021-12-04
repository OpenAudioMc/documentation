package main

import (
	"fmt"
	"github.com/OpenAudioMc/documentation/md_loader"
)

func main() {
	var pageMeta = md_loader.LoadPages()
	fmt.Println(pageMeta)
}
