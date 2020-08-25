package ds

import "fmt"

type User struct {
	Name string
	next *User
	prev *User
}

type UserList struct {
	root   *User
	length int
}

func (u *UserList) Insert(newuser *User) {
	if u.root == nil {
		u.root = newuser
		u.length++
		return
	}
	insElement := u.root
	for i := 0; i < u.length; i++ {
		if insElement.next == nil {
			insElement.next = newuser
			newuser.prev = insElement
			u.length++
			return
		}
		insElement = insElement.next
	}
}

func (u *UserList) String() string {
	var str string
	str = fmt.Sprintf("%v", u.length)
	for usr := u.root; usr != nil; usr = usr.next {
		str += fmt.Sprintf(" %s", usr.Name)
	}
	return str
}

func (u *UserList) Lenght() int {
	return u.length
}

func (u *UserList) Delete(usr *User) {
	usr.prev.next = usr.next
	usr.next.prev = usr.prev
	usr.next = nil
	usr.prev = nil
	u.length--
}

func (u *UserList) Reverse() {
	var current *User
	var prev *User
	var next *User
	current = u.root

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	u.root = prev
}
func (u *UserList) IsPalindrome() bool {
	if (u.length % 2) != 0 {
		return false
	}
	var left *User
	var right *User
	var middle *User
	//go to middle of list
	middle = u.root
	for i := 1; i <= u.length/2; i++ {
		middle = middle.next
	}
	left = middle.prev
	right = middle

	if left.Name != right.Name {
		return false
	}
	if (left.prev != nil) && right.next != nil {
		left = left.prev
		right = right.next
		if left.Name != right.Name {
			return false
		}
	}
	return true
}
