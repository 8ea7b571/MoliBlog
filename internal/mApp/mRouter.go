package mApp

func (ma *MApp) loadRoutes() {
	ma.engine.GET("/", ma.IndexHandler)
	ma.engine.GET("/article/:hash", ma.ArticleHandler)
	ma.engine.PUT("/update", ma.UpdateArticleHandler)
}
