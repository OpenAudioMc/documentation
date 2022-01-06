package md_loader

type DocumentationPage struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Icon        string     `json:"icon"`
	Filename    string     `json:"path"`
	Tags        []string   `json:"tags"`
	Commands    [][]string `json:"commands"`
	Html        string     `json:"html"`
}
