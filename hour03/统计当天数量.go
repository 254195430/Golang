package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"regexp"
	"strconv"
	"sync"
)

var (
	chanImageUrls chan string
	chanTask      chan string
	waitGroup     sync.WaitGroup
	reDate        = `houseInfo.{40}1[2-9]\d`
)

func main() {
	chanImageUrls = make(chan string, 100000)
	chanTask = make(chan string, 16)

	for i := 1; i < 32; i++ {
		waitGroup.Add(1)
		//查看成交数		go getImgUrls("https://xa.lianjia.com/chengjiao/pg" + strconv.Itoa(i))
		// 查看二手房
		go getImgUrls("https://xa.lianjia.com/ershoufang/pg" + strconv.Itoa(i))
	}

	//	waitGroup.Add(1)
	waitGroup.Wait()
}

func getImgUrls(url string) {
	urls := GetImgs(url)
	for _, url = range urls {
		chanImageUrls <- url
	}
	chanTask <- url
	waitGroup.Done()
}
func GetImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reDate)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))
	fmt.Println(url)
	//    for _, result := range results{
	//   	fmt.Println(result)
	//	}

	return
}

func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	pageStr = string(pageBytes)
	//	fmt.Println(pageStr)
	return pageStr
}
