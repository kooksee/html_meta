package utils

import (
	"encoding/json"
	"fmt"
	"github.com/storyicon/graphquery"
)

func UnMashallHtml(data []byte, pattern string) ([]byte, error) {
	res := graphquery.ParseFromString(string(data), pattern)
	if len(res.Errors) != 0 {
		return nil, fmt.Errorf(res.Errors[0])
	}

	if res.Data == nil {
		return nil, fmt.Errorf("parse error")
	}

	dt, err := json.Marshal(res.Data)
	if err != nil {
		return nil, err
	}
	return dt, nil
}
