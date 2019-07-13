package main

import (
	"os/exec"
	"strings"
)

func mv(source string, destination string) (err error) {
	cmd := exec.Command("mv", source, destination)
	_, err = cmd.Output()
	return
}

func rm(filePath string) (err error) {
	if !pathIsExist(filePath) {
		return nil
	}

	cmd := exec.Command("rm", filePath)
	_, err = cmd.Output()
	return
}

func rmDir(filePath string) (err error) {
	if !pathIsExist(filePath) {
		return nil
	}

	cmd := exec.Command("rm", "-r", filePath)
	_, err = cmd.Output()
	return
}

func tarUnArchive(source string, destination string) (err error) {
	cmd := exec.Command("tar", "-zxf", source, "-C", destination)
	_, err = cmd.Output()
	return
}

// gzip先解压，再移动到目标文件夹
func gzUnArchive(source string, destination string) (err error) {
	cmd := exec.Command("gzip", "-d", source)
	_, err = cmd.Output()
	if err != nil {
		return
	}

	newFile := strings.TrimSuffix(source, FileSuffixGz)

	return mv(newFile, destination)
}

// gzip先压缩，再移动到目标文件夹
func gzArchive(source string, destination string) (err error) {
	cmd := exec.Command("gzip", source)
	_, err = cmd.Output()
	if err != nil {
		return
	}

	newFile := source + FileSuffixGz

	return mv(newFile, destination)
}
