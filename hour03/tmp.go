package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
     re := regexp.MustCompile(reStr)：根据匹配规则，返回一个正则对象
     ret := re.FindAllStringSubmatch(srcStr,-1)：src是大字符串，-1是取所有
*/

var (
	reQQEmail = `(\d+)@qq.com`
	reEmail   = `\w+@\w+\.\w+(\.\w+)?`
	reLink    = `href="(https?://[\s\S]+?)"`
	rePhone   = `houseInfo.{40}\d{3}`
	//410222198611270512
	reIdcard = `[123456]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg    = `(https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif))))`
)

func main() {
	//1.爬邮箱
	//GetEmail()
	//2.方法抽取
	//GetEmail2("http://tieba.baidu.com/p/2544042204")
	//3.爬超链接
	//GetLink("http://www.baidu.com/s?wd=岛国%20留下邮箱")
	//4.爬手机号
	GetPhone("https://xa.lianjia.com/ershoufang/pg2/")
	//5.身份证号
	//GetIdcard("http://henan.qq.com/a/20171107/069413.htm")
	//6.超链接
	//	GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
}

//爬邮箱
func GetEmail() {
	//1.爬所有数据
	resp, err := http.Get("http://tieba.baidu.com/p/2544042204")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	//接收当前页面的数据
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.Read")
	fmt.Println(string(pageBytes))
	//2.取数据，通过正则
	pageStr := string(pageBytes)
	re := regexp.MustCompile(reQQEmail)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	//遍历数组
	for _, result := range ret {
		//fmt.Println(result)
		fmt.Printf("email=%s qq=%s\n", result[0], result[1])
	}
}

//处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

//抽取的爬邮箱
func GetEmail2(url string) {
	//获取页面数据
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range ret {
		fmt.Println(result)
	}
}

//爬页面所有数据
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	//接收当前页面的数据
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.Read")
	pageStr = string(pageBytes)
	return pageStr
}

//爬超链接
func GetLink(url string) {
	//获取页面数据
	pageStr := GetPageStr(url)
	fmt.Println(pageStr)
	re := regexp.MustCompile(reLink)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果：\n", len(ret))
	for _, result := range ret {
		fmt.Println(result)
	}
}

//爬手机号
func GetPhone(url string) {
	//获取页面数据
	pageStr := GetPageStr(url)
	fmt.Println(pageStr)
	re := regexp.MustCompile(rePhone)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果：\n", len(ret))
	for _, result := range ret {
		fmt.Println(result)
	}
}

//爬手机号
func GetIdcard(url string) {
	//获取页面数据
	pageStr := GetPageStr(url)
	fmt.Println(pageStr)
	re := regexp.MustCompile(reIdcard)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果：\n", len(ret))
	for _, result := range ret {
		fmt.Println(result)
	}
}

//爬图片链接
func GetImg(url string) {
	//获取页面数据
	pageStr := GetPageStr(url)
	fmt.Println(pageStr)
	re := regexp.MustCompile(reImg)
	ret := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果：\n", len(ret))
	for _, result := range ret {
		fmt.Println(result)
	}
}
