package main

import (
	"main/providers"
)

func main() {
	providers.InitRouter()
	r := providers.InitRouter()
	r.Run()

}
