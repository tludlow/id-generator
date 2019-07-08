package main

import "id-generator/comms"

func main() {
	//Start our HTTP router.
	comms.CreateRouter()
}
