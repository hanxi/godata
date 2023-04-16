package hello

import (
	"fmt"
)

type Observer interface {
	OnDirty(interface{})
}

type DataObject interface {
	NotifyDirty()
	Attach(o Observer)
}

type Base struct {
	DataObject
	observer Observer
	parent DataObject
	root DataObject
	self DataObject
}

func (x *Base) NotifyDirty() {
	if x.observer != nil {
		x.observer.OnDirty(x)
	}
	if x.root != nil && x.root != x.self {
		// 非根节点往上传递消息
		x.root.NotifyDirty()
	}
}

func (x *Base) Attach(o Observer) {
	x.observer = o
}

type PhoneNumber struct {
	Base
	_number  string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	_my *User
	_users  map[uint32]string `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	wrap_users *WrapPhoneNumberUsers
}

func NewPhoneNumber() *PhoneNumber {
	this := &PhoneNumber {}
	this.self = this
	this.root = this
	return this
}

func (x *PhoneNumber) GetNumber() string {
	if x == nil {
		return ""
	}

	return x._number
}

func (x *PhoneNumber) SetNumber(_number string) {
	if x == nil {
		return
	}
	x._number = _number
	x.NotifyDirty()
}

func (x *PhoneNumber) GetMy() *User {
	if x == nil {
		return nil
	}
	return x._my
}

func (x *PhoneNumber) SetMy(v *User) {
	if x == nil {
		return
	}

	x._my = v
	// v 的 root 为 x 的 root
	v.root = x.root
	x.NotifyDirty()
}

type User struct {
	Base
	_name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	_age  uint32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	_friends map[string]*User `protobuf:"bytes,3,rep,name=friends,proto3" json:"friends,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	wrap_friends *WrapUserFriends
	_sun *User
}

func NewUser() *User {
	this := &User {}
	this.self = this
	this.root = this
	return this
}

func (x *User) GetName() string {
	if x != nil {
		return x._name
	}
	return ""
}

func (x *User) SetName(name string) {
	if x != nil {
		x._name = name
		fmt.Printf("SetName name:%s\n", name)
		x.NotifyDirty()
	}
}

func (x *User) GetAge() uint32 {
	if x != nil {
		return x._age
	}
	return 0
}

func (x *User) SetAge(age uint32) {
	if x != nil {
		fmt.Println("SetAge:", age)
		x._age = age
		x.NotifyDirty()
	}
}

type WrapPhoneNumberUsers struct {
	Base
	parent *PhoneNumber
}

func (x *WrapPhoneNumberUsers) Set(key uint32, value string) {
	if x.parent == nil {
		return
	}
	if x.parent._users == nil {
		return
	}
	x.parent._users[key] = value
	fmt.Println("WrapPhoneNumberUsers Set")
	x.NotifyDirty()
}

func (x *WrapPhoneNumberUsers) Delete(key uint32) {
	if x.parent == nil {
		return
	}
	if x.parent._users == nil {
		return
	}
	delete(x.parent._users, key)
	x.NotifyDirty()
}

func (x *WrapPhoneNumberUsers) Get(key uint32) string {
	if x.parent == nil {
		return ""
	}
	if x.parent._users == nil {
		return ""
	}
	return x.parent._users[key]
}

func (x *PhoneNumber) GetUsers() *WrapPhoneNumberUsers {
	if x != nil {
		return x.wrap_users
	}
	return nil
}

func (x *PhoneNumber) SetUsers(v *WrapPhoneNumberUsers) {
	fmt.Println("SetUsers")
	if x != nil {
		x.wrap_users = v
		// v 的 root 为 x 的 root
		v.root = x.root
		x.NotifyDirty()
	}
}

func NewWrapPhoneNumberUsers(x *PhoneNumber) *WrapPhoneNumberUsers {
	this := &WrapPhoneNumberUsers {}
	x._users = make(map[uint32]string)
	this.parent = x
	this.self = this
	this.root = this
	return this
}

type WrapUserFriends struct {
	Base
	parent *User
}

func (x *WrapUserFriends) Set(key string, v *User) {
	if x.parent == nil {
		return
	}
	if x.parent._friends == nil {
		return
	}
	x.parent._friends[key] = v
	v.root = x.root
	fmt.Println("WrapUserFriends Set")
	x.NotifyDirty()
}

func (x *WrapUserFriends) Delete(key string) {
	if x.parent == nil {
		return
	}
	if x.parent._friends == nil {
		return
	}
	delete(x.parent._friends, key)
	x.NotifyDirty()
}

func (x *WrapUserFriends) Get(key string) *User {
	if x.parent == nil {
		return nil
	}
	if x.parent._friends == nil {
		return nil
	}
	return x.parent._friends[key]
}

func (x *User) GetFriends() *WrapUserFriends {
	if x != nil {
		return x.wrap_friends
	}
	return nil
}

func (x *User) SetFriends(v *WrapUserFriends) {
	fmt.Println("SetFriends")
	if x != nil {
		x.wrap_friends = v
		// v 的 root 为 x 的 root
		v.root = x.root
		x.NotifyDirty()
	}
}

func (x *User) GetSun() *User {
	if x != nil {
		return x._sun
	}
	return nil
}

func (x *User) SetSun(v *User) {
	if x != nil {
		x._sun = v
		// v 的 root 为 x 的 root
		v.root = x.root
		fmt.Printf("SetSun name:%s\n", v.GetName())
		x.NotifyDirty()
	}
}

func NewWrapUserFriends(x *User) *WrapUserFriends {
	this := &WrapUserFriends {}
	x._friends = make(map[string]*User)
	this.parent = x
	this.self = this
	this.root = this
	return this
}

////////////////////////

type Customer struct{}

var count uint32
func (c *Customer) OnDirty(i interface{}) {
	count++
	fmt.Println("OnDirty", i, count)
}

////////////////////////
