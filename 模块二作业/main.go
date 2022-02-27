package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/healthz", Index)
	_ = http.ListenAndServe(":80", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {
	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		for _, val := range v {
			w.Header().Set(k, val)
		}
	}

	// 2. 设置环境变量
	os.Setenv("VERSION", "go1.15.5")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	ip := getClientIp(r)
	fmt.Println("addr: ", ip, "code: ", 200)
	// 当访问 localhost/healthz 时，应返回 200
	w.WriteHeader(200)
}

func getClientIp(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
