package main

import "github.com/pathcl/fakeme/cmd"

func main() {
	root := cmd.Root()
	root.Execute()
}
