package config

import (
	"database/sql"
	"github.com/kooksee/html_meta/internal/utils"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/gorp.v1"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	debug        bool
	patternCache *cache.Cache
	id           string
	cacheTime    time.Duration
	mysqlDb      *gorp.DbMap
	mysqlUrl     string
}

func (t *config) GetMysql() *gorp.DbMap {
	if t.mysqlDb == nil {
		panic("please init mysql db")
	}
	return t.mysqlDb
}

func (t *config) IsDebug() bool {
	return t.debug
}

func (t *config) PatternGet(name string) interface{} {
	dt, b := t.patternCache.Get(name)
	if b {
		return dt
	}
	return nil
}

func (t *config) PatternSet(name string, dt interface{}) {
	t.patternCache.SetDefault(name, dt)
}

func (t *config) Init() {
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if !t.debug {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	t.id = utils.IpAddress()
	if t.id == "" {
		panic("获取不到ip地址")
	}

	log.Logger = log.
		Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().
		Str("service_name", "mworker").
		Str("service_ip", t.id).
		Str("service_id", t.id).
		Bool("is_debug", t.debug).
		Caller().
		Logger()

	t.patternCache = cache.New(t.cacheTime, t.cacheTime*2)

	log.Debug().Msg("init mysql")
	db, err := sql.Open("mysql", t.mysqlUrl)
	utils.MustNotError(err)

	t.mysqlDb = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
}

var cfg *config
var once sync.Once

func DefaultConfig() *config {
	once.Do(func() {
		cfg = &config{
			debug:     true,
			cacheTime: time.Minute * 10,
		}

		if e := env("DEBUG"); e != "" {
			cfg.debug = e == "true"
		}

		if e := env("mysql_url"); e != "" {
			cfg.mysqlUrl = e
		}

	})
	return cfg
}
