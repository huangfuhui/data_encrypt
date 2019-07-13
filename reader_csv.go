package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path"
)

type CsvReader struct {
	FilePath string
}

func (c *CsvReader) Read() {
	// 打开源文件
	inputFile, err := os.Open(c.FilePath)
	if err != nil {
		log.Printf("reader_csv.go -> Read -> os.Open %v ; error: %v", c.FilePath, err)
		return
	}
	defer func() {
		_ = inputFile.Close()
	}()

	// 删除空文件和异常文件
	if info, err := inputFile.Stat(); err != nil || info.Size() == 0 {
		_ = rm(c.FilePath)
		return
	}

	// 读取文件
	reader := csv.NewReader(inputFile)
	record, err := reader.Read()
	if err == io.EOF {
		_ = rm(c.FilePath)
		return
	} else if err != nil {
		log.Printf("reader_csv.go -> Read -> reader.Read ; error: %v", err)
		return
	}

	// 创建写入文件
	outputFile := conf.DataOutPutPath + "/" + path.Base(c.FilePath)
	writer := CsvWriter{FilePath: outputFile}
	err = writer.Open()
	if err != nil {
		log.Printf("reader_csv.go -> Read -> w.Open ; error: %v", err)
		return
	}
	defer func() {
		if writer.File != nil {
			_ = writer.File.Close()
		}
	}()

	temp := template{}
	// 读取第一行文件，获取目标字段的索引
	indexArr := temp.GetTargetIndex(record)

	// 写入首行
	err = writer.Write(record)
	if err != nil {
		log.Printf("reader_scv.go -> Read -> writer.Write ; error: %v", err)
		return
	}

	// 循环读取剩余行，加密后写入新文件
	for {
		record, err = reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Printf("reader_csv.go -> Read -> reader.Read ; error: %v", err)
				break
			}
		}

		// 加密或解密
		isEncrypt := conf.IsEncrypt
		if isEncrypt {
			for _, v := range indexArr {
				result, err := aesEncrypt([]byte(record[v]), conf.GetEncryptKey())
				if err != nil {
					log.Printf("reader_csv.go -> Read -> aesEncrypt ; error: %v", err)
					break
				}

				record[v] = result
			}
		} else {
			for _, v := range indexArr {
				result, err := aesDecrypt(record[v], conf.GetEncryptKey())
				if err != nil {
					log.Printf("reader_csv.go -> Read -> aesEncrypt ; error: %v", err)
					break
				}

				record[v] = string(result)
			}
		}

		// 写入加密数据
		err = writer.Write(record)
		if err != nil {
			log.Printf("reader_scv.go -> Read -> writer.Write ; error: %v", err)
			break
		}
	}

	writer.Writer.Flush()
	err = writer.Writer.Error()
	if err != nil {
		log.Printf("reader_scv.go -> Read -> writer.Writer.Flush ; error: %v", err)
	}

	// 压缩加密后的文件
	archive := Archive{
		Suffix:      FileSuffixGz,
		Source:      outputFile,
		Destination: outputFile + FileSuffixGz,
	}
	err = archive.Archive()
	if err != nil {
		log.Printf("reader_tar.go -> Read -> archive.UnArchive ; error: %v", err)
		return
	}

	// 删除加密文件
	_ = rm(outputFile)

	// 删除源文件
	_ = rm(c.FilePath)
}
