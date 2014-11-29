package ConfigCenter

import (
	"fmt"
	"os"
	"strings"
	"io"
)

/**
配置中心的Site配置文件
 */

type SiteConfig struct{

}

func main() {

	url := SiteConfig{}
	fmt.Println(url.ReadSite())

}

/**
author:linx
date:2014-11-20
解析url
 */
func (SiteConfig) ReadUrl() ([]string) {
	url := SiteConfig{}
	sites, _ := url.ReadSite()

	urls := strings.Split(sites, ",")


	return urls
}

/**
author:linx
date:2014-11-20
从urlFile文件读取url
 */
func (SiteConfig) ReadSite() (string, error) {
	//文件完整路径

	f, err := os.Open("./src/solinx/sospider/ConfigCenter/urlFile.txt")
	//文件相对路径

	buf := make([]byte, 100)
	var result  []byte
	defer f.Close()
	if err == nil {
		//读取文件中数据
		for {
			n, err := f.Read(buf[0:])

			result = append(result, buf[0:n]...)
			if err != nil {
				if err == io.EOF {
					break
				}
				return "", err
			}
		}

	}
	return string(result), err
}

