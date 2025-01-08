package main

import "fmt"

var (
	Version   string
	Commit    string
	BuildTime string
)

func main() {
	fmt.Printf("Version: %s\nCommit: %s\nBuildTime: %s\n", Version, Commit, BuildTime)
	fmt.Println("hello nocgo!")
}
