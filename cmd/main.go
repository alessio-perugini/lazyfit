package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"lazyfit"
	"log"
)

var filePath = "."
var telegramToken = ""

func main() {
	flag.StringVar(&filePath, "config", ".", "file path of yaml config file.")
	flag.StringVar(&telegramToken, "token", "", "token telegram")

	flag.Usage = func() { //help flag
		fmt.Fprintf(flag.CommandLine.Output(), "\n\nUsage: lazyfit [options]\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	lazyfit.Conf = getConf()
	lazyfit.Start()
}

func getConf() *lazyfit.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filePath)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &lazyfit.Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		log.Fatalf("unable to decode into config struct, %v", err)
	}

	return conf
}
