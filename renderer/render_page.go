package renderer

import (
	"fmt"
	"github.com/Mindgamesnl/toothpaste"
	"github.com/OpenAudioMc/documentation/md_loader"
	"os"
	"strings"
)

func renderPage(page md_loader.DocumentationPage, tp *toothpaste.Renderer, docHtml string) {
	var context = toothpaste.NewRenderContext()
	context.SetVariable("content", page.Html)
	fmt.Println(docHtml)
	// render
	out, _ := tp.Render(context, docHtml)

	var fn = "out/" + strings.ReplaceAll(page.Filename, ".md", ".html")
	os.Remove(fn)
	os.WriteFile(fn, []byte(out), 0644)
}
