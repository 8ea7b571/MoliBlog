package mApp

func (ma *MApp) loadRoutes() {
	ma.engine.GET("/", ma.IndexHandler)
	ma.engine.PUT("/update", ma.UpdateBlogHandler)
}
