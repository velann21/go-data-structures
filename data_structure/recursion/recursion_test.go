package recursion

import (
	"fmt"
	"testing"
)

func TestHeadRecursion(t *testing.T){
	HeadRecursion(3)
}

func TestTailRecursion(t *testing.T){
	TailRecursion(3)
}

func TestRecusrionWithGlobal(t *testing.T){
	fmt.Println(RecusrionWithGlobal(3))
}



