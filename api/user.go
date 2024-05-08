package api

import "errors"

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

var users = []User{
	{"anil", "nc"},
	{"raj", "tx"},
	{"kim", "sc"},
	{"kal", "ca"},
}

func (s *Store) GetUser(name string) (User, error) {
	for _, v := range users {
		if v.Name == name {
			return v, nil
		}
	}
	return User{}, errors.New("not found")
}

func (s *Store) CreateUser(user User) error {
	for _, v := range users {
		if v.Name == user.Name {
			return errors.New("already exists")
		}
	}
	users = append(users, user)
	return nil
}

func (s *Store) UpdateUser(user User) error {
	for i, v := range users {
		if v.Name == user.Name {
			users[i].Address = user.Address
			return nil
		}
	}
	return errors.New("not found")
}

func (s *Store) DeleteUser(name string) error {
	indexToDelete := -1
	for i, v := range users {
		if v.Name == name {
			indexToDelete = i
		}
	}
	if indexToDelete != -1 {
		users = append(users[:indexToDelete], users[indexToDelete+1:]...)
		return nil
	}
	return errors.New("not found")
}
