package main

import (
	"APIDemo/handler"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

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
