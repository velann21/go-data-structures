package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ctc := make(chan []*Contact)
	wg.Add(1)
	getPagesInfo("temp.com", 1, wg, ctc)
	wg.Wait()
	defer close(ctc)
	for {
		select {
		case ch1Rcv, ok := <-ctc:
			if ok{
				fmt.Println(ch1Rcv)
			}else {
				ch1Rcv= nil
				close(ctc)
			}

		}
	}
	//for contact := range ctc{
	//	fmt.Println(contact)
	//}
}


func getPagesInfo(url string, depth int, wg *sync.WaitGroup, ctc chan []*Contact){
	//resultChan := make(chan []*Contact, 0)
	//contactsChan := make(chan []*Contact, 0)
	defer wg.Done()
	if depth == 0{
		return
	}
	page := NewPage(url)
	urls := page.getUrls() //["Google", "youtube", "FB"]

	for _, url := range urls{
		wg.Add(1)
		go func(){
			getPagesInfo(url, depth-1, wg, ctc) //["gmail, gmaps", "gsuite"]
			contacts := page.getContacts()
			ctc <- contacts
		}()
	}
	return
}

func consumerChannel(){

}


//func sendRequest(urls chan []string, contactsChan chan []*Contact, page PageI){
//	go func(){
//		for url := range urls {
//			page.getUrls()
//			contacts := page.getContacts()
//			contactsChan <- contacts
//		}
//	}()
//}
//
//func ConsumeContact(contactsChan chan []*Contact, resultChan chan []*Contact,page PageI){
//	go func(){
//		for contact := range contactsChan{
//			resultChan <- contact
//		}
//	}()
//}

type PageI interface {
	getUrls()[]string
	getContacts()[]*Contact
}

type Page struct {
	url string
}

func NewPage(url string)PageI{
	return &Page{url: url}
}

type Contact struct {
	name string
	age string
}

func NewContact()*Contact{
	return &Contact{}
}

func (page *Page) getUrls()[]string{
	return []string{"Google", "Youtube", "FB"}
}

func (page *Page) getContacts()[]*Contact{
	return []*Contact{NewContact(), NewContact(), NewContact()}
}