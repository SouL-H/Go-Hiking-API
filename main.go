package main

import (
	"fmt"
	//"likyaapi/api"
	"likyaapi/db"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//database connect
	db.DbConnect()
	//api.ApiMain()
	//Program done 
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigs
		_ = sig
		done <- true
	}()
	<-done
	fmt.Println(("Program closed"))

}
