package kts

import (
	"github.com/kooksee/html_meta/internal/config"
	"gopkg.in/gorp.v1"
)

func NewHtmlPattern() *HtmlPattern {
	return &HtmlPattern{}
}

type HtmlPattern struct {
	Name    string `json:"name" db:"name"`
	Pattern string `json:"pattern" db:"pattern"`
}

func (t *HtmlPattern) TableName() string {
	return "html_pattern"
}

func (t *HtmlPattern) getDb() *gorp.DbMap {
	return config.DefaultConfig().GetMysql()
}

func (t *HtmlPattern) Save() error {
	return t.getDb().Insert(t)
}

func (t *HtmlPattern) GetPattern() error {
	return t.getDb().SelectOne(t, "select * from ? where name=?", t.TableName(), t.Name)
}
