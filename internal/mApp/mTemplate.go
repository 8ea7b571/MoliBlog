package mApp

func (ma *MApp) loadTemplates() {
	assetsPath := "templates/default/assets"
	htmlPath := "templates/default/html/*.html"

	ma.engine.Static("/assets", assetsPath)
	ma.engine.LoadHTMLGlob(htmlPath)
}
