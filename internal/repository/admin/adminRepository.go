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
}

func (a *AdminRepository) CreateAdmin(admin models.Admin) error {
	stmt := `INSERT INTO  admins (id,email,password,created_at)
	VALUES(?,?,?, datetime('now', 'localtime'));`
	if _, err := a.db.Exec(stmt, admin.ID, admin.Email, admin.Password); err != nil {
		return fmt.Errorf("Create admin error - %v", err)
	}
	return nil
}

func (*AdminRepository) GetAdminByEmail(email string) (models.Admin, error) {
	return models.Admin{}, nil
}
