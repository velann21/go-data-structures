package recursion

var i int = 0

func RecusrionWithGlobal(n int)int{
	if n > 0{
		i = i+1
		return RecusrionWithGlobal(n-1)+i
	}
	return 0
}



