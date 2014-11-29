package main

import "fmt"
import "gopkg.in/mgo.v2"

func main() {
	fmt.Printf("病毒!")
	 mgo.Dial("127.0.0.1:27017")


}
