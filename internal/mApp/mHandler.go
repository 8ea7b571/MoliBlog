package mApp

import (
	"net/http"
	"strings"

	"github.com/8ea7b571/MoliBlog/internal/model"
	"github.com/8ea7b571/MoliBlog/utils"
	"github.com/gin-gonic/gin"
)

func (ma *MApp) IndexHandler(ctx *gin.Context) {
	// generate recent posts
	var recentPosts []*model.MPost
	for i := 0; i < utils.Min(len(ma.Posts), ma.Config.MSite.RecentPostNum); i++ {
		tmpPost := ma.Posts[i]
		tmpPost.Date = strings.Split(tmpPost.Date, " ")[0]

		recentPosts = append(recentPosts, tmpPost)
	}

	resData := gin.H{
		"site_info": gin.H{
			"title":  ma.Config.MSite.Title,
			"author": ma.Config.MSite.Author,
		},
		"recent_posts": recentPosts,
	}

	ctx.HTML(http.StatusOK, "index.html", resData)
}

func (ma *MApp) UpdateBlogHandler(ctx *gin.Context) {
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
