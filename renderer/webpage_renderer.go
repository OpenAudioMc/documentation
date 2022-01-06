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
		renderer.RegisterComponent("documentation_command_list", getCommandListArticles())
		var output, _ = renderer.Render(context, readHtmlTest("html_templates/base.html"))
		var fn = "docs/" + filename
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
	out += getDocCard("technical", "Technical Notes", "Technical notes and requirements for server owners")

	return out
}
func getCommandListArticles() string {
	var out = ""

	out += getCommandList()

	return out
}
func getDocCard(topic string, title string, about string) string {
	var cards = `
<div class="container px-5 py-10 mx-auto border-b-2 border-indigo-500">
        <div class="flex flex-wrap w-full mb-10 flex-col items-center text-center hide_on_search">
            <h1 class="sm:text-3xl text-2xl font-medium title-font mb-2 text-white">` + title + `</h1>
            <p class="lg:w-1/2 w-full leading-relaxed text-opacity-80">` + about + `</p>
        </div>
        <div class="flex flex-wrap -m-4">`
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
<div class="search_cards xl:w-1/3 md:w-1/2 p-4 ">
<a href="` + strings.ReplaceAll(pages[i].Filename, ".md", ".html") + `">
	<div class="border border-gray-700 border-opacity-75 p-6 rounded-lg">
		<div class="flex flex-row items-stretch">
			<div class="xl:w-1/3 md:w-2/5 p-1 flex h-40">
				<div class=" mb-auto mt-auto  ">
					<i class="` + pages[i].Icon + ` fa-7x"></i>
				</div>
			</div>
			<div class="xl:w-2/3 md:w-3/5 p-4">
				<h2 class="text-lg text-white font-medium title-font mb-2 pl-5">` + pages[i].Title + `</h2>
				<p class="leading-relaxed text-base pl-5">` + pages[i].Description + `</p>
				
				<p class="search_tags leading-relaxed text-base pl-5 hidden">Tags: ` + strings.Join(pages[i].Tags, ", ") + `</p>

			</div>
		</div>
	</div>
</a>
</div>
`
	}

	cards += ` </div>
    </div>`

	return cards
}

func getCommandList() string {
	var cards = `
<div class="container px-5 py-10 mx-auto border-b-2 border-indigo-500">
        <div class="flex flex-wrap w-full mb-10 flex-col items-center text-center hide_on_search ">
			<table class="border-collapse border border-gray-400 bg-gray-800">
				<tr>
					<th class="border border-gray-300 w-3/12">Command</th>
					<th class="border border-gray-300">Description</th>
					<th class="border border-gray-300 w-3/12">More info  </th>
				</tr>`
	var pages = md_loader.LoadPages()
	for i := range pages {
		for y := range pages[i].Commands {
			cards += `
<tr>
<td class="border border-gray-300 search_column">` + pages[i].Commands[y][0] + `</td>
<td class="border border-gray-300 text-left pl-3 pt-4 pb-4 search_column">` + pages[i].Commands[y][1] + `</td>
<td class="border border-gray-300 text-left pl-3 ml-20 mr-20 text-blue-500"><a href="` + strings.ReplaceAll(pages[i].Filename, ".md", ".html") + `">Learn More...</a></td>
</tr>`
			cards += ``
		}
	}
	cards += `</table> </div>
    </div>`
	return cards
}
