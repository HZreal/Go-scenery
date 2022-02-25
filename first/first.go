// 声明当前所在包
package first

//一个文件夹下的所有文件必须使用同一个包名

import (
	"fmt"
)

func main() {
	var username = "huang"
	fmt.Println(username)

	var password string
	password = "zhen"
	fmt.Println(password)

	var res string
	res = username + password
	fmt.Println(res)

	var ccc bool
	fmt.Println(ccc)

	aaa := "hahahaha" // var a string = "hahahaha"
	fmt.Println(aaa)

	const WIDTH = 20
	const LENGTH int = 10

	//逻辑位运算
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c)

	// 条件判断
	var _num_1 = 10
	var _num_2 = 5

	if _num_1 > _num_2 {
		fmt.Println(_num_1)
	} else {
		fmt.Println(_num_2)
	}

	var marks = 90
	var grade string = "D"
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println(grade)

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	//	循环
	var sum = 0 // init 和 post 参数是可选的，我们可以直接省略
	for he := 1; he <= 10; he++ {
		sum += he
	}
	fmt.Println("\nsum-----", sum)

	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}

	// goto语句
	var _aa int = 10
LOOP:
	for _aa < 20 {
		if _aa == 15 {
			/* 跳过迭代 */
			_aa = _aa + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", _aa)
		_aa++
	}

	resNum := num_max(1, 6)
	fmt.Println("resNum---------------", resNum)

	str1, str2 := swap("happy", "nice a day")
	fmt.Println("swap result-------------", str1, str2)

}

func num_max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
