package Analyse

import (
	"fmt"
	"github.com/opesun/goquery"
	"log"
	"strings"
	"solinx/sospider/Common"
	"solinx/sospider/DataStore"
)

/**
页面分析器
author：linx
date：2014-11-22
 */

type AnalysePage struct{
	Name string
}

func (self AnalysePage) AnalyseBody(body string) {

	//	fmt.Println(body)
	dataStore := DataStore.MongoStore{}

	p, err := goquery.ParseString(body)
	if err == nil {
		toutiao := p.Find(".toutiao")
		//		for i := 0; i < toutiao.length(); i++ {
		title := toutiao.Find("h2").Text()
		siteUrl := p.Find(".logo a").Attr("href")
		url := toutiao.Find("h2 a").Attr("href")
		//不包含Http：//则不是绝对路径
		isAbsolute := strings.Contains(url, "http://")
		reqBody := Common.HttpRequester{}
		reqUrl := ""
		if (isAbsolute) {
			reqUrl = url
		}else {
			reqUrl = siteUrl+url
		}
		postHtml := reqBody.RequestBody(reqUrl)
		//		post := reqBody.RequestBody(reqUrl)
		post := analysePost(postHtml)

		log.Println(url)
		log.Println(title)
		log.Println(siteUrl)
		log.Println(post)
		dataStore.PushMongoStore(title, post, url, siteUrl, isAbsolute)
		//
		//		}

	}else {
		fmt.Println(err)
	}
}

func analysePost(post string) (string) {
	body, err := goquery.ParseString(post)
	retVal := ""
	if err == nil {
		body.Find(".a_self pull-right btn btn-primary add-collect").Remove()

		retVal = body.Find(".neirong-box").Eq(0).Html()
	}
	return retVal
}

func (self AnalysePage) Analyse() {

}
