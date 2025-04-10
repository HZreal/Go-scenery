package ipaddr

import (
	"log"
	"os"
	"strconv"
	"strings"
)

/**
 * @author xiao.luo
 * @description This is the go version for IpSearch
https://github.com/zengzhan/qqzeng-ip/blob/master/qqzeng-ip-dat
 高性能IP数据库格式详解 qqzeng-ip.dat

编码：UTF8和GB2312  字节序：Little-Endian

返回多个字段信息（如：亚洲|中国|香港|九龙|油尖旺|新世界电讯|810200|Hong Kong|HK|114.17495|22.327115）


------------------------ 文件结构 3.0 版 -------------------------
 //文件头  4字节
 [IP段记录]

//前缀区   8字节(4-4)   256*8
[索引区start第几个][索引区end第几个]


//索引区    8字节(4-3-1)   ip段行数*8
[结束IP数字][地区流位置][流长度]


 //内容区    长度无限制
[地区信息][地区信息]……唯一不重复

------------------------ 文件结构 3.0 版 ---------------------------

优势：压缩形式将数据存储在内存中，通过减少将相同数据读取到内存的次数来减少I/O.
     较高的压缩率通过使用更小的内存中空间提高查询性能。
     前缀区为作为缩小查询范围,索引区和内容区长度一样,
     解析出来一次性加载到数组中,查询性能提高3-5倍！

压缩：原版txt为38.5M,生成dat结构为3.68M 。
     和上一版本2.0不同的是索引区去掉了[开始IP数字]4字节,节省多1-2M。
     3.0版本只适用[全球版]，条件为ip段区间连续且覆盖所有IPV4。
     2.0版本适用[全球版][国内版][国外版]

性能：每秒解析1000多万ip (环境：CPU i7-7700K  + DDR2400 16G  + win10 X64)

创建：qqzeng-ip 于 2018-04-08
*/

type ipIndex struct {
	startip, endip             uint32
	local_offset, local_length uint32
}

type prefixIndex struct {
	start_index, end_index uint32
}

type IpSearch struct {
	data               []byte
	prefixMap          map[uint32]prefixIndex
	firstStartIpOffset uint32
	prefixStartOffset  uint32
	prefixEndOffset    uint32
	prefixCount        uint32
}

var ips *IpSearch = nil

func MustNew(file string) IpSearch {
	ips, _ = loadIpDat(file)
	return *ips
}

func New(file string) (IpSearch, error) {
	if ips == nil {
		var err error
		ips, err = loadIpDat(file)
		if err != nil {
			log.Fatal("the IP Dat loaded failed!")
			return *ips, err
		}
	}
	return *ips, nil
}

func loadIpDat(file string) (*IpSearch, error) {

	p := IpSearch{}
	//加载ip地址库信息
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	p.data = data
	p.prefixMap = make(map[uint32]prefixIndex)

	p.firstStartIpOffset = bytesToLong(data[0], data[1], data[2], data[3])
	p.prefixStartOffset = bytesToLong(data[8], data[9], data[10], data[11])
	p.prefixEndOffset = bytesToLong(data[12], data[13], data[14], data[15])
	p.prefixCount = (p.prefixEndOffset-p.prefixStartOffset)/9 + 1 // 前缀区块每组

	// 初始化前缀对应索引区区间
	indexBuffer := p.data[p.prefixStartOffset:(p.prefixEndOffset + 9)]
	for k := uint32(0); k < p.prefixCount; k++ {
		i := k * 9
		prefix := uint32(indexBuffer[i] & 0xFF)

		pf := prefixIndex{}
		pf.start_index = bytesToLong(indexBuffer[i+1], indexBuffer[i+2], indexBuffer[i+3], indexBuffer[i+4])
		pf.end_index = bytesToLong(indexBuffer[i+5], indexBuffer[i+6], indexBuffer[i+7], indexBuffer[i+8])
		p.prefixMap[prefix] = pf
	}
	return &p, nil
}

// 通过地址，获取下面的ip地址集合
func (p IpSearch) GetIpListByArea(area string) []string {
	return []string{}
}

func (p IpSearch) Get(ip string) string {
	ips := strings.Split(ip, ".")
	x, _ := strconv.Atoi(ips[0])
	prefix := uint32(x)
	intIP := ipToLong(ip)

	var high uint32 = 0
	var low uint32 = 0

	if _, ok := p.prefixMap[prefix]; ok {
		low = p.prefixMap[prefix].start_index
		high = p.prefixMap[prefix].end_index
	} else {
		return ""
	}

	var my_index uint32
	if low == high {
		my_index = low
	} else {
		my_index = p.binarySearch(low, high, intIP)
	}

	ipindex := ipIndex{}
	ipindex.getIndex(my_index, &p)

	if ipindex.startip <= intIP && ipindex.endip >= intIP {
		return ipindex.getLocal(&p)
	} else {
		return ""
	}
}

// 二分逼近算法
func (p IpSearch) binarySearch(low uint32, high uint32, k uint32) uint32 {
	var M uint32 = 0
	for low <= high {
		mid := (low + high) / 2

		endipNum := p.getEndIp(mid)
		if endipNum >= k {
			M = mid
			if mid == 0 {
				break // 防止溢出
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return M
}

// 只获取结束ip的数值
// 索引区第left个索引
// 返回结束ip的数值
func (p IpSearch) getEndIp(left uint32) uint32 {
	left_offset := p.firstStartIpOffset + left*12
	return bytesToLong(p.data[4+left_offset], p.data[5+left_offset], p.data[6+left_offset], p.data[7+left_offset])

}

func (p *ipIndex) getIndex(left uint32, ips *IpSearch) {
	left_offset := ips.firstStartIpOffset + left*12
	p.startip = bytesToLong(ips.data[left_offset], ips.data[1+left_offset], ips.data[2+left_offset], ips.data[3+left_offset])
	p.endip = bytesToLong(ips.data[4+left_offset], ips.data[5+left_offset], ips.data[6+left_offset], ips.data[7+left_offset])
	p.local_offset = bytesToLong3(ips.data[8+left_offset], ips.data[9+left_offset], ips.data[10+left_offset])
	p.local_length = uint32(ips.data[11+left_offset])
}

// / 返回地址信息
// / 地址信息的流位置
// / 地址信息的流长度
// 亚洲|中国|湖北|武汉|洪山|移动|420111|China|CN|114.34375|30.49989
func (p *ipIndex) getLocal(ips *IpSearch) string {
	bytes := ips.data[p.local_offset : p.local_offset+p.local_length]
	return string(bytes)

}

func ipToLong(ip string) uint32 {
	quads := strings.Split(ip, ".")
	var result uint32 = 0
	a, _ := strconv.Atoi(quads[3])
	result += uint32(a)
	b, _ := strconv.Atoi(quads[2])
	result += uint32(b) << 8
	c, _ := strconv.Atoi(quads[1])
	result += uint32(c) << 16
	d, _ := strconv.Atoi(quads[0])
	result += uint32(d) << 24
	return result
}

// 字节转整形
func bytesToLong(a, b, c, d byte) uint32 {
	a1 := uint32(a)
	b1 := uint32(b)
	c1 := uint32(c)
	d1 := uint32(d)
	return (a1 & 0xFF) | ((b1 << 8) & 0xFF00) | ((c1 << 16) & 0xFF0000) | ((d1 << 24) & 0xFF000000)
}

func bytesToLong3(a, b, c byte) uint32 {
	a1 := uint32(a)
	b1 := uint32(b)
	c1 := uint32(c)
	return (a1 & 0xFF) | ((b1 << 8) & 0xFF00) | ((c1 << 16) & 0xFF0000)

}
