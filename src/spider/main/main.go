package main
import (
	"fmt"
	"spider/core"
	"regexp"
)
var(
	ptnIndexItem=regexp.MustCompile(`<a target="_blank" href="(.+\.html)" title=".+" >(.+)</a>`)
)
type IndexItem struct {
	url string
	title string
}
func findIndex(content string)(index []IndexItem,err error)  {
	matches:=ptnIndexItem.
}

func main()  {
	fmt.Println("Get index...")
	s,statusCode:=core.Get("http://www.yifan100.com/dir/15136/");
if statusCode!=200{
	return
}
	index,_:=findIndex(s)

}
