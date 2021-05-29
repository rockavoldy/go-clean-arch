package repository

import (
	"database/sql"
	"go-clean-arch/entity"
	"time"
)

// sqlite repo for User entity
type UserMySQL struct {
	db *sql.DB
}

// create new MySQL repository
func NewUserMySQL(db *sql.DB) *UserMySQL {
	return &UserMySQL{
		db: db,
	}
}

// create User
func (r *UserMySQL) Create(e *entity.User) (entity.ID, error) {
	stmt, err := r.db.Prepare(`INSERT INTO user (id, email, password, name, created_at) values (?,?,?,?,?)`)
	defer stmt.Close()
	if err != nil {
		return e.ID, err
	}

	_, err = stmt.Exec(
		e.ID,
		e.Email,
		e.Password,
		e.Name,
		time.Now().Format(entity.FormatDateTimeSQL),
	)
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

// get user
func (r *UserMySQL) Get(id entity.ID) (*entity.User, error) {
	rows, err := r.db.Query(`SELECT id, email, name, created_at FROM user WHERE id = ? AND deleted_at IS NULL`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user entity.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)
	}
	if user.Email == "" {
		return nil, entity.ErrNotFound
	}

	rows, err = r.db.Query(`SELECT book_id FROM book_user WHERE user_id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var i entity.ID
		err = rows.Scan(&i)
		user.Books = append(user.Books, i)
	}

	return &user, nil
}

// list user
func (r *UserMySQL) List() ([]*entity.User, error) {
	rows, err := r.db.Query(`SELECT id, email, name, created_at FROM user WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		_ = rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)

		stmt, err := r.db.Prepare(`SELECT book_id FROM book_user WHERE user_id = ?`)
		if err != nil {
			return nil, err
		}

		rowsBook, err := stmt.Query(user.ID)
		if err != nil {
			return nil, err
		}

		for rowsBook.Next() {
			var i entity.ID
			_ = rowsBook.Scan(&i)
			user.Books = append(user.Books, i)
		}

		users = append(users, &user)
	}

	return users, nil
}

// Search User by name
func (r *UserMySQL) Search(query string) ([]*entity.User, error) {
	stmt, err := r.db.Prepare(`SELECT id, email, name, created_at FROM user WHERE name LIKE ? AND delete_at IS NULL`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		_ = rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)

		stmt, err := r.db.Prepare(`SELECT book_id FROM book_user WHERE user_id = ?`)
		if err != nil {
			return nil, err
		}

		rowsBook, err := stmt.Query(user.ID)
		if err != nil {
			return nil, err
		}

		for rowsBook.Next() {
			var i entity.ID
			_ = rowsBook.Scan(&i)
			user.Books = append(user.Books, i)
		}

		users = append(users, &user)
	}

	return users, nil
}

// Update an User
func (r *UserMySQL) Update(e *entity.User) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`UPDATE user SET email = ?, password = ?, name = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL`,
		e.Email, e.Password, e.Name, e.UpdatedAt.Format(entity.FormatDateTimeSQL), e.ID)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`DELETE FROM book_user WHERE user_id = ?`, e.ID)
	if err != nil {
		return err
	}

	for _, b := range e.Books {
		_, err := r.db.Exec(`INSERT INTO book_user VALUES (?,?,?)`, e.ID, b, time.Now().Format(entity.FormatDateTimeSQL))
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete an User
func (r *UserMySQL) Delete(id entity.ID) error {
	_, err := r.db.Exec(`UPDATE user SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL`, time.Now().Format(entity.FormatDateTimeSQL), id)
	if err != nil {
		return err
	}

	return nil
}

// Get deleted User
func (r *UserMySQL) GetDeletedUser(id entity.ID) (*entity.User, error) {
	rows, err := r.db.Query(`SELECT id, email, name, created_at FROM user WHERE id = ? AND deleted_at IS NOT NULL`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user entity.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.DeletedAt)
	}

	rows, err = r.db.Query(`SELECT book_id FROM book_user WHERE user_id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var i entity.ID
		err = rows.Scan(&i)
		user.Books = append(user.Books, i)
	}

	return &user, nil
}

// Restore an User
func (r *UserMySQL) Restore(id entity.ID) error {
	_, err := r.db.Exec(`UPDATE user SET deleted_at = null WHERE id = ? AND deleted_at IS NOT NULL`, id)
	if err != nil {
		return err
	}

	return nil
}
