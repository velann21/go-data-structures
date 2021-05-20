//package main


package main

import "fmt"

//input : [2,4,6,8](1)
//op := 2,4
//input2 := [6,2,4,8]
//op: [2,4]
//input3 := [6,2,-2,8]
//op:=[-2,8]

func main() {
	o1, o2 := findPair([]int{5,6,1,2,3}, 5)
	fmt.Println(o1, ",",o2)
}

func findPair(input []int, target int)(int, int){
	tracker := make(map[int]int)
	pair1 := 0
	pair2 := 0

	//Write the value as the position of array in the map
	for i, v := range input{
		tracker[v] = i
	}
	//Iterate the array and find the diff with target and find whether the diff is present in array
	//If present the return the current index and diffed elements index
	for j, v := range input{
		diff := target-v//1
		if tracker[diff] != 0 && tracker[diff] != j{
			pair1 = j
			pair2 = tracker[diff]
			break
		}
	}
	return pair1, pair2
}








//
//import (
//	//"context"
//	_ "encoding/json"
//	"errors"
//	"fmt"
//
//	//"fmt"
//	v3 "github.com/velann21/etcd/clientv3"
//	"github.com/velann21/etcd/pkg/transport"
//	"google.golang.org/grpc"
//
//	//"io/ioutil"
//	//"log"
//	//"os"
//	"time"
//)
//
//const (
//	ClientPEM  = "/mnt/etcd/ssl/certs/client.pem"
//	ClientKey  = "/mnt/etcd/ssl/certs/client-key.pem"
//	TrustedCA  = "/mnt/etcd/ssl/certs/ca.pem"
//	CLOUD_INFO = "/cloudinfo"
//	RI = "/cluster/resource_info/withfalco"
//)
//
//func main() {
//
//
//
//	err := returnError()
//	if err != nil{
//		if err.Error() == "Invalid request"{
//			goto backupstatement
//		}
//	}
//	fmt.Println("Normal exec")
//	backupstatement:
//		fmt.Println("Am excsd")
//
//	//ec, err := New([]string{"https://192.168.1.221:2379"})
//	//if err != nil{
//	//
//	//}
//	//jsonFile, err := os.Open("./resource_info.json")
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//
//	//// defer the closing of our jsonFile so that we can parse it later on
//	//defer jsonFile.Close()
//    //ctx := context.Background()
//	//byteValue, err := ioutil.ReadAll(jsonFile)
//	//if err != nil {
//	//	log.Println(err)
//	//	log.Println("Error occured")
//	//}
//	//fmt.Println(byteValue)
//	//
//	//_, err = ec.client.Put(ctx, RI, string(byteValue))
//	//if err != nil {
//	//	log.Println("Error occured")
//	//}
//}
//
//type EtcdClient struct {
//	client *v3.Client
//	url    []string
//}
//
//
//func New(url []string) (*EtcdClient, error) {
//	tlsInfo := transport.TLSInfo{
//		CertFile:      ClientPEM,
//		KeyFile:       ClientKey,
//		TrustedCAFile: TrustedCA,
//	}
//
//	tlsConfig, err := tlsInfo.ClientConfig()
//	if err != nil {
//		return nil, err
//	}
//
//	cfg := v3.Config{
//		Endpoints:   url,
//		DialTimeout: 5 * time.Second,
//		DialOptions: []grpc.DialOption{grpc.WithBlock()},
//		TLS:         tlsConfig,
//	}
//	c, err := v3.New(cfg)
//	if err != nil {
//		return nil, err
//	}
//
//	return &EtcdClient{client: c, url: url}, nil
//}
//
//
//
//func returnError()error{
//	return  errors.New("Invalid request12")
//}
