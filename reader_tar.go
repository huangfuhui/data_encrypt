package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type TarReader struct {
	FilePath string
}

func (t *TarReader) Read() {
	fullFileName := path.Base(t.FilePath)

	var suffixWithPoint string
	if isExt(fullFileName, FileSuffixGz) {
		suffixWithPoint = FileSuffixGz
	} else if isExt(fullFileName, FileSuffixTarGz) {
		suffixWithPoint = FileSuffixTarGz
	} else if isExt(fullFileName, FileSuffixTar) {
		suffixWithPoint = FileSuffixTar
	} else if isExt(fullFileName, FileSuffixTarBz) {
		suffixWithPoint = FileSuffixTarBz
	} else if isExt(fullFileName, FileSuffixTarBz2) {
		suffixWithPoint = FileSuffixTarBz2
	} else if isExt(fullFileName, FileSuffixTarXz) {
		suffixWithPoint = FileSuffixTarXz
	} else if isExt(fullFileName, FileSuffixTarSz) {
		suffixWithPoint = FileSuffixTarSz
	} else if isExt(fullFileName, FileSuffixTarLz4) {
		suffixWithPoint = FileSuffixTarLz4
	} else if isExt(fullFileName, FileSuffixZip) {
		suffixWithPoint = FileSuffixZip
	} else {
		suffixWithPoint = path.Ext(fullFileName)
	}

	// 将文件从源文件夹移至备份目录文件夹下
	backupPath := conf.DataBackupPath
	if err := mv(t.FilePath, backupPath); err != nil {
		log.Printf("reader_tar.go -> Read -> mv %v %v ; error: %v", t.FilePath, backupPath, err)
		return
	}

	// 在备份路径下创建解压文件夹
	unArchiveDirName := strings.TrimSuffix(fullFileName, suffixWithPoint) + strings.ReplaceAll(suffixWithPoint, ".", "_")
	destinationPath := backupPath + "/" + unArchiveDirName
	source := backupPath + "/" + fullFileName
	if err := mkdirNotExist(destinationPath, os.ModePerm); err != nil {
		log.Printf("reader_tar.go -> Read -> mkdir %v %v ; error: %v", destinationPath, os.ModePerm, err)
		return
	}

	defer func() {
		// 删除创建的解压文件夹
		if err := rmDir(destinationPath); err != nil {
			log.Printf("reader_tar.go -> Read -> rmDir %v; error: %v", destinationPath, err)
		}

		// 删除备份文件夹内的源文件
		if err := rm(source); err != nil {
			log.Printf("reader_tar.go -> Read -> rm %v; error: %v", source, err)
		}
	}()

	// 解压文件
	archive := Archive{
		Suffix:      suffixWithPoint,
		Source:      source,
		Destination: destinationPath,
	}
	err := archive.UnArchive()
	if err != nil {
		log.Printf("reader_tar.go -> Read -> archive.UnArchive ; error: %v", err)
		return
	}

	// 源文件所在目录
	parentDirPath := path.Dir(t.FilePath)
	// 遍历解压后的文件，防止多层文件夹嵌套的情况
	err = filepath.Walk(destinationPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("reader_tar.go -> Read -> filepath.Walk %v ; error: %v", destinationPath, err)
			return err
		}

		// 将解压后的文件移回原目录
		if !info.IsDir() {
			err = mv(path, parentDirPath)
			if err != nil {
				log.Printf("reader_tar.go -> Read -> mv %v %v; error: %v", path, parentDirPath, err)
				return err
			}
		}
		return nil
	})
}
