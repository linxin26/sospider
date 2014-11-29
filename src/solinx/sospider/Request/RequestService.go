package Request

import (
	"fmt"
	"strconv"
	"solinx/sospider/ConfigCenter"
	"net/http"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"sync"
	"solinx/sospider/Analyse"
)

/**
Request模块入口
auther:linx
date:2014-11-20
 */

type HttpRequest struct{
	Name string
}

var wg2 sync.WaitGroup

/**
请求网页
author:linx
date:2014-11-20
从配置中心模块读取站点site，在for中启用go并发请求url
 */
func (HttpRequest) Requsturl() {

	//从配置中心取得站点
	site := ConfigCenter.SiteConfig{}
	urls := site.ReadUrl()

	for _, url := range urls {
		fmt.Printf("----------------%s-----------", url+"--")
		fmt.Println("")
		wg2.Add(1)
		go RequestPager(url)
	}
	wg2.Wait()
}

/**
请求内容
 */
func (HttpRequest) RequestBody(url string) (string) {
	resp , err := http.Get(url)
	retVal := ""
	if err == nil {
		if resp.StatusCode == 200 {
//			body, err := ioutil.ReadAll((resp.Body))
//			if err == nil {
////				retVal = body
//			}
		}
	}
	return retVal
}

/**
发起请求
author:linx
date:2014-11-20
发起请求
 */
func RequestPager(url string) {
	fmt.Println("go Request")
	//页面分析器
	analysePage := Analyse.AnalysePage{"analyse"}

	resp, err := http.Get(url)
	if err == nil {
		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				str := string(body)
				enc := mahonia.NewDecoder("utf-8")
				analysePage.AnalyseBody(enc.ConvertString(str))

			}
		}
	}else {
		fmt.Println(err)
	}
	defer wg2.Done()
}

func decode(body []byte) {
	json := ""
	html := ""
	for _, r := range body {
		rint := int(r)
		if rint < 128 {
			json +=string(r)
			html +=string(r)
		}else {
			json +="\\u"+strconv.FormatInt(int64(rint), 32)
			html +="&#"+strconv.Itoa(int(r))+";"
		}
	}
	fmt.Printf("JSON:%s", json)
	fmt.Printf("HTML:%s", html)
}


