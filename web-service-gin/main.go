package main

import "example.com/api"

func main() {
	api.Router().Run("localhost:7000")
}
