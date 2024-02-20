package domain

import (
	"errors"
	"time"
)

type User struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Count      int    `json:"count"`
	Updated_at int    `json:"updated_at"`
	Nickname   string `json:"nickname"`
}
type CreateUserOpts struct {
	Id       string
	Username string
	Nickname string
}

/*
	NewUser instances an NONEXISTENT user

this constructor should not be used inside repository. instead use the User struct directly
*/
func NewUser(opts CreateUserOpts) (*User, error) {
	user := &User{
		Id:       opts.Id,
		Username: opts.Username,
		Nickname: opts.Nickname,
		Count:    0,
	}
	err := user.validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Pay() {
	u.touch()
	u.Count++
}

func (u *User) GetNickname() string {
	nickname := u.Username
	if u.Nickname != "" {
		nickname = u.Nickname
	}
	return nickname
}

func (u *User) AlreadySubmitted() bool {
	today := time.Now().Day()
	return u.Updated_at == today
}
func (u *User) IsNotSubscribed() bool {
	return u.Id == ""
}

func (u *User) validate() error {
	var errorList []error
	if u.Username == "" {
		errorList = append(errorList, errors.New("username is required"))
	}
	if u.Id == "" {
		errorList = append(errorList, errors.New("id is required"))
	}
	if u.Count < 0 {
		errorList = append(errorList, errors.New("count must be greater than or equal to 0"))
	}
	if len(errorList) > 0 {
		concatenateMessage := ""
		for _, err := range errorList {
			concatenateMessage += err.Error() + ", "
		}
		return errors.New(concatenateMessage)
	}
	return nil
}

func (u *User) touch() {
	u.Updated_at = time.Now().Day()
}
