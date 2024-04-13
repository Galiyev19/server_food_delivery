package admin

import (
	"database/sql"
	"fmt"
	"food_delivery/internal/models"
)

type AdminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

type IAdminRepository interface {
	CreateAdmin(admin models.Admin) error
	GetAdminByEmail(email string) (models.Admin, error)
	ChangePassword(email, password string) error
}

// create new admin
func (a *AdminRepository) CreateAdmin(admin models.Admin) error {
	stmt := `INSERT INTO  admins (id,email,password,created_at)
	VALUES(?,?,?, datetime('now', 'localtime'));`

	if _, err := a.db.Exec(stmt, admin.ID, admin.Email, admin.Password); err != nil {
		return fmt.Errorf("Create admin error - %v", err)
	}
	return nil
}

// get admin from db
func (a *AdminRepository) GetAdminByEmail(email string) (models.Admin, error) {
	var admin models.Admin

	stmt := `SELECT * FROM admins WHERE email = ?`
	if err := a.db.QueryRow(stmt, email).Scan(&admin.ID, &admin.Email, &admin.Password, &admin.CreatedAt); err != nil {
		return models.Admin{}, fmt.Errorf("NOT FIND ADMIN THIS EMAIL %s -", email)
	}
	return admin, nil
}

// Update admin data
func (a *AdminRepository) ChangePassword(email, password string) error {
	stmt := `UPDATE admins SET password = ? WHERE email = ?`
	if _, err := a.db.Exec(stmt, password, email); err != nil {
		return fmt.Errorf("wrong password or email")
	}
	return nil
}
