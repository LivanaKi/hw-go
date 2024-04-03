package user

type User struct {
	Name     string // 'json:"name"' //будет из за большой букви, тоесть публичного индефикатора
	password string //'json:"password"'// не будет работать, из-за приватного индефикатора
}

func (m *User) SetPass(psw string) {
	m.password = psw
}
