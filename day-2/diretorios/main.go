package main

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
)

func main() {
	gottaGoMyOwnWay()
}

func gottaGoMyOwnWay() {
	dir := path.Join("whatAboutUs", "WhatAboutEverythingWeveBeenThrought")
	fmt.Println(dir)
}

func mexendoEmDir() {
	filepath.WalkDir(".", learningToWalkAgain)
}

func learningToWalkAgain(path string, d fs.DirEntry, err error) error {
	fmt.Println("is it my goal??", d.IsDir())
	info, err := d.Info()
	if err != nil {
		return err
	}
	fmt.Println("who we are ", info.Name())

	return nil
}
