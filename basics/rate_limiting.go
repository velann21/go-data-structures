package main

import "fmt"

var rateLimiter = make(map[string]*User, 0)

func main() {
	done := make(chan bool)
	users := Getdata()
	channel := make(chan User, 100)

	go Producer(channel, *users)
	go Consumer(channel, done)

	fmt.Println("")

	<-done
}

func Producer(user chan User, data []User) {
	for _, v := range data {
		user <- v
	}
	close(user)
}

func Consumer(user chan User, done chan bool) {
	for {
		select {
		case msg, ok := <-user:
			if ok {
				fmt.Println("received job", msg)
				if rateLimiter[msg.ID] != nil{
					fmt.Println("Not nil")
				}else{
					rateLimiter[msg.ID] = &msg
				}
			}else {
				//done <- true
			}
		default:
			fmt.Println("No value ready, moving on.")
			done <- true

		}


	}
}

type User struct {
	ID       string
	Name     string
	IP       string
	Location UserLocation
}

type UserLocation struct {
	Lat float64
	Lng float64
}

func (user *User) Login() {

}

func Getdata() *[]User {
	user := User{ID: "123", Name: "velan", IP: "192.168.1.117", Location: UserLocation{Lat: 80.3400, Lng: 4.9209}}
	user1 := User{ID: "123", Name: "velan", IP: "192.168.1.117", Location: UserLocation{Lat: -80.3400, Lng: 4.9209}}
	user2 := User{ID: "123", Name: "velan", IP: "192.168.1.117", Location: UserLocation{Lat: -8980.3400, Lng: 4.9209}}
	user3 := User{ID: "123", Name: "velan", IP: "192.168.1.117", Location: UserLocation{Lat: 79.3400, Lng: 4.9209}}
	user4 := User{ID: "123", Name: "velan", IP: "192.168.1.117", Location: UserLocation{Lat: 35.3400, Lng: 4.9209}}
	user5 := User{ID: "123", Name: "velan", IP: "192.168.1.117", Location: UserLocation{Lat: 55.3400, Lng: 4.9209}}
	user6 := User{ID: "123", Name: "velan", IP: "192.168.1.117", Location: UserLocation{Lat: 77.3400, Lng: 4.9209}}
	user7 := User{ID: "125", Name: "velan21", IP: "192.168.1.118", Location: UserLocation{Lat: 81.3400, Lng: 4.9209}}
	user8 := User{ID: "127", Name: "indu21", IP: "192.168.1.119", Location: UserLocation{Lat: 22.3400, Lng: 4.9209}}
	user9 := User{ID: "128", Name: "mani", IP: "192.168.1.120", Location: UserLocation{Lat: 4.3400, Lng: 4.9209}}
	user10 := User{ID: "129", Name: "karthu", IP: "192.168.1.121", Location: UserLocation{Lat: 1.3400, Lng: 4.9209}}
	user11 := User{ID: "130", Name: "Naveen", IP: "192.168.1.122", Location: UserLocation{Lat: -81.3400, Lng: 4.9209}}
	user12 := User{ID: "131", Name: "nandu", IP: "192.168.1.123", Location: UserLocation{Lat: -80.3400, Lng: 4.9209}}
	users := []User{user, user1, user2, user3, user4, user5, user6, user7, user8, user9, user10, user11, user12}
	return &users
}