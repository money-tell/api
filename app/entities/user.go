package entities

import "time"

type User struct {
	ID       string    `json:"_"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

/*func HashPassword(password string) error {
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
*/
