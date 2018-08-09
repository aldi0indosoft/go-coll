package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var isRunHugo, isRunVue, isRunGit bool
var tmp bytes.Buffer
var theme string

func main() {
	hugoCommand := []string{"/S", "/C", "hugo"}
	vueCommand := []string{"/S", "/C", "vue"}
	gitCommand := []string{"/S", "/C", "git"}
	for i, arg := range os.Args {
		if i > 0 {
			switch arg {
			case "new":
				fmt.Println("Generating New Site")
				isRunHugo = true
				hugoCommand = append(hugoCommand, arg)
				hugoCommand = append(hugoCommand, "site")
				isRunVue = true
				vueCommand = append(vueCommand, "create")
			case "theme":
				fmt.Println("Adding template")
				isRunGit = true
				gitCommand = append(gitCommand, "clone")
			case "prod":
				fmt.Println("Building Site")
			// TODO: better testing
			case "test":
				fmt.Println("Testing Command")
				cmd := exec.Command("cmd", "/S", "/C", "mkdir", "testing")
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					panic(err)
				}
			default:
				if i == 1 {
					fmt.Println("Command not supported.")
				} else {
					hugoCommand = append(hugoCommand, arg)
					vueCommand = append(vueCommand, arg)
					if isRunGit && i == len(os.Args) - 1 {
						theme = arg
						tmp.WriteString("themes/")
						tmp.WriteString(arg)
						arg = tmp.String()
					}
					gitCommand = append(gitCommand, arg)
				}
			}
		}
	}

	if isRunHugo {
		hugoCmd := exec.Command("cmd", hugoCommand...)
		hugoCmd.Stdout = os.Stdout
		hugoCmd.Stdin = os.Stdin
		hugoCmd.Stderr = os.Stderr
		if err := hugoCmd.Run(); err != nil {
			panic(err)
		}
	}

	if isRunVue {
		vueCmd := exec.Command("cmd", vueCommand...)
		vueCmd.Stdout = os.Stdout
		vueCmd.Stdin = os.Stdin
		vueCmd.Stderr = os.Stderr
		if err := vueCmd.Run(); err != nil {
			panic(err)
		}
	}

	if isRunGit {
		fmt.Printf("args: %v", gitCommand)
		// gitCmd := exec.Command("cmd", gitCommand...)
		// gitCmd.Stdout = os.Stdout
		// gitCmd.Stdin = os.Stdin
		// gitCmd.Stderr = os.Stderr
		// if err := gitCmd.Run(); err != nil {
		// 	panic(err)
		// }

		// edit config

	}

}
