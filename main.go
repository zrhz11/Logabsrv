package main

import (
	"Logabsrv/cmd"
	"Logabsrv/internal"
)

func main() {
	internal.InitDB()
	cmd.Execute()
	internal.CloseDB()
}
