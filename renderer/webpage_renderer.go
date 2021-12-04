package renderer

import (
	"github.com/Mindgamesnl/toothpaste"
	"github.com/OpenAudioMc/documentation/md_loader"
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

func RenderHtmlPages() {
	var files = FindHtmlPages()

	// build html for documentation index lists

	for i := range files {
		var file = files[i]
		var filename = filepath.Base(file)
		var renderer = toothpaste.NewRenderer()
		var context = toothpaste.NewRenderContext()
		// register global component in the renderer
		renderer.RegisterComponent("page_content", readHtmlTest(file))
		renderer.RegisterComponent("documentation_article_cards", getDocumentationArticles())
		var output, _ = renderer.Render(context, readHtmlTest("html_templates/base.html"))
		var fn = "out/" + filename
		os.Remove(fn)
		os.WriteFile(fn, []byte(output), 0644)
	}
}

func getDocumentationArticles() string {
	var out = ""

	out += getDocCard("installation", "Installation", "All the articles guiding you through the OpenAudioMc plugin installation and configuration")
	out += getDocCard("voicechat", "Proximity Voice Chat", "How to talk and interact with other players, and how to moderate it as a server owner")
	out += getDocCard("media", "Media", "Controlling media playback and effects")
	out += getDocCard("accounts", "Administration Accounts", "How to setup, use, and link Craftmend accounts to get all the best that OpenAudioMc has to offer")
	out += getDocCard("commands", "Utility Commands", "How to use some of the coolest Utility commands to take full advantage of the plugin")

	return out
}

func getDocCard(topic string, title string, about string) string {
	var cards = `
<div class="container px-5 py-10 mx-auto border-b-2 border-indigo-500">
        <div class="flex flex-wrap w-full mb-10 flex-col items-center text-center">
            <h1 class="sm:text-3xl text-2xl font-medium title-font mb-2 text-white">` + title + `</h1>
            <p class="lg:w-1/2 w-full leading-relaxed text-opacity-80">` + about + `</p>
        </div>
        <div class="flex flex-wrap -m-4">
`
	var pages = md_loader.LoadPages()
	for i := range pages {

		var hasTag = false
		for i2 := range pages[i].Tags {
			if pages[i].Tags[i2] == topic {
				hasTag = true
			}
		}

		if !hasTag {
			continue
		}

		if pages[i].Title == "" {
			continue
		}
		cards += `
<div class="xl:w-1/3 md:w-1/2 p-4">
	<div class="border border-gray-700 border-opacity-75 p-6 rounded-lg">
		<h2 class="text-lg text-white font-medium title-font mb-2">` + pages[i].Title + `</h2>
		<p class="leading-relaxed text-base">` + pages[i].Description + `</p>
	</div>
</div>
`
	}

	cards += ` </div>
    </div>`
	return cards
}
