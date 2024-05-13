package main

import (
	"fmt"
	"os"

	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/version"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-version" || os.Args[1] == "version" || os.Args[1] == "-v") {
		fmt.Println(version.GetBuildInfo())
		os.Exit(0)
	}

	fmt.Println("Hello, World!")
}
