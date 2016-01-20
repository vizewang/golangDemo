package main
import (
	"net/http"
	"fmt"
"strings"
	"html/template"
	"log"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

func sayhelloName(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello astaxie")
}

func login1(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("method:",r.Method)
	if r.Method=="GET"{
		t,_:=template.ParseFiles("login.gtpl")
		t.Execute(w,nil)
	}else {
		r.ParseForm()
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
	}
}

func login2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:",r.Method)
	if r.Method=="GET"{
		crutime:=time.Now().Unix()
		h:=md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token:=fmt.Sprintf("%x",h.Sum(nil))
		t,_:=template.ParseFiles("login.gtpl")
		t.Execute(w,token)
	}else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("token is "+token)//验证token的合法性
		} else {
			fmt.Println("there is no token")
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

func main() {
	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login",login2)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		log.Fatal("ListenAndServe:",err)
	}
}
