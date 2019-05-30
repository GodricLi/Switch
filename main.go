package main

import (
	"fmt"
)

func testSwitch() {
	//switch语句外部定义a，
	a := 2
	switch a {
	case 1: //相当于if
		fmt.Println("a=1")
	case 2: //else if
		fmt.Println("a=2")
	case 3:
		fmt.Println("a=3")
	default: //else
		fmt.Println("invalid")
	}
}

func testSwitch2() {
	switch a := 3; a {
	case 2:
		fmt.Println("a=2")
	case 3:
		fmt.Println("a=3")
	}
	//switch语句内部定义的a在语句结束后清空
}

func getNumber(a int) int {
	return a
}

func testSwitch3() {
	//a可以接受一个函数
	switch a := getNumber(5); a {
	case 2:
		fmt.Println("a=2")
	case 3:
		fmt.Println("a=3")
	default:
		fmt.Println("out of switch")
	}
}

func testSwitch4() {
	switch a := 1; a {
	//一个case可以存在多个条件
	case 1, 2, 3, 4:
		fmt.Println("a in (1,2,3,4,5)")
	default:
		fmt.Println("out of switch")
	}
}

func testSwitch5() {
	num := getNumber(66)
	switch {
	case num >= 80 && num <= 100:
		fmt.Println("60~100")
	case num >= 60 && num <= 80:
		fmt.Println("60~80")
		fallthrough //穿透到下一个case，会执行下一个case的代码
	case num == 66:
		fmt.Println("666")
	}
}

func testMulti(){
	//9*9乘法表
	for i:=1;i < 10;i++{
		for j:=1;j < i;j++{
			fmt.Printf("%d*%d=%d \t",i,j,i*j)
		}
		fmt.Println()
	}
}

func main() {
	testSwitch()
	testSwitch2()
	testSwitch3()
	testSwitch4()
	testSwitch5()
	testMulti()
}
