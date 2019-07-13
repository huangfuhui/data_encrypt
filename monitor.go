package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path"
	"path/filepath"
)

type FileWatcher struct {
	Watcher *fsnotify.Watcher
}

// 监控目录
func (w *FileWatcher) watch() {
	// 数据源路径
	dataSourcePath := conf.DataSourcePath

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-w.Watcher.Events:
				if !ok {
					return
				}

				// 文件创建事件
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Printf("file: %v has been created", event.Name)

					file, err := os.Stat(event.Name)
					if err != nil {
						log.Printf("monitor.go -> watch -> fsnotify.Create ; error: %v", err)
						return
					}

					// 如果新增的是文件夹，则添加监控
					if file.IsDir() {
						err = w.Watcher.Add(event.Name)
						if err != nil {
							log.Printf("monitor.go -> run -> fsnotify.Create -> watch.Add ; error: %v", err)
							return
						}
					} else {
						suffix := path.Ext(path.Base(event.Name))
						if suffix == FileSuffixCsv {
							go func() {
								csvReader := CsvReader{FilePath: event.Name}
								csvReader.Read()
							}()
						} else {
							go func() {
								tarReader := TarReader{FilePath: event.Name}
								tarReader.Read()
							}()
						}
					}
				}

				// 文件移除事件
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					_ = w.Watcher.Remove(event.Name)
				}

				// 文件重命名事件
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					_ = w.Watcher.Remove(event.Name)
				}

			case err, ok := <-w.Watcher.Errors:
				if !ok {
					return
				}
				log.Printf("monitor.go -> watch -> watch.Errors ; error: %v", err)
			}
		}
	}()

	// 监听数据源文件夹
	err := w.Watcher.Add(conf.DataSourcePath)
	if err != nil {
		log.Printf("monitor.go -> watch -> watch.add: %v ; error: %v", conf.DataSourcePath, err)
	}

	// 如果数据源路径下有多个文件夹，则遍历它们并加入监听
	err = filepath.Walk(dataSourcePath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = w.Watcher.Add(absPath)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("monitor.go -> run -> filepath.Walk ; error: %v", err)
		return
	}

	<-done
}
