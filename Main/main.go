package main

import (
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"VideoAPI/TencentVideoAPI/Controller"
)

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(Controller.VideoInfoController)
}