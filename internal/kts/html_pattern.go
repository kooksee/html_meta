package kts

import (
	"database/sql"
	"fmt"
	"github.com/kooksee/html_meta/internal/config"
	"github.com/rs/zerolog/log"
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
	Id      string `json:"id" db:"id"`
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
	_, err = t.getDb().Select(&names, fmt.Sprintf("select name from %s", t.TableName()))
	return
}

func (t *HtmlPattern) Update() error {
	_, err := t.getDb().Exec(fmt.Sprintf("update %s set pattern=? where name=?", t.TableName()), t.Pattern, t.Name)
	return err
}

func (t *HtmlPattern) Exist() bool {
	i, err := t.getDb().SelectInt(fmt.Sprintf("select count(*) from %s where name=?", t.TableName()), t.Name)
	if err == sql.ErrNoRows {
		return false
	}

	if err != nil {
		log.Error().Err(err).Msg("debug")
		return false
	}

	return i != 0
}

func (t *HtmlPattern) GetPattern() error {
	pt := config.DefaultConfig().PatternGet(t.Name)
	if pt != nil {
		t.Pattern = pt.(string)
		return nil
	}

	if err := t.getDb().SelectOne(t, fmt.Sprintf("select * from %s where name=?", t.TableName()), t.Name); err != nil {
		return err
	}

	config.DefaultConfig().PatternSet(t.Name, t.Pattern)
	return nil
}
