package main

import (
	"fmt"
	"play/algo"
	"play/ds"
	"sort"
)

func main() {

	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 9

	elem := algo.BinarySearchNearest(a, len(a), x)
	if elem == -1 {
		fmt.Printf("%d not found in %v\n", x, a)
	} else {
		fmt.Printf("found %d at index %d in %v\n", x, elem, a)
	}

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
	return
	// fmt.Println(playg.AddNumber(1, 2))
	// fmt.Println("gjergj")
	usrList := &ds.UserList{}
	usrList.Insert(&ds.User{Name: "gjergj"})
	usrList.Insert(&ds.User{Name: "gjergj"})
	// usrList.Insert(&ds.User{Name: "klara"})
	// usrList.Insert(&ds.User{Name: "gjergj"})
	fmt.Println(usrList)

	//usrList.Reverse()
	fmt.Println(usrList)
	fmt.Println(usrList.IsPalindrome())
}
