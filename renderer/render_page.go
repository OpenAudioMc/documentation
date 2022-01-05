package renderer

import (
	"github.com/Mindgamesnl/toothpaste"
	"github.com/OpenAudioMc/documentation/md_loader"
	"os"
	"strings"
)

func renderDocs(page md_loader.DocumentationPage, docHtml string) {
	var renderer = toothpaste.NewRenderer()
	var context = toothpaste.NewRenderContext()
	// register global component in the renderer
	renderer.RegisterComponent("page_content", readHtmlTest("html_templates/doc_article.html"))
	renderer.RegisterComponent("documentation_content", docHtml)

	context.SetVariable("doc_title", page.Title)
	context.SetVariable("doc_about", page.Description)
	context.SetVariable("doc_icon", page.Icon)
	context.SetVariable("doc_tags", page.Tags)
	var output, _ = renderer.Render(context, readHtmlTest("html_templates/base.html"))
	var fn = "docs/" + strings.ReplaceAll(page.Filename, ".md", ".html")
	os.Remove(fn)
	os.WriteFile(fn, []byte(output), 0644)
}
