package renderer

import (
	"github.com/Mindgamesnl/toothpaste"
	"github.com/OpenAudioMc/documentation/md_loader"
	"io/ioutil"
)

var tpRenderer = toothpaste.NewRenderer()

func prepareToothpaste()  {

}

func RenderPages(pages []md_loader.DocumentationPage) {
	tpRenderer.RegisterComponent("")
}

func readHtmlTest(f string) string {
	content, _ := ioutil.ReadFile(f)
	return string(content)
}
