package renderer

import (
	"github.com/Mindgamesnl/toothpaste"
	"github.com/OpenAudioMc/documentation/md_loader"
	"io/ioutil"
)

var tpRenderer = toothpaste.NewRenderer()

func prepareToothpaste()  {
	tpRenderer.RegisterComponent("doc_page", readHtmlTest("../html_templates/doc.html"))
}

func RenderPages(pages []md_loader.DocumentationPage) {
	prepareToothpaste()
	docPage := readHtmlTest("html_templates/doc.html")
	tpRenderer.RegisterComponent("doc_page", docPage)

	// render pages
	for i := range pages {
		renderPage(pages[i], tpRenderer, docPage)
	}
}

func readHtmlTest(f string) string {
	content, _ := ioutil.ReadFile(f)
	return string(content)
}
