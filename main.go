package main

import (
	"fmt"
	"play/ds"
)

func main() {
	// fmt.Println(playg.AddNumber(1, 2))
	// fmt.Println("gjergj")
	usrList := &ds.UserList{}
	usrList.Insert(&ds.User{Name: "gjergj"})
	usrList.Insert(&ds.User{Name: "klara"})
	usrList.Insert(&ds.User{Name: "sibel"})
	fmt.Println(usrList)

	usrList.Reverse()
	fmt.Println(usrList)
}
