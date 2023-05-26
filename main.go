package main

import (
	"main/providers"
)

func main() {
	providers.InitRouter()
	r := providers.InitRouter()
	//default size 32 MiB
	r.MaxMultipartMemory = 8 << 20 //8MiB setting a upper limit of 8 MiB for the uploaded file

	r.Run()

}
