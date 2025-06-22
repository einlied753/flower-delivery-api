package repository

import (
	"api/internal/model/user"
	"log"
	"strconv"
	"sync"
)

var (
	userSlice []*user.User
	userMutex sync.RWMutex
)

const userInputFile string = "./input_files/users.csv"
const userOutputFile string = "./output_files/users.csv"

func GetUsers() []*user.User {
	userMutex.RLock()
	defer userMutex.RUnlock()
	return userSlice
}

func ReadNewUsers(itemChan chan<- Item) {
	users, err := readFile(userInputFile)
	if err != nil {
		log.Fatalf("An error occurred while reading the new users from file: %v", err)
	}

	for i, u := range users {
		itemChan <- user.NewUser(i+1, u[1], u[2], u[3], u[4])
	}
}

func SaveUser(user *user.User) {
	userMutex.Lock()
	defer userMutex.Unlock()
	userSlice = append(userSlice, user)
	SaveUserToFile(user)
}

func SaveUserToFile(user *user.User) {

	strUser := []string{
		strconv.Itoa(user.GetId()),
		user.GetFio(),
		user.GetEmail(),
		user.GetPhone(),
		user.GetAddress(),
		strconv.FormatBool(user.GetIsActive()),
	}

	writeFile(userOutputFile, strUser)
}
