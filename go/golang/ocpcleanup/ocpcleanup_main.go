package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	contents, err := ioutil.ReadFile("../config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	clusterNameRegEx := regexp.MustCompile(`"ocp-(.*)"`)
	contentsNew := []byte(clusterNameRegEx.ReplaceAllString(string(contents), fmt.Sprintf("\"ocp-%s\"", os.Args[1])))

	err = ioutil.WriteFile("config.json", contentsNew, fs.ModeDevice)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
