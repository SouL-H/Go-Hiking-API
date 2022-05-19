package main

import (
	"fmt"
	//"likyaapi/api"
	"likyaapi/db"
	//jsonTodb "likyaapi/json_db"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//Json data Insert DB
	//pathArr := []string{"./jsonData/1_fethiye_kayakoy_ovacık.json", "./jsonData/2_ovacık_kozagac_faralya.json", "./jsonData/3_faralya_kabak.json"}

	// for _, k := range pathArr{
	// 	jsonTodb.JsonToDB(k)
	// }
	//Route Id =>> return RotaName,Lat,Lon, 
	a:=db.GetRoute(7)
	for _, k := range a{
		fmt.Print(k.RouteName)
	}
	b:=db.GetRoute(9)
	for _, k := range b{
		fmt.Print(k.RouteName)
	}
	//database connect
	//db.DbConnect()
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
