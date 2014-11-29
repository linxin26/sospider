package Common

import (
	"io/ioutil"
	"net/http"
	"github.com/axgle/mahonia"
)

type HttpRequester struct{
	Name string
}

func (HttpRequester) RequestBody(url string) (string) {

	resp , err := http.Get(url)
	retVal := ""
	if err == nil {
		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll((resp.Body))
			if err == nil {
				str := string(body)
				enc := mahonia.NewDecoder("utf-8")
				retVal = enc.ConvertString(str)
			}
		}
	}
	return retVal


}



