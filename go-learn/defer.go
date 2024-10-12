// https://blog.51cto.com/tinywan/6161993
// defer 最后执行
// defer 调用的函数参数的值在defer定义时就确定了
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("asd")
	fmt.Println("aaa")
}
