package helper

import "net/http"

type Global interface {
	GetIPAddress(request *http.Request)
}

func GetIPAddress(request *http.Request) string {
	ipAddress := request.Header.Get("X-Forwarded-For")
	if ipAddress == "" {
		ipAddress = request.RemoteAddr
	}
	return ipAddress
}
