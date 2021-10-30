package ipx

import (
	"fmt"
	"testing"
)

func TestIpUtil(t *testing.T) {
	ip := GetIntranetIp()
	for _, ipStr := range ip {
		if len(ipStr) != 0 {
			fmt.Println(ipStr)
		}
	}
	fmt.Println(IsPublicIP("192.168.0.10"))
	fmt.Println(IpBetween("192.168.0.10"))
	fmt.Println(inetAToN("192.168.0.10"))
	fmt.Println(inetNToA(inetAToN("192.168.0.10")))
}
