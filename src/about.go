package main

func (r *Render) RenderAbout() error {
	content, err := readFile("./data/about.md")
	aboutPage := r.NewPage("about.html", "", content)

	aboutPage.Title = r.Ego.Config.Site.Title
	aboutPage.Params["about"] = ParseMarkdown(string(content))

	_, err = aboutPage.Render()
	if err != nil {
		return err
	}
	aboutPage.URL = "/about.html"
	return aboutPage.Save()
}
