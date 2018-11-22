package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的姓名")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Printf("An Error Occurred :%+v \n", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("你好， %s! 请问我能为你做什么？", name)
	}
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An Error Occurred :%+v \n", err)
			continue
		}
		input = input[:len(input)-1]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "hello":
			fmt.Println("Hello")
		case "bye", "nothing":
			fmt.Println("Bye")
			os.Exit(2)
		default:
			fmt.Println("I can't help you")
		}
	}
}
