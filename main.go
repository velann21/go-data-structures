package main

//
//
//package main

//
//import (
//	"errors"
//	"fmt"
//)
//
////input : [2,4,6,8](1)
////op := 2,4
////input2 := [6,2,4,8]
////op: [2,4]
////input3 := [6,2,-2,8]
////op:=[-2,8]
//
//type CircularQueue struct{
//	arr [10]int
//	front int
//	rear int
//	size int
//}
//
//func NewCircularQueue()*CircularQueue{
//	arr := [10]int{}
//	return &CircularQueue{arr: arr, front: 0, rear: 0}
//}
//
//func (cq *CircularQueue) Enqueue(data int){
//	fmt.Println((cq.rear+1)%(len(cq.arr)))
//	if (cq.rear+1)%(len(cq.arr)) == cq.front{
//		fmt.Println("Queue is full")
//		return
//	}else{
//		cq.rear = (cq.rear+1)%(len(cq.arr))
//		cq.arr[cq.rear] = data
//	}
//	return
//}
//
//func (cq *CircularQueue) IsQueueEmpty()bool{
//	return cq.front == cq.rear
//}
//
//func (cq *CircularQueue) Dequeue()int{
//	if cq.front == cq.rear{
//		fmt.Println("Queue is empty")
//	}else{
//		cq.front = (cq.front+1)%(len(cq.arr))
//		return cq.arr[cq.front]
//	}
//	return -1
//}
//
//func main() {
//	//sum, _ := FindPairSum([]int{3, 3, 4, 2}, 6)
//	//fmt.Println(sum)
//
//	num := 122
//	n := num %10
//	fmt.Println(n)
//
//	v := 11/10
//	fmt.Println(v)
//
//	err := printAllOperations(10, 20)
//	if err != nil{
//		fmt.Println(err)
//	}
//	err = printAllOperations(10, 0)
//	if err != nil{
//		fmt.Println(err)
//	}
//    fmt.Println("Continuing")
//
//	cq := NewCircularQueue()
//	cq.Enqueue(10)
//	cq.Enqueue(20)
//	cq.Enqueue(30)
//	cq.Enqueue(40)
//	cq.Enqueue(50)
//	cq.Enqueue(60)
//	cq.Enqueue(70)
//	cq.Enqueue(80)
//	cq.Enqueue(90)
//	cq.Enqueue(100)
//
//	fmt.Println(cq.Dequeue())
//	fmt.Println(cq.Dequeue())
//	fmt.Println(cq.Dequeue())
//	cq.Enqueue(100)
//
//	//o1, o2 := findPair([]int{5,6,1,2,3}, 5)
//	//fmt.Println(o1, ",",o2)
//	lastDigit := 12344346 % 10
//	number := 12344346 / 10
//	fmt.Println(lastDigit)
//	fmt.Println(number)
//	op := convertRec(120)
//	fmt.Println(op)
//}
//
//
//func printAllOperations(x int, y int)(err error){
//	defer func() {
//		r := recover()
//		if  r != nil {
//			fmt.Printf("Recovering from panic in printAllOperations error is: %v \n", r)
//			fmt.Println("Proceeding to alternative flow skipping division.")
//            err = errors.New("Something went wrong")
//		}
//	}()
//	sum, subtract, multiply, divide := x+y, x-y, x*y, x/y
//	fmt.Printf("sum=%v, subtract=%v, multiply=%v, divide=%v \n", sum, subtract, multiply, divide)
//	return err
//}
//
//
//
//func convertRec(num int)int{
//	if num==0 {
//		return 0
//	}
//	digit := num%10
//	if digit==0{
//		digit=5
//	}
//	head := convertRec(num/10)
//	tail := (head*10)+digit
//	return tail
//}
//
//
//func base62(deci uint64)string{
//	base62 := "0123456789ABCDEFGHIKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
//	var base  uint64 = 62
//	hashStr := ""
//	for deci > 0{
//		r := deci % base
//		deci /= base
//		hashStr = string(base62[r]) + hashStr
//	}
//	return hashStr
//}
//
//func findPair(input []int, target int)(int, int){
//	tracker := make(map[int]int)
//	pair1 := 0
//	pair2 := 0
//
//	//Write the value as the position of array in the map
//	for i, v := range input{
//		tracker[v] = i
//	}
//	//Iterate the array and find the diff with target and find whether the diff is present in array
//	//If present the return the current index and diffed elements index
//	for j, v := range input{
//		diff := target-v//1
//		if tracker[diff] != 0 && tracker[diff] != j{
//			pair1 = j
//			pair2 = tracker[diff]
//			break
//		}
//	}
//	return pair1, pair2
//}
//
//
//func FindPairSum(array []int, sum int)([]int, error){
//	InvalidData := errors.New("Invalid")
//	if array == nil{
//		return nil, InvalidData
//	}
//
//	if len(array) <=0 {
//
//		return nil, InvalidData
//
//	}
//
//	if len(array) == 1{
//		return nil, InvalidData
//	}
//
//	auxilaryMap := make(map[int]bool, 0)
//	resultArray := make([]int, 0)
//	diffSum := 0
//	for _, value := range array{
//		diffSum = sum-value
//		if auxilaryMap[diffSum] == false {
//			auxilaryMap[value] = true
//			continue
//		}
//		resultArray = append(resultArray, value)
//		resultArray = append(resultArray, diffSum)
//	}
//
//	return resultArray, nil
//}
//
//
//
//
//
//
////
////import (
////	//"context"
////	_ "encoding/json"
////	"errors"
////	"fmt"
////
////	//"fmt"
////	v3 "github.com/velann21/etcd/clientv3"
////	"github.com/velann21/etcd/pkg/transport"
////	"google.golang.org/grpc"
////
////	//"io/ioutil"
////	//"log"
////	//"os"
////	"time"
////)
////
////const (
////	ClientPEM  = "/mnt/etcd/ssl/certs/client.pem"
////	ClientKey  = "/mnt/etcd/ssl/certs/client-key.pem"
////	TrustedCA  = "/mnt/etcd/ssl/certs/ca.pem"
////	CLOUD_INFO = "/cloudinfo"
////	RI = "/cluster/resource_info/withfalco"
////)
////
////func main() {
////
////
////
////	err := returnError()
////	if err != nil{
////		if err.Error() == "Invalid request"{
////			goto backupstatement
////		}
////	}
////	fmt.Println("Normal exec")
////	backupstatement:
////		fmt.Println("Am excsd")
////
////	//ec, err := New([]string{"https://192.168.1.221:2379"})
////	//if err != nil{
////	//
////	//}
////	//jsonFile, err := os.Open("./resource_info.json")
////	//if err != nil {
////	//	fmt.Println(err)
////	//}
////	//
////	//// defer the closing of our jsonFile so that we can parse it later on
////	//defer jsonFile.Close()
////    //ctx := context.Background()
////	//byteValue, err := ioutil.ReadAll(jsonFile)
////	//if err != nil {
////	//	log.Println(err)
////	//	log.Println("Error occured")
////	//}
////	//fmt.Println(byteValue)
////	//
////	//_, err = ec.client.Put(ctx, RI, string(byteValue))
////	//if err != nil {
////	//	log.Println("Error occured")
////	//}
////}
////
////type EtcdClient struct {
////	client *v3.Client
////	url    []string
////}
////
////
////func New(url []string) (*EtcdClient, error) {
////	tlsInfo := transport.TLSInfo{
////		CertFile:      ClientPEM,
////		KeyFile:       ClientKey,
////		TrustedCAFile: TrustedCA,
////	}
////
////	tlsConfig, err := tlsInfo.ClientConfig()
////	if err != nil {
////		return nil, err
////	}
////
////	cfg := v3.Config{
////		Endpoints:   url,
////		DialTimeout: 5 * time.Second,
////		DialOptions: []grpc.DialOption{grpc.WithBlock()},
////		TLS:         tlsConfig,
////	}
////	c, err := v3.New(cfg)
////	if err != nil {
////		return nil, err
////	}
////
////	return &EtcdClient{client: c, url: url}, nil
////}
////
////
////
////func returnError()error{
////	return  errors.New("Invalid request12")
////}
