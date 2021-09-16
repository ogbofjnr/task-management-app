package entities

type User struct {
	Model
	UUID      string `db:"uuid"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Status    string `db:"status"`
}

func (u *User) GetFullName() string {
	return u.FirstName + u.LastName
}
