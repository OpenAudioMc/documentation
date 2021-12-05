package renderer

import (
	"github.com/OpenAudioMc/documentation/md_loader"
	"io/ioutil"
)

func RenderPages(pages []md_loader.DocumentationPage) {
	for i := range pages {
		renderDocs(pages[i], pages[i].Html)
	}
}

func readHtmlTest(f string) string {
	content, _ := ioutil.ReadFile(f)
	return string(content)
}
