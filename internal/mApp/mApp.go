package mApp

import (
	"fmt"
	"log"

	"github.com/88250/lute"
	"github.com/8ea7b571/MoliBlog/config"
	"github.com/8ea7b571/MoliBlog/internal/model"
	"github.com/gin-gonic/gin"
)

type MApp struct {
	Host   string
	Port   int
	Config *config.MConfig

	lute   *lute.Lute
	engine *gin.Engine

	tags       map[string][]*model.MPost // A collection of Posts with a certain tag
	categories map[string][]*model.MPost // A collection of Posts in a certain category

	Posts    []*model.MPost
	SrcFiles []model.MFileInfo
}

const (
	SRC = "markdowns/src"
	DST = "markdowns/dst"
)

func (ma *MApp) Run() {
	ma.loadRoutes()
	ma.loadTemplates()

	addr := fmt.Sprintf("%s:%d", ma.Host, ma.Port)
	err := ma.engine.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}

func NewMApp(cfg *config.MConfig) *MApp {
	return &MApp{
		Host:   cfg.Host,
		Port:   cfg.Port,
		Config: cfg,

		lute:   lute.New(),
		engine: gin.Default(),

		tags:       make(map[string][]*model.MPost),
		categories: make(map[string][]*model.MPost),
	}
}
