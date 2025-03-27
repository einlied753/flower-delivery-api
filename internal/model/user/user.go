package user

import (
	"api/internal/model"
	"fmt"
	"time"
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

func NewUser(fio string, email string, phone string, address string) model.Item {
	return User{
		id:       1,
		fio:      fio,
		email:    email,
		phone:    phone,
		address:  address,
		isActive: true,
	}
}

func (u *User) SetFio(fio string) {
	u.fio = fio
}

func (u *User) GetFio() string {
	return u.fio
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetPhone(phone string) {
	u.phone = phone
}

func (u *User) GetPhone() string {
	return u.phone
}

func (u *User) SetAddress(address string) {
	u.address = address
}

func (u *User) GetAddress() string {
	return u.address
}

func (u User) GetItem() User {
	return u
}

func (u User) SaveItem() {
	time.Sleep(time.Second)
	fmt.Println("The User \"" + u.fio + "\"  was saved")
}
