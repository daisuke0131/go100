package a

import "fmt"

func init(){
	fmt.Println("call init() 1")
}

func init(){
	fmt.Println("call init() 2")
}

func init(){
	fmt.Println("call init() 3")
}

func a_call1(){
	fmt.Println("call method")
}

func b_call1(){
	fmt.Println("call method")
}

type C struct{

}

func (c C) c_call(){
	fmt.Println("call c_call method")
}
