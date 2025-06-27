package main

import (
	"fmt"
	"log"

	gogurt_proto "github.com/fozzbuzz/gogurt-rpc/proto"
	server "github.com/fozzbuzz/gogurt-rpc/server"
)

func main() {
	fmt.Println("gogurt")
	vendor := server.CreateGogurtVendor()
	request := &gogurt_proto.GetGogurtRequest{}
	gogurt, err := vendor.GetGogurt(request)
	if err != nil {
		log.Fatalf("Couldn't serve request: %q", err)
	}
	fmt.Printf("%+v\n", gogurt)
}
