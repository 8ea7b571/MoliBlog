package mApp

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/8ea7b571/MoliBlog/internal/model"
	"github.com/8ea7b571/MoliBlog/utils"
	"gopkg.in/yaml.v3"
)

// loadMarkdownFiles load markdown source files
func (ma *MApp) loadMarkdownFiles() error {
	var markdownList []model.MFileInfo

	markdownPath := fmt.Sprintf("%s/%s", ma.Config.Root, SRC)
	err := filepath.Walk(markdownPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// filter markdown files
		if filepath.Ext(path) == ".md" {
			markdownList = append(markdownList, model.MFileInfo{
				Name: info.Name(),
				Path: path,
			})
		}

		return nil
	})

	if err != nil {
		return err
	}

	ma.markdownFiles = markdownList
	return nil
}

// parseMarkdowns parse markdown files to html
func (ma *MApp) parseMarkdowns() error {
	htmlPath := fmt.Sprintf("%s/%s", ma.Config.Root, DST)

	for _, file := range ma.markdownFiles {
		// read markdown file
		_mdFile, err := os.Open(file.Path)
		if err != nil {
			return err
		}
		defer _mdFile.Close()

		_mdByte, err := io.ReadAll(_mdFile)
		if err != nil {
			return err
		}

		// extract frontmatter and clean markdown bytes
		_frontmatter, _cleanbytes := utils.ExtractFrontMatter(_mdByte)

		// parse frontmatter to article metadata
		var article model.MArticle
		err = yaml.Unmarshal(_frontmatter, &article)
		if err != nil {
			return err
		}

		// convert into html format and save to html file
		_htmlPath := fmt.Sprintf("%s/%s.html", htmlPath, file.Name)
		_htmlFile, err := os.Create(_htmlPath)
		if err != nil {
			return err
		}
		defer _htmlFile.Close()

		_htmlbyte := ma.lute.Markdown(file.Name, _cleanbytes)
		_, err = _htmlFile.Write(_htmlbyte)
		if err != nil {
			return err
		}

		// save html path to article's metadata
		article.HtmlHash = utils.Sha256Hash(_htmlbyte)
		article.HtmlPath = _htmlPath

		ma.articles = append(ma.articles, &article)

		// save tags and categories to map
		for _, tag := range article.Tags {
			ma.tags[tag] = append(ma.tags[tag], &article)
		}

		for _, category := range article.Categories {
			ma.categories[category] = append(ma.categories[category], &article)
		}
	}

	// sort articles by date
	model.SortArticlesByDate(ma.articles)
	return nil
}
