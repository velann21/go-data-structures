package main

import (
	//"context"
	_ "encoding/json"
	"errors"
	"fmt"

	//"fmt"
	v3 "github.com/velann21/etcd/clientv3"
	"github.com/velann21/etcd/pkg/transport"
	"google.golang.org/grpc"
	
	//"io/ioutil"
	//"log"
	//"os"
	"time"
)

const (
	ClientPEM  = "/mnt/etcd/ssl/certs/client.pem"
	ClientKey  = "/mnt/etcd/ssl/certs/client-key.pem"
	TrustedCA  = "/mnt/etcd/ssl/certs/ca.pem"
	CLOUD_INFO = "/cloudinfo"
	RI = "/cluster/resource_info/withfalco"
)

func main() {


	err := returnError()
	if err != nil{
		if err.Error() == "Invalid request"{
			goto backupstatement
		}
	}
	fmt.Println("Normal exec")
	backupstatement:
		fmt.Println("Am excsd")

	//ec, err := New([]string{"https://192.168.1.221:2379"})
	//if err != nil{
	//
	//}
	//jsonFile, err := os.Open("./resource_info.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// defer the closing of our jsonFile so that we can parse it later on
	//defer jsonFile.Close()
    //ctx := context.Background()
	//byteValue, err := ioutil.ReadAll(jsonFile)
	//if err != nil {
	//	log.Println(err)
	//	log.Println("Error occured")
	//}
	//fmt.Println(byteValue)
	//
	//_, err = ec.client.Put(ctx, RI, string(byteValue))
	//if err != nil {
	//	log.Println("Error occured")
	//}
}

type EtcdClient struct {
	client *v3.Client
	url    []string
}


func New(url []string) (*EtcdClient, error) {
	tlsInfo := transport.TLSInfo{
		CertFile:      ClientPEM,
		KeyFile:       ClientKey,
		TrustedCAFile: TrustedCA,
	}

	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		return nil, err
	}

	cfg := v3.Config{
		Endpoints:   url,
		DialTimeout: 5 * time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
		TLS:         tlsConfig,
	}
	c, err := v3.New(cfg)
	if err != nil {
		return nil, err
	}

	return &EtcdClient{client: c, url: url}, nil
}


func returnError()error{
	return  errors.New("Invalid request12")
}
