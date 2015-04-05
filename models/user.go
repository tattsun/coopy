package models

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"

	"github.com/tattsun/coopy/helpers"
)

func hashPassword(password string, salt string) string {
	salted := []byte(password + salt)
	sum := sha512.Sum512(salted)
	return base64.StdEncoding.EncodeToString(sum[:])
}

func NewAuthUserInfo(userid string, password string) *UserAuthInfo {
	salt := helpers.RandomStr(20)
	hashed := hashPassword(password, salt)
	return &UserAuthInfo{UserID: userid, HashedPassword: hashed, Salt: salt}
}

func CreateUser(user *User, password string) (*UserAuthInfo, error) {
	if IsExists(User{UserID: user.UserID}) {
		return nil, fmt.Errorf("UserID %s duplicates.", user.UserID)
	}
	db.Create(user)

	auth := NewAuthUserInfo(user.UserID, password)
	db.Create(auth)

	return auth, nil
}

func IsExists(user User) bool {
	cnt := 0
	db.Model(&user).Where(user).Count(&cnt)
	return cnt > 0
}

func FindUserOne(user *User) (*User, error) {
	found := &User{}
	db.First(found, user)
	if found.UserID == "" {
		return nil, fmt.Errorf("user %#v not found", user)
	}
	return found, nil
}

func (self *User) getAuthInfo() (*UserAuthInfo, error) {
	found := &UserAuthInfo{}
	db.First(found, &UserAuthInfo{UserID: self.UserID})
	if found.UserID == "" {
		return nil, fmt.Errorf("user_auth_info %#v not found", self)
	}
	return found, nil
}

func (self *User) Authorize(password string) (bool, error) {
	authinfo, err := self.getAuthInfo()
	if err != nil {
		return false, err
	}

	hashed := hashPassword(password, authinfo.Salt)
	if hashed == authinfo.HashedPassword {
		return true, nil
	}
	return false, nil
}

func (self *User) ReassignToken() string {
	newtkn := helpers.RandomStr(30)
	auth, err := self.getAuthInfo()
	if err != nil {
		panic(err)
	}

	auth.Token = newtkn
	db.Save(auth)
	return newtkn
}
