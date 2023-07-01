package utils

import "net"

func IsValidIPv4Address(ip string) bool {
	trial := net.ParseIP(ip)
	return trial.To4() != nil
}
