package main

import (
	"github.com/mholt/archiver"
)

const (
	FileSuffixGz     = ".gz"
	FileSuffixTarGz  = ".tar.gz"
	FileSuffixTar    = ".tar"
	FileSuffixTarBz  = ".tar.bz"
	FileSuffixTarBz2 = ".tar.bz2"
	FileSuffixTarXz  = ".tar.xz"
	FileSuffixTarSz  = ".tar.sz"
	FileSuffixTarLz4 = ".tar.lz4"
	FileSuffixZip    = ".zip"
	FileSuffixCsv    = ".csv"
)

type Archive struct {
	Suffix      string
	Source      string
	Destination string
}

// 解压文件
func (a *Archive) UnArchive() error {
	var err error = nil
	if a.Suffix == FileSuffixGz {
		err = gzUnArchive(a.Source, a.Destination)
	} else if a.Suffix == FileSuffixTarBz {
		err = tarUnArchive(a.Source, a.Destination)
	} else {
		err = archiver.Unarchive(a.Source, a.Destination)
	}

	return err
}

// 压缩文件
func (a *Archive) Archive() error {
	var err error = nil
	if a.Suffix == FileSuffixGz {
		err = gzArchive(a.Source, a.Destination)
	}  else {
		err = archiver.Archive([]string{a.Source}, a.Destination)
	}

	return err
}
