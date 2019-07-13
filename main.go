package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

var conf *Conf

func main() {
	// 加载配置
	conf = &Conf{}
	if err := conf.LoadConf("conf.json", conf); err != nil {
		log.Printf("main.go -> main -> LoadConf ; error: %v", err)
		return
	}

	// 创建相关目录
	if err := mkdirNotExist(conf.DataSourcePath, os.ModePerm); err != nil {
		log.Printf("main.go -> main -> mkdir: %v ; error: %v", conf.DataSourcePath, err)
		return
	}
	if err := mkdirNotExist(conf.DataOutPutPath, os.ModePerm); err != nil {
		log.Printf("main.go -> main -> mkdir: %v; error: %v", conf.DataOutPutPath, err)
	}
	if err := mkdirNotExist(conf.DataBackupPath, os.ModePerm); err != nil {
		log.Printf("main.go -> main -> mkdir: %v; error: %v", conf.DataBackupPath, err)
	}

	// 启动文件监听
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("main.go -> main -> fsnotify.NewWatcher ; error: %v", err)
		return
	}

	defer func() {
		_ = watcher.Close()
	}()

	// 执行监听逻辑任务
	monitor := FileWatcher{Watcher: watcher}
	monitor.watch()
}
