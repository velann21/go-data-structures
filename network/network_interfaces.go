package main

import (
	"fmt"
	"log"
	"net"
	"syscall"
)

func main() {
	fmt.Println(GetInterfaces("en0"))

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

func CreateSocketAddress(iface net.Interface)( fd int,  err error ) {
	fd, err = syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 0 )

	if err != nil {
		return -1,err
	}

	if err = syscall.BindToDevice(fd, iface.Name); err != nil {
		return -1, err
	}
	return fd, nil
}

func BindToAddress(){

}