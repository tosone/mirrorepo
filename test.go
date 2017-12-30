package main

import (
	"fmt"

	"github.com/tosone/mirror-repo/bash"
)

func main() {
	//fmt.Println(IsRepo("/Users/tosone/awesome/bolt"))
	//fmt.Println("test")
	//fmt.Println(bash.ShortLoooooooog("/Users/tosone/awesome/bolt"))
	//c, err := exec.Command("sh", "-c", "git shortlog -n -s -e").CombinedOutput()
	//fmt.Println(string(c), err)

	fmt.Println(bash.RepoAge1("/Users/tosone/awesome/bolt"))
}
