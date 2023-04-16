package hello

import (
	"testing"
)

func TestExample(t *testing.T) {
	pn := NewPhoneNumber()
    observer := &Customer{}
    pn.Attach(observer)
	pn.SetNumber("123") // dirty 1

	user := NewUser()
	user.SetAge(18) // no dirty: because not set real root
	pn.SetMy(user) // dirty 2

	user.SetName("hanxi") // dirty 3

	users := NewWrapPhoneNumberUsers(pn)
	users.Set(1, "hanxi1") // no dirty: because not set real root
	pn.SetUsers(users) // dirty4

	users.Set(2, "hanxi2") // dirty5

	friend1 := NewUser()
	friend1.SetName("f1")
	friends := NewWrapUserFriends(user)
	friends.Set("f1", friend1)
	user.SetFriends(friends) // dirty6
	friend1.SetAge(11) // dirty7

	friend2 := NewUser()
	friend2.SetName("f2")
	friends.Set("f2", friend2) // dirty8

	sun := NewUser()
	sun.SetName("sun")
	user.SetSun(sun) // dirty9
	sun.SetAge(9) // dirty10

	sun.SetSun(friend1) // dirty11
	friend1.SetAge(10) // dirty12
}
