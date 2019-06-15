package main

import (
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"VideoAPI/Controller"
)

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(Controller.VideoListController)
}