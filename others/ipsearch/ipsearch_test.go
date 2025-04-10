package ipaddr

import (
	"fmt"
	"strings"
	"testing"
)

/**
 * @author xiao.luo
 * @description This is the unit test for IpSearch
 */

func TestLoad(t *testing.T) {
	fmt.Println("Test Load IP Dat ...")
	p, err := New("./ip.dat")
	if len(p.data) <= 0 || err != nil {
		t.Fatal("the IP Dat did not loaded successfully!")
	}
}

func TestGet(t *testing.T) {
	fmt.Println("Test Get IP ...")
	p, _ := New("./ip.dat")
	ip := "210.51.200.123"
	ipstr := p.Get(ip)
	fmt.Println(ipstr)
	if ipstr != `亚洲|中国|湖北|武汉||中国联通|420100|China|CN|114.298572|30.584355` {
		t.Fatal("the IP convert by ipSearch component is not correct!")
	}
}

func TestGet1(t *testing.T) {
	ip := "36.149.36.0"
	p, _ := New("./ip.dat")
	ipstr := p.Get(ip)
	fmt.Println(ipstr)
	// 亚洲|中国|江苏|扬州||中国移动|321000|China|CN|119.421003|32.393159
}

func TestLocal(t *testing.T) {
	// ip := "172.16.10.139" // |保留||||专用网络/局域网|||||
	ip := "127.0.0.1" // |保留||||本地环回|||||
	p, _ := New("./ip.dat")
	ipstr := p.Get(ip)
	fmt.Println(ipstr)
	fmt.Printf("%v\n", strings.Split(ipstr, "|"))
	for k, v := range strings.Split(ipstr, "|") {
		fmt.Printf("%v=>:%v.\n", k, v)
	}
}
