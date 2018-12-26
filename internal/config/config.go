package config

import (
	"github.com/kooksee/html_meta/internal/utils"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
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
	pattern      []byte
}

func (t *config) GetPattern() []byte {
	return t.pattern
}

func (t *config) IsDebug() bool {
	return t.debug
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
	t.pattern, _ = ioutil.ReadFile("patterns/common.txt")
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
	})
	return cfg
}
