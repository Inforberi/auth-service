package helpers

import "net/http"

type ClientInfo struct {
	IP        *string
	UserAgent *string
	DeviceID  *string
}

func ExtractClientInfo(r *http.Request) ClientInfo {

	var ipPtr *string
	var uaPtr *string
	var devicePtr *string

	// IP
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	if ip != "" {
		ipPtr = &ip
	}

	// User Agent
	ua := r.UserAgent()
	if ua != "" {
		uaPtr = &ua
	}

	// Device ID
	device := r.Header.Get("X-Device-ID")
	if device != "" {
		devicePtr = &device
	}

	return ClientInfo{
		IP:        ipPtr,
		UserAgent: uaPtr,
		DeviceID:  devicePtr,
	}
}
