package main

import(
	"fmt"
	"math/rand"
)

func recycle1() int {
	// for循环第一种写法
	for i:=1;i<=10;i++{
		fmt.Println("i的值",i)
	}
	return 1
}
func recycle2() {
	// for循环的第二种写法
	j := 1
	for j<=10 {
		fmt.Println("j的值",j)
		j++
	}
}
func recycle3() {
	// for循环的第3种写法
	k := 1
	for {
		if k<10 {
			fmt.Println("k的值",k)
			k++
		}else {
			break;
		}
	}
}

func main(){
	// recycle1()
	// recycle2()
	recycle3()
	// GO语言没有while和dowhile语句.这就是go的设计风格:尽量只有一种解决方案
	
	a := rand.Intn(100)
	fmt.Println("随机数",a)

}