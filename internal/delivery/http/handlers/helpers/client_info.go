package helpers

import (
	"net"
	"net/http"
	"strings"
)

type ClientInfo struct {
	IP        *string
	UserAgent *string
	DeviceID  *string
}

func ExtractClientInfo(r *http.Request) ClientInfo {
	return ClientInfo{
		IP:        extractIP(r),
		UserAgent: stringPtr(strings.TrimSpace(r.UserAgent())),
		DeviceID:  stringPtr(strings.TrimSpace(r.Header.Get("X-Device-ID"))),
	}
}

func extractIP(r *http.Request) *string {
	if ip := strings.TrimSpace(r.Header.Get("X-Real-IP")); ip != "" {
		return &ip
	}

	if xff := strings.TrimSpace(r.Header.Get("X-Forwarded-For")); xff != "" {
		parts := strings.Split(xff, ",")
		ip := strings.TrimSpace(parts[0])
		if ip != "" {
			return &ip
		}
	}

	host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err == nil && host != "" {
		return &host
	}

	remote := strings.TrimSpace(r.RemoteAddr)
	if remote != "" {
		return &remote
	}

	return nil
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
