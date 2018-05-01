package main

import (
	"github.com/pluyckx/MusicManager/server"
)

func main() {
	err := server.ListenAndServe(12345)

	if err != nil {
		panic(err)
	}
}
