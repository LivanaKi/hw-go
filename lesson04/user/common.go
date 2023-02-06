package user

type User struct {
	Name     string
	password string
}

func (m *User) SetPass(psw string) {
	m.password = psw
}
