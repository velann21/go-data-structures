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
