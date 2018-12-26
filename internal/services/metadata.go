package services

import (
	"encoding/json"
	"fmt"
	"github.com/kooksee/html_meta/internal/utils"
	"github.com/kooksee/html_meta/readability"
	"strings"
)

func GetMetadataByData(data string, pattern string) (map[string]interface{}, error) {
	red, err := readability.NewReadabilityByHtml(data)
	if err != nil {
		return nil, err
	}
	red.Parse()

	dt, err := utils.UnMashallHtml([]byte(red.Html), pattern)
	if err != nil {
		return nil, err
	}
	d1 := make(map[string]interface{})
	if err := json.Unmarshal(dt, &d1); err != nil {
		return nil, err
	}

	d1["content_1"] = red.Content

	dtd := make(map[string][]string)
	for k, v := range d1 {
		k = strings.Split(k, "_")[0]
		switch v1 := v.(type) {
		case string:
			dtd[k] = append(dtd[k], strings.TrimSpace(v1))
		case []interface{}:
			for _, v2 := range v1 {
				switch v3 := v2.(type) {
				case string:
					dtd[k] = append(dtd[k], strings.TrimSpace(v3))
				case map[string]interface{}:
					for _, v4 := range v3 {
						if v4 == nil || strings.TrimSpace(v4.(string)) == "" {
							continue
						}
						dtd[k] = append(dtd[k], strings.TrimSpace(v4.(string)))
					}
				}
			}
		}
	}

	//dddd, _ := json.Marshal(dtd)
	//fmt.Println("\n",string(dddd),"\n")
	//d3, _ := json.Marshal(dtd["img"])
	//fmt.Println(string(d3))

	d2 := make(map[string]interface{})
	for k, v := range dtd {
		if k == "img" {
			d2[k] = v
			continue
		}

		v2 := ""
		for _, v1 := range v {
			if v1 == "" {
				continue
			}

			if len(v1) > len(v2) {
				v2 = v1
			}
		}
		d2[k] = v2
	}

	return d2, nil
}

func GetMetadata(url string, pattern string) (map[string]interface{}, error) {

	red, err := readability.NewReadability(url)
	if err != nil {
		return nil, err
	}
	red.Parse()

	dt, err := utils.UnMashallHtml([]byte(red.Html), pattern)
	if err != nil {
		return nil, err
	}
	d1 := make(map[string]interface{})
	if err := json.Unmarshal(dt, &d1); err != nil {
		return nil, err
	}

	d1["content_1"] = red.Content

	dtd := make(map[string][]string)
	for k, v := range d1 {
		k = strings.Split(k, "_")[0]
		switch v1 := v.(type) {
		case string:
			dtd[k] = append(dtd[k], strings.TrimSpace(v1))
		case []interface{}:
			fmt.Println(k)
			for _, v2 := range v1 {
				dtd[k] = append(dtd[k], strings.TrimSpace(v2.(string)))
			}
		}
	}

	//d3, _ := json.Marshal(dtd["img"])
	//fmt.Println(string(d3))

	d2 := make(map[string]interface{})
	for k, v := range dtd {
		v2 := ""
		for _, v1 := range v {
			if v1 == "" {
				continue
			}

			if len(v1) > len(v2) {
				v2 = v1
			}
		}
		d2[k] = v2
	}

	return d2, nil
}
