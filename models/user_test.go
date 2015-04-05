package models

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	actual := hashPassword("test", "foobar")
	expected := "GOuD+uRChbA2FXl1XF/7ksklYAsKyDEOj+w0tc7xsb76cUnOl7pWuwP696+uDTyphhgH2HH628qUNciJ0kVG4Q=="
	if actual != expected {
		t.Errorf("got: %s\nwant: %s", actual, expected)
	}
}

func TestCreateUser(t *testing.T) {
	before_each()

	_, _, err := CreateUser("test", "test@test.com", "John", "pass")
	if err != nil {
		t.Error(err)
	}
	_, _, err = CreateUser("test2", "test2@test2.com", "John2", "pass2")
	if err != nil {
		t.Error(err)
	}

	user2, err := FindUserOne(&User{UserID: "test"})
	if err != nil {
		t.Error(err)
		return
	}
	if user2.Name != "John" {
		t.Errorf("name invalid")
	}
}

func TestAuthorize(t *testing.T) {
	before_each()
	user, pass, err := specCreateUser()
	if err != nil {
		t.Error(err)
	}

	ok, err := user.Authorize(pass)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("cannot authorize")
	}
}

func TestReassignToken(t *testing.T) {
	before_each()
	user, _, err := specCreateUser()
	if err != nil {
		t.Error(err)
	}

	_, err = user.getAuthInfo()
	if err != nil {
		t.Error(err)
	}
	newtkn := user.ReassignToken()
	newauth, err := user.getAuthInfo()
	if newauth.Token != newtkn {
		t.Errorf("got: %s\nwant: %s", newauth.Token, newtkn)
	}
}
