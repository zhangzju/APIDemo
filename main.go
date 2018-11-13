package main

import (
	"APIDemo/handler"
	"flag"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ListenIP   string `yaml:"listen_ip"`
	ListenPort uint16 `yaml:"listen_port"`
}

var (
	configPath = flag.String("c", "modules/basic/config/config.yaml", "Runtime config file")
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

	addr := fmt.Sprintf("%s:%d", config.ListenIP, config.ListenPort)
	go handler.Run(addr)
}
