package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//fmt.Println(GetInterfaces("en0"))
    fmt.Println(time.Now().UTC().Format(time.RFC3339Nano))
	fmt.Println(time.Now().Add(10*time.Minute).UTC().Format(time.RFC3339Nano))

}


func GetInterfaces(name string)(net.Interface, error){

	iface, err := net.Interfaces()
	if err != nil{
		log.Fatal("error occured while listing iface")
	}

	for _, i := range iface{

		fmt.Println(i)
	}


	return net.Interface{}, nil
}

//func CreateDataLinkSocket(iface net.Interface)( fd int,  err error ) {
//	fd, err = syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 0 )
//
//	if err != nil {
//		return -1,err
//	}
//
//	if err = syscall.BindToDevice(fd, iface.Name); err != nil {
//		return -1, err
//	}
//	return fd, nil
//}
//
//func BindSocketToAddress(){
//
//}
