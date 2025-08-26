package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
)

type userDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser() *rest_err.RestErr
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
