package user

import (
	"fmt"
)

// User who can be a customer of the store or a temporary user who makes an order
type User struct {
	id       int
	fio      string
	email    string
	phone    string
	address  string
	isActive bool
}

func NewUser(id int, fio string, email string, phone string, address string) *User {
	return &User{
		id:       id,
		fio:      fio,
		email:    email,
		phone:    phone,
		address:  address,
		isActive: true,
	}
}

func (u User) GetId() int {
	return u.id
}

func (u *User) SetFio(fio string) {
	u.fio = fio
}

func (u User) GetFio() string {
	return u.fio
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u User) GetEmail() string {
	return u.email
}

func (u *User) SetPhone(phone string) {
	u.phone = phone
}

func (u User) GetPhone() string {
	return u.phone
}

func (u *User) SetAddress(address string) {
	u.address = address
}

func (u User) GetAddress() string {
	return u.address
}

func (u User) GetIsActive() bool {
	return u.isActive
}

func (u User) GetItem() User {
	return u
}

func (u *User) SaveItemLog() {
	fmt.Println("The User \"" + u.GetFio() + "\"  was saved")
}
