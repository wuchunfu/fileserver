package iputil

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func GetIntranetIp() []string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	var ipArr = make([]string, 0)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ipStr := ipNet.IP.String()
				ipArr = append(ipArr, ipStr)
			}
		}
	}
	return ipArr
}

func IsPublicIP(ipStr string) bool {
	parseIp := net.ParseIP(ipStr)
	if parseIp.IsLoopback() || parseIp.IsLinkLocalMulticast() || parseIp.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := parseIp.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// 判断ip地址区间
func IpBetween(checkIp string) bool {
	ipStr := net.ParseIP(checkIp)
	fromIp := net.ParseIP("0.0.0.0")
	toIp := net.ParseIP("255.255.255.255")
	if fromIp == nil || toIp == nil || ipStr == nil {
		return false
	}
	from16 := fromIp.To16()
	to16 := toIp.To16()
	test16 := ipStr.To16()
	if from16 == nil || to16 == nil || test16 == nil {
		return false
	}
	if bytes.Compare(test16, from16) >= 0 && bytes.Compare(test16, to16) <= 0 {
		return true
	}
	return false
}

func inetNToA(ipNum int64) net.IP {
	var byteArr [4]byte
	byteArr[0] = byte(ipNum & 0xFF)
	byteArr[1] = byte((ipNum >> 8) & 0xFF)
	byteArr[2] = byte((ipNum >> 16) & 0xFF)
	byteArr[3] = byte((ipNum >> 24) & 0xFF)
	return net.IPv4(byteArr[3], byteArr[2], byteArr[1], byteArr[0])
}

func inetAToN(ipStr string) int64 {
	parseIp := net.ParseIP(ipStr)
	bits := strings.Split(parseIp.String(), ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)
	return sum
}
