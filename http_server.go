package main

import (
    "fmt"
    "net/http"
    //"strings"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Server", "ngx_cloud")
    r.ParseForm()  //解析参数，默认是不会解析的
    /*fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
    fmt.Println("path: ", r.URL.Path)
    fmt.Println("scheme: ", r.URL.Scheme)
    fmt.Println("ID: ", r.Form["id"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }*/
    fmt.Fprintf(w, "Hello Go-Language!") //这个写入到w的是输出到客户端的
}

func HttpResponse(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Server", "ngx_cloud")    
    w.Write([]byte("<a href='http://192.168.137.129:8080/myhttpDNS'>video</a>"))
}

func HttpDNS(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Server", "ngx_cloud")
    w.Write([]byte("<a href='http://192.168.137.129/video_test/player_org.html'>ngx_cloud</a>\n"))
}



func main() {
	http_map := map[string](func(http.ResponseWriter, *http.Request)) {
		"/myhello":    SayHello,
		"/myhttpDNS":  HttpDNS,
		"/myhttp":     HttpResponse,
	}
	
    for location, http_handler := range http_map {
    	    http.HandleFunc(location, http_handler)
    	}
	
    err := http.ListenAndServe(":8080", nil)
	if err != nil {
	    fmt.Println("ListenAndServe error:", err.Error())
	}
}
