package main

import (
	"log"
	"time"

	"github.com/warnawski/space-evolution/pkg/sfp"
)

func main() {

	//logger.ConfigureLogger()
	//
	//conf := configurate.NewConf("../../examples/configuration/config.yaml")
	//if conf == nil {
	//	return
	//}
	//
	//err := conf.LoadConfig()
	//if err != nil {
	//	return
	//}
	//log.Printf("Server configuration loaded: server_name: %s", conf.Data.ServerName)

	//Test

	rb := sfp.NewRingBuffer(1000)

	time.Sleep(1 * time.Second)

	go func() {
		for {
			example_package := []byte{125, 124, 215}
			rb.Push(example_package)
		}
	}()

	go func() {
		for {
			example_data := rb.Pop()
			if example_data != nil {
				log.Printf("good package: %s", example_data)
			} else {
				time.Sleep(1 * time.Millisecond)
				log.Printf("buffer is empty")
			}
		}
	}()

	select {}

}
