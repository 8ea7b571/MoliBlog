package mApp

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (ma *MApp) IndexHandler(ctx *gin.Context) {
	resData := gin.H{
		"site_info": gin.H{
			"title": "MoliBlog",
		},
		"articles": ma.articles,
	}

	ctx.HTML(http.StatusOK, "index.html", resData)
}

func (ma *MApp) ArticleHandler(ctx *gin.Context) {
	hash := ctx.Param("hash")

	for _, article := range ma.articles {
		if article.HtmlHash == hash {
			file, err := os.Open(article.HtmlPath)
			if err != nil {
				_ = ctx.Error(err)
				return
			}
			defer file.Close()

			data, err := io.ReadAll(file)
			if err != nil {
				_ = ctx.Error(err)
				return
			}

			resData := gin.H{
				"site_info": gin.H{
					"title": "MoliBlog",
				},
				"article": gin.H{
					"title":      article.Title,
					"date":       article.Date,
					"tags":       article.Tags,
					"categories": article.Categories,
					"content":    template.HTML(data),
				},
			}

			ctx.HTML(http.StatusOK, "article.html", resData)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"msg": "not found"})
}

func (ma *MApp) UpdateArticleHandler(ctx *gin.Context) {
	var err error

	err = ma.loadMarkdownFiles()
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = ma.parseMarkdowns()
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
