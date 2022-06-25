package main

import (
	"fmt"
	"log"
	"pro/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	req := pb.HelloReq{
		Name: "hello proto",
	}

	rsp, err := proto.Marshal(&req)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	fmt.Printf("rsp = %s\n", rsp)
}
