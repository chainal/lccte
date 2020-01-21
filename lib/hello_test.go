package lib

import "testing"

func TestSayHello(t *testing.T) {
	arr := [4]int{1, 2, 4, 5}
	var arr1 [4]int
	var arr2 = new([4]int)
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2)
}
