package repositories

import (
	"database/sql"
	"go_dnd/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT * FROM usuario")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Find(u *models.User) ([]models.User, error) {
	query_str := "SELECT * FROM usuarios WHERE id=?"

	rows, err := r.db.Query(query_str, u.ID, u.Name, u.Email)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Save(user *models.User) error {
	_, err := r.db.Exec("INSERT INTO usuario(name, email) VALUES(?,?)", user.Name, user.Email)
	return err
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	rows, err := r.db.Query("SELECT * FROM usuario WHERE id=?", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	_, err := r.db.Exec("UPDATE name=?, email=? WHERE id=?", &user.Name, &user.Email, &user.ID)
	return err
}
