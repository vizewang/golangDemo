package main
import (
	"fmt"
	"spider/core"
	"regexp"
	"io/ioutil"
)
var(
	ptnIndexItem=regexp.MustCompile(`<a target="_blank" href="(.+\.html)" title=".+" >(.+)</a>`)
	ptnContentRough=regexp.MustCompile(`(?s).*<div class="artcontent">(.*)<div id="zhanwei">.*`)
	ptnBrTag=regexp.MustCompile(`<br>`)
	ptnHTMLTag=regexp.MustCompile(`(?s)</?.*?>`)
	ptnSpace=regexp.MustCompile(`(^\s+)|()`)
)
type IndexItem struct {
	url string
	title string
}
func findIndex(content string)(index []IndexItem,err error)  {
	matches:=ptnIndexItem.FindAllStringSubmatch(content,10000)
	index=make([]IndexItem,len(matches))
	for i,item:=range matches{
		index[i]=IndexItem{"http://www.yifan100.com"+item[1],item[2]}
	}
	return
}
func readContent(url string)(content string)  {
	raw,statusCode:=core.Get(url)
	if statusCode!=200 {
		fmt.Print("Fail to get the raw data from",url,"\n")
		return
	}
	match:=ptnContentRough.FindStringSubmatch(raw)
	if match!=nil{
		content=match[1]
	}else{
		return
	}
	content=ptnBrTag.ReplaceAllString(content,"\r\n")
	content=ptnHTMLTag.ReplaceAllString(content,"")
	content=ptnSpace.ReplaceAllString(content,"")
	return
}
func main()  {
	fmt.Println("Get index...")
	s,statusCode:=core.Get("http://www.yifan100.com/dir/15136/");
if statusCode!=200{
	return
}
	index,_:=findIndex(s)
	fmt.Println(`Get contents and write to file...`)
	for _,itm :=range index{
		fmt.Printf("Get content %s from %s and write to file.\n",itm.title,itm.url)
		fileName:=fmt.Sprintf("c:\\%s.txt",itm.title)
		content:=readContent(itm.url)
		ioutil.WriteFile(fileName,[]byte(content),0644)
		fmt.Printf("Finish writing to %s.\n",fileName)
	}

}
