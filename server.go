package main

import "kv_server/kv"

func main() {
	kv.RunServer("localhost:55555")
}
