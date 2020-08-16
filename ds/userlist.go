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

func (u *UserList) Insert(user *User) {
	if u.root == nil {
		u.root = user
		u.length++
		return
	}
	usr := u.root
	for i := 0; i < u.length; i++ {
		if usr.next == nil {
			usr.next = user
			u.length++
			return
		}
		usr = usr.next
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
