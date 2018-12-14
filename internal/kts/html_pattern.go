package kts

import (
	"github.com/kooksee/html_meta/internal/config"
	"gopkg.in/gorp.v1"
)

func Init() {
	hp := &HtmlPattern{}
	hp.getDb().AddTableWithName(*hp, hp.TableName())
}

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

func (t *HtmlPattern) GetPatternNames() (names []string, err error) {
	_, err = t.getDb().Select(&names, "select name from ?", t.TableName())
	return
}

func (t *HtmlPattern) GetPattern() error {
	pt := config.DefaultConfig().PatternGet(t.Name)
	if pt != nil {
		t.Pattern = pt.(string)
		return nil
	}

	if err := t.getDb().SelectOne(t, "select * from ? where name=?", t.TableName(), t.Name); err != nil {
		return err
	}

	config.DefaultConfig().PatternSet(t.Name, t.Pattern)
	return nil
}
