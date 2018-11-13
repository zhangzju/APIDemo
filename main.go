package main

import (
	"APIDemo/handler"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ListenIP   string `yaml:"listen_ip"`
	ListenPort uint16 `yaml:"listen_port"`
}

var (
	configPath = flag.String("c", "modules/basic/config/runtime.yaml", "Runtime config file")
)

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		fmt.Printf("Load monitor config file %s failed: %v", *configPath, err)
		return
	}
	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Printf("Unmarshal yaml failed: %v", err)
		return
	}

	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watch.Close()
	err = watch.Add("./hot-reload.yaml")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						log.Println("创建文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("写入文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("删除文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.Println("重命名文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						log.Println("修改权限 : ", ev.Name)
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()

	addr := fmt.Sprintf("%s:%d", config.ListenIP, config.ListenPort)
	go handler.Run(addr)

	<-c
}
