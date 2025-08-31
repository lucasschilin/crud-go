package model

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

type UserDomainInterface interface {
	SetID(string)
	GetID() string
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAge() int8

	EncryptPassword()
}
