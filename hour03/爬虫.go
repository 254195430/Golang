package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	chanImageUrls chan string
	chanTask      chan string
	waitGroup     sync.WaitGroup
)

var (
	reMail = `(\w+)@\w+\.\w+(\.\w)?`
	reLink = `href="(https?://[\s\S]+?)"`
	//bak	reLink = `href="(https?://[\s\S]+?)"`
	//	reLink2 = `("http://[\s\S]+?")`
	rePhone = `1[2-8]\d{9}`
	//	reIDcard =
	rePhoto = `(https?://[^"]+?(\.((jpg)|(png))))`
	reImg   = `(https://[^"]+?(\.((jpg))))`
)

func main() {
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 16)

	for i := 1; i < 4; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://xa.lianjia.com/chengjiao/pg" + strconv.Itoa(i))
	}

	waitGroup.Add(1)
	go checkOK()
	for i := 0; i < 2; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()

	//	getEmail()
	//    GetPageHome("https://xa.lianjia.com/chengjiao/pg1/")
	// go  GetLink()
	//    GetPhoneNumber("https://www.zhaohaowang.com/")

	//	GetImage("http://desk.zol.com.cn/youxi/wangzherongyao/")
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
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))

	for _, result := range results {
		url := result[1]
		fmt.Println(url)
		urls = append(urls, url)
	}
	return
}

func checkOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成爬取任务\n", url)
		count++
		if count == 16 {
			close(chanImageUrls)
			break
		}

	}
	waitGroup.Done()
}

func DownloadImg() {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}

func GetFilenameFromUrl(url string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex+1:]
	timePrefix := strconv.Itoa(int(time.Now().Unix()))
	filename = timePrefix + "_" + filename
	return
}

func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	if err != nil {
		HandleError(err, "http.Get")
		return
	}
	defer resp.Body.Close()
	fBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil")
	err = ioutil.WriteFile("D:/Golang/src/test"+filename, fBytes, 644)
	HandleError(err, "write")
	if err != nil {
		return false
	} else {
		return true
	}
}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func getEmail() {
	resp, err := http.Get("http://tieba.baidu.com/p/6205999330")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(string(pageBytes))
	pageStr := string(pageBytes)
	re := regexp.MustCompile(reMail)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	for i, result := range ret {
		fmt.Printf("%d Email=%s\n", i+1, result[0])
	}
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

func getEmail2() {
	pageStr := GetPageStr("http://tieba.baidu.com/p/6205999330")
	re := regexp.MustCompile(reMail)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	for i, result := range ret {
		fmt.Printf("%d Email=%s\n", i+1, result[0])
	}
}

func GetLink() {
	pagestr := GetPageStr("http://www.downza.cn/")
	re := regexp.MustCompile(reLink)
	ret := re.FindAllStringSubmatch(pagestr, -1)
	for i, result := range ret {
		fmt.Printf("%d 超链接地址是=%s \n", i+1, result[1])
	}
	//    re :=regexp.MustCompile(reLink)
	//    ret := re.FindAllString(pagestr,-1)
	//	re2 := regexp.MustCompile(reLink2)
	//	for i, result := range ret {
	//		ret2 := re2.FindAllStringSubmatch(result[:],-1)
	//		fmt.Printf("%d 超链接地址是=%v \n", i+1, ret2)
	//	}
}

func GetPhoneNumber(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range ret {
		fmt.Println(result)
	}
}

func GetImage(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhoto)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d张图片\n", len(ret))
	for _, result := range ret {
		fmt.Println(result[1])
	}

}

func GetPageHome(url string) {
	pagestr := GetPageStr(url)
	re := regexp.MustCompile(reLink)
	ret := re.FindAllStringSubmatch(pagestr, -1)
	for i, result := range ret {
		fmt.Printf("%d 超链接地址是=%s \n", i+1, result[1])
	}
}
