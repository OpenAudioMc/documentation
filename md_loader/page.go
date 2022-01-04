package md_loader

type DocumentationPage struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	Filename    string   `json:"path"`
	Tags        []string `json:"tags"`
	Html        string   `json:"html"`
}
