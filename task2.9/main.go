package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	ex, _ := os.Executable()
	filepath.Dir(ex)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ready for input.")

	for scanner.Scan() {
		text := scanner.Text()
		inputs := strings.Fields(text)

		switch inputs[0] {
		case "pwd":
			res, err := os.Getwd()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(res)
		case "echo":
			fmt.Println(inputs[1:])
		case "ps":
			matches, _ := filepath.Glob("/proc/*/exe")
			for _, file := range matches {
					target, _ := os.Readlink(file)
					if len(target) > 0 {
							fmt.Printf("%+v\n", target)
					}
			}
		case "kill":
			id, err := strconv.Atoi(inputs[1])
			if err != nil {
				fmt.Printf("Wrong process id %v\n", id)
				continue
			}

			p, err := os.FindProcess(id) 

			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			err3 := p.Kill()

			if err3 != nil {
				fmt.Println(err3.Error())
			}

		case "cd":
			err := os.Chdir(inputs[1])
			if err != nil {
				fmt.Println(err.Error())
			}
		case "exec":
			cmd := exec.Command(inputs[1], inputs[2:]...)
			res, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Println(string(res))
		case "quit":
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Printf("Unknown command: %v\n", inputs[0])
		}

	}
}