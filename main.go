package main

import "morpher-controller/route"

func main() {
	r := route.InitRoute()
	r.Run(":9000")
}
