package main

import (
	"fmt"
)

func main() {
	var dclient Defclient = Defclient{20, 15}
	var lclient Limclient = Limclient{dclient, 150, 1.33}
	var prClient1 Prefclient1 = Prefclient1{dclient, 0.67}
	var prclient2 Prefclient2 = Prefclient2{dclient, 50, 0.0}
	fmt.Println("dClient")
	fmt.Println(dclient.Calculate())
	fmt.Println("lClient")
	fmt.Println(lclient.Calculate())
	fmt.Println("pr1Client")
	fmt.Println(prClient1.Calculate())
	fmt.Println("pr2Client")
	fmt.Println(prclient2.Calculate())

}
