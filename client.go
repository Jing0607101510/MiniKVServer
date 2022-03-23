package main

import (
	"bufio"
	"fmt"
	"kv_server/kv"
	"os"
	"regexp"
	"strings"
)

func main() {
	c := kv.NewClient("localhost:55555")
	scanner := bufio.NewScanner(os.Stdin)
	reg := regexp.MustCompile("\\s+")
	fmt.Print("> ")
	for scanner.Scan() {
		input := strings.Trim(scanner.Text(), " ")
		input = reg.ReplaceAllString(input, " ")
		args := strings.Split(input, " ")
		switch len(args) {
		case 2:
			if args[0] == "get" {
				key := args[1]
				status, val := c.Get(key)
				if status == kv.Success {
					fmt.Println(val)
				} else {
					fmt.Printf("%q\n", "The Key is not in the kv db!")
				}
			} else {
				fmt.Printf("%q\n", "The cmd does not correct.")
			}
		case 3:
			if args[0] == "set" {
				key := args[1]
				val := args[2]
				status := c.Set(key, val)
				if status == kv.Success {
					fmt.Printf("%q\n", "The Key: Value is Saved Successfully.")
				} else {
					fmt.Printf("%q\n", "The Key: Value is not Savad Successfully.")
				}
			} else {
				fmt.Printf("%q\n", "The cmd does not correct.")
			}
		default:
			fmt.Printf("%q\n", "The cmd does not correct.")
		}
		fmt.Print("> ")
	}
}
