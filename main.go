package main

import (
	"fmt"
	"github.com/spf13/viper"
	"flag"
)

func main() {
	fileName := flag.String("name", "config", "fileName to load")
	filePath := flag.String("path", ".", "filePath to load")
	flag.Parse()

	viper.SetConfigName(*fileName)
	viper.AddConfigPath(*filePath)
	if err := viper.ReadInConfig(); err != nil {
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }

  var parsedConfig interface{}
  if err := viper.Unmarshal(&parsedConfig); err != nil {
    panic(fmt.Errorf("Fatal unmarshaling config: %s \n", err))
  }

  fmt.Printf("%#v")
}