package md_loader

import (
	"bufio"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yosssi/gohtml"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parseMarkdownComments(files []string) []DocumentationPage {
	pageList := make([]DocumentationPage, 0)
	for i := range files {
		fmt.Println("Loading markdown " + files[i])
		file, err := os.Open(files[i])
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)
		var page = DocumentationPage{}
		var contentFixed strings.Builder
		var first = true
		for scanner.Scan() {
			var line = scanner.Text()
			line = strings.ReplaceAll(line, ".md", ".html")

			if strings.Contains(line, "[//]: # (") {
				// it is a comment!
				if last := len(line) - 1; last >= 0 && line[last] == ')' {
					line = line[:last]
				}
				line = trim(line)

				var key = strings.Split(line, ":")[0]
				var value = strings.Split(line, ":")[1]

				switch key {
				case "TITLE":
					page.Title = value
					break
				case "DESCRIPTION":
					page.Description = value
					break
				case "ICON":
					page.Icon = value
					break
				case "TAGS":
					page.Tags = strings.Split(value, ",")
				case "COMMANDS":
					var builder []string
					builder = strings.SplitN(value, ",", 2)
					page.Commands = append(page.Commands, builder)

				}
			} else {
				if strings.HasPrefix(line, "#") && !first {
					contentFixed.Write([]byte("<br />"))
				}

				if len(line) > 5 {
					first = false
				}

				contentFixed.Write([]byte("\n"))
				contentFixed.Write([]byte(line))
			}
		}

		extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.OrderedListStart | parser.DefinitionLists | parser.Tables | parser.HardLineBreak | parser.LaxHTMLBlocks
		parser := parser.NewWithExtensions(extensions)
		maybeUnsafeHTML := markdown.ToHTML([]byte(contentFixed.String()), parser, nil)
		html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)
		page.Html = gohtml.Format(string(html))

		page.Html = strings.ReplaceAll(page.Html, "::warningstart::", `<div style="padding: 20px;
  background-color: #f44336;
  color: white;">`)

		page.Html = strings.ReplaceAll(page.Html, "::warningend::", `</div>`)

		filename := filepath.Base(files[i])
		page.Filename = filename

		pageList = append(pageList, page)

		file.Close()
	}
	return pageList
}

func trim(s string) string {
	if idx := strings.Index(s, "("); idx != -1 {
		return s[idx+1:]
	}
	return s
}
