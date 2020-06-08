package main

import "fmt"

func main() {
	rect := Rectangle{MinLat:90, MaxLat:-90, MinLng:180, MaxLng:-180}
	location := Location{}
	locdata := location.GetData()
	rectNew := formRectFromCoords(&rect, &locdata)
	fmt.Println(rectNew.MaxLng)
	fmt.Println(rectNew.MinLng)
	fmt.Println(rectNew.MaxLat)
	fmt.Println(rectNew.MinLat)
}

type Rectangle struct {
	MinLat float64
	MaxLat float64
	MinLng float64
	MaxLng float64
}

type Location struct {
	Lat float64
	Lng float64
}

func (loc *Location) GetData() []Location{

	// +Lat(North) -Lat(South) +Lng(Eastern) -Lng(Eastern)
	l1 := Location{Lat:80.3400, Lng: 4.9209}
	l2 := Location{Lat:-52.3400, Lng: -4.9209}
	l3 := Location{Lat:-51.3400, Lng: -50.9209}
	l4 := Location{Lat:62.3400, Lng: 80.9209}
	l5 := Location{Lat:-43.3400, Lng: -33.9209}
	l6 := Location{Lat:-50.3400, Lng: 179.9209}
	l7 := Location{Lat:52.3400, Lng: -168.9209}
	l8 := Location{Lat:89.3400, Lng: 150.9209}
	l9 := Location{Lat:70.3400, Lng: -142.9209}
	l10 := Location{Lat:43.3400, Lng: 179.9209}

	locs := make([]Location, 0)
	locs = append(locs, l1)
	locs = append(locs, l2)
	locs = append(locs, l3)
	locs = append(locs, l4)
	locs = append(locs, l5)
	locs = append(locs, l6)
	locs = append(locs, l7)
	locs = append(locs, l8)
	locs = append(locs, l9)
	locs = append(locs, l10)

	return locs
}

func formRectFromCoords(rect *Rectangle, locdata *[]Location)*Rectangle{

	if len(*locdata) < 0{
		return rect
	}

	for _,loData := range *locdata{
		if  loData.Lat < rect.MinLat{
			rect.MinLat = loData.Lat
		}

		if loData.Lng < rect.MinLng{
			rect.MinLng = loData.Lng
		}

		if loData.Lat > rect.MaxLat{
			rect.MaxLat = loData.Lat
		}

		if loData.Lng > rect.MaxLng{
			rect.MaxLng = loData.Lng
		}
	}
	return rect
}


=======
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
