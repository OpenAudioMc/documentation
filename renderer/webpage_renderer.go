package renderer

import (
	"github.com/Mindgamesnl/toothpaste"
	"os"
	"path/filepath"
	"strings"
)

func FindHtmlPages() []string {
	searchDir := "webpages/"

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(strings.ToLower(path), ".html") {
			fileList = append(fileList, path)
		}
		return err
	})

	if e != nil {
		panic(e)
	}

	return fileList
}

func RenderHtmlPages()  {
	var files = FindHtmlPages()

	for i := range files {
		var file = files[i]
		var filename = filepath.Base(file)
		var renderer = toothpaste.NewRenderer()
		var context = toothpaste.NewRenderContext()
		// register global component in the renderer
		renderer.RegisterComponent("page_content", readHtmlTest(file))
		var output, _ 	= renderer.Render(context, readHtmlTest("html_templates/base.html"))
		var fn = "out/" + filename
		os.Remove(fn)
		os.WriteFile(fn, []byte(output), 0644)
	}

}