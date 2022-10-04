package http

import (
	"log"
	"net"
	"net/http"
)

// Get request IP address
func IP(req *http.Request) string {
	// req.RemoteAddr <- 192.168.131.229:49159
	// ip <- 19.168.131.229
	ip, _, err := net.SplitHostPort(req.RemoteAddr)

	// エラーチェック
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return string(ip)
}
