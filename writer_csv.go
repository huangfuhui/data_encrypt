package main

import (
	"encoding/csv"
	"os"
	"path"
)

type CsvWriter struct {
	FilePath string
	Writer   *csv.Writer
	File     *os.File
}

func (c *CsvWriter) Open() error {
	err := mkdirNotExist(path.Dir(c.FilePath), os.ModePerm)
	if err != nil {
		return err
	}

	c.File, err = os.OpenFile(c.FilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	c.Writer = csv.NewWriter(c.File)

	return nil
}

func (c *CsvWriter) Write(record []string) error {
	if len(record) == 0 {
		return nil
	}

	return c.Writer.Write(record)
}
