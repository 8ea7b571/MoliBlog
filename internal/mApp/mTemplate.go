package mApp

import (
	"fmt"
)

func (ma *MApp) loadTemplates() {
	assetsPath := fmt.Sprintf("%s/templates/default/assets", ma.Config.Root)
	htmlPath := fmt.Sprintf("%s/templates/default/html/*.html", ma.Config.Root)

	ma.engine.Static("/assets", assetsPath)
	ma.engine.LoadHTMLGlob(htmlPath)
}
