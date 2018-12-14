package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
	"math/big"
	"reflect"
	"regexp"
)

func MustNotError(err error) {
	if err != nil {
		log.Error().Err(err).Str("method", "MustNotError").Msg("error")
		panic(err.Error())
	}
}

func P(d ... interface{}) {
	for _, i := range d {
		dt, err := json.MarshalIndent(i, "", "\t")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
	}
}

func ParseOssUrl(url string) *oss.Config {
	c := &oss.Config{}
	dt := regexp.MustCompile(`oss://(?P<username>.*):(?P<password>.*)@(?P<host>.*)`).FindStringSubmatch(url)
	if len(dt) == 0 {
		panic(fmt.Sprintf("url %s parse error", url))
	}

	c.AccessKeyID = dt[1]
	c.AccessKeySecret = dt[2]
	c.Endpoint = dt[3]
	return c
}

func ParseRedisUrl(url string) *redis.Options {
	c := &redis.Options{}
	dt := regexp.MustCompile(`redis://:(?P<password>.*)@(?P<host>.*)/(?P<db>.*)`).FindStringSubmatch(url)
	if len(dt) == 0 {
		panic(fmt.Sprintf("url %s parse error", url))
	}

	c.Password = dt[1]
	c.Addr = dt[2]

	db, _ := big.NewInt(0).SetString(dt[3], 10)
	c.DB = int(db.Int64())
	return c
}
