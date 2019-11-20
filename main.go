package main

import (
	"github.com/cisordeng/beego/xenon"

	_ "dolphin/model"
	_ "dolphin/rest"
)

func main() {
	xenon.Run()
}
