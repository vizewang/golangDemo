package core
import (
	"net/http"
	"io/ioutil"
)


func Get(url string)(content string,statusCode int)  {
	resp,err1:=http.Get(url)
	if err1!=nil{
		statusCode=-100
		return
	}
	defer resp.Body.Close()
	data,err2:=ioutil.ReadAll(resp.Body)
	if err2!=nil{
		statusCode=-200
		return
	}
	statusCode=resp.StatusCode
	content=string(data)
	return
}