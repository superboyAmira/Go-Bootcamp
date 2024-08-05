package main

import (
	"log"
	"team00/internal/cfg"
	"team00/pkg/transmitter_v1/team00/pkg/transmitter_v1"
)


const (
	xml_configuration string = "../../internal/cfg/cfg.xml"
)


func main() {
	config, err := cfg.LoadCfg(xml_configuration)
	if err != nil {
		log.Fatalf("Erorr reading config: %v", err.Error())
	}

	
}