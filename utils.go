package main

import (
	"os"
	"strings"
)

// 文件夹不存在则创建
func mkdirNotExist(path string, perm os.FileMode) error {
	if !pathIsExist(path) {
		return os.MkdirAll(path, perm)
	}

	return nil
}

// 判断文件或文件夹是否存在
func pathIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// 判断文件后缀是否正确
func isExt(fileName, ext string) bool {
	return len(fileName)-strings.LastIndex(fileName, ext) == len(ext)
}
