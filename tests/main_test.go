package tests

import (
	"encoding/json"
	"fmt"
	"github.com/kooksee/html_meta/internal/services"
	"github.com/kooksee/html_meta/internal/utils"
	"github.com/kooksee/html_meta/readability"
	"io/ioutil"
	"testing"
)

func TestAll(t *testing.T) {
	test, err := readability.NewReadability("https://mp.weixin.qq.com/s/7I8oDh6TyrFev7yvqqg4zA")
	if err != nil {
		fmt.Println("failed.", err)
		return
	}
	test.Parse()
	utils.P(test.Content)
	utils.P(test.Html)
}

func TestName123(t *testing.T) {
	d1, err := ioutil.ReadFile("../patterns/common.txt")
	utils.MustNotError(err)
	//dt,err:=services.GetMetadata("https://mp.weixin.qq.com/s/7I8oDh6TyrFev7yvqqg4zA", string(d1))
	//dt,err:=services.GetMetadata("https://news.mbalib.com/story/27003", string(d1))
	dt, err := services.GetMetadata("http://8btc.com/article-4772-1.html", string(d1))
	utils.MustNotError(err)
	dd, err := json.Marshal(dt)
	utils.MustNotError(err)
	fmt.Println(string(dd))
}
