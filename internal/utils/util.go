package utils

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"reflect"
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
