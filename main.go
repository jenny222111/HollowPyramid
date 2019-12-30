package main

import (
	"fmt"
)
//author : wuyijia
//func : pring pyramid 
//date: 2019-12-30 submit to github

//print pyramid
func pyramid(level int) {
	for i := 1; i <= level; i++ { //i = print layer 
		for k := 1; k <= level-i; k++ {// k print " "
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j==1||j==2*i-1||i==level{ //how many * you need to print
				fmt.Print("*")
			}else{
				fmt.Print(" ") //print " "
			}
			
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Println("input how many layers do you want to print")
	fmt.Scanln(&n)
	fmt.Printf("so you want to print %d layers, please wait\n", n)
	pyramid(n)
}
