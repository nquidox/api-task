package user

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Id         uint   `json:"id" gorm:"column:id;"`
	Username   string `json:"username"`
}

func Migrate(repo Repo) error {
	err := repo.Migrate(User{})
	if err != nil {
		return err
	}
	return nil
}

func (u *User) create(repo Repo) error {
	return repo.Create(u)
}

func (u *User) readAll(repo Repo) (*[]User, error) {
	users := &[]User{}
	err := repo.ReadAll(users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) update(repo Repo) error {
	err := repo.Update(u, u.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) delete(repo Repo) error {
	err := repo.Delete(u, u.Id)
	if err != nil {
		return err
	}
	return nil
}
