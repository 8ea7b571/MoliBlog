package model

import (
	"sort"
	"time"
)

// MPost post metadata
type MPost struct {
	Title      string   `yaml:"title" json:"title"`
	Cover      string   `yaml:"cover" json:"cover"`
	Date       string   `yaml:"date" json:"date"`
	Tags       []string `yaml:"tags" json:"tags"`
	Categories []string `yaml:"categories" json:"categories"`

	HtmlHash string `yaml:"htmlHash" json:"html_hash"`
	HtmlPath string `yaml:"htmlPath" json:"html_path"`
}

type MPostSlice []*MPost

func (a MPostSlice) Len() int {
	return len(a)
}

func (a MPostSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a MPostSlice) Less(i, j int) bool {
	timeFormat := "2006-01-02 15:04:05"
	t1, err1 := time.Parse(timeFormat, a[i].Date)
	t2, err2 := time.Parse(timeFormat, a[j].Date)

	if err1 != nil || err2 != nil {
		return false
	}

	return t1.After(t2)
}

func SortPostsByDate(Posts []*MPost) []*MPost {
	sort.Sort(MPostSlice(Posts))
	return Posts
}
