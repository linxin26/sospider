package DataStore

import (
	"log"
	"gopkg.in/mgo.v2"
	"fmt"
)

type MongoStore struct{
	Name string
}

const (
	Url = "lld:lld@localhost:27017"
)

type post struct{
	Title      string
	Body       string
	Url        string
	SiteUrl    string
	IsAbsolute bool
}

func (MongoStore) PushMongoStore(title string, body string, url string, siteUrl string, isabsolute bool) {

	session , err := mgo.Dial("127.0.0.1:27017")
	if err == nil {
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		spider := session.DB("blog").C("spider")
		err := spider.Insert(&post{title, body, url, siteUrl, isabsolute})
		if err != nil {
			log.Fatal(err)
		}


	}else {
		fmt.Println(err)
		panic(err)

	}

}
