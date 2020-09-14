package repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
	"kecamatan_app/models"
)

// RegisterStruct ...
type RegisterStruct struct {
	db *gorm.DB
}

// RegisterRepoInterface ...
type RegisterRepoInterface interface {
	AddRole(data *models.Role) (*models.Role, error)
	RegisterAdmin(data *models.User) (*models.User, error)
	CreateAuth(username string, role string) (*models.Auth, error)
	DeleteAuthData(givenUuid string) (int, error)
	GetAllAdmin() (*models.ResponseGetAllAdmin, error)
	GetAdminByID(id int) (*models.User, error)
	GetAdminByUUID(uuid string) (*models.User, error)
	GetAdminByUsername(username string) (*models.User, error)
	GetUserRoleByID(userID int) (*models.UserRole, error)
	GetRoleByID(id int) (*models.Role, error)
	UpdateAdmin(uuid string, data *models.User) (*models.User, error)
	DeleteAdmin(uuid string) error
	BeginTrans() *gorm.DB
	CheckData(username string) bool
	CreateUserRole(data *models.UserRole)
}

// CreateRegisterRepoImpl ...
func CreateRegisterRepoImpl(db *gorm.DB) RegisterRepoInterface {
	return &RegisterStruct{db}
}

// RegisterAdmin ...
func (r *RegisterStruct) AddRole(data *models.Role) (*models.Role, error) {
	tx := r.db.Begin()

	if err := tx.Debug().Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[RegisterAdminRepoImpl.Insert] Error when query save data with : %w", err)
	}

	tx.Commit()

	return data, nil
}

// RegisterAdmin ...
func (r *RegisterStruct) RegisterAdmin(data *models.User) (*models.User, error) {
	tx := r.db.Begin()

	if err := tx.Debug().Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[RegisterAdminRepoImpl.Insert] Error when query save data with : %w", err)
	}

	tx.Commit()

	return data, nil
}

// RegisterAdmin ...
func (r *RegisterStruct) CreateUserRole(data *models.UserRole) {
	tx := r.db.Begin()

	if err := tx.Debug().Create(&data).Error; err != nil {
		tx.Rollback()
		fmt.Errorf("[RegisterAdminRepoImpl.Insert] Error when query save data with : %w", err)
		return
	}

	tx.Commit()

	return
}

// GetByUUID ...
func (r *RegisterStruct) GetAdminByUUID(uuid string) (*models.User, error) {
	data := models.User{}
	err := r.db.Debug().Where("uuid = ? and is_delete=?", uuid, 1).Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[StudentRepo.GetTabunganByID] Error when query get by id with error: %w", err)
	}
	return &data, nil
}

// GetByUsername ...
func (r *RegisterStruct) GetAdminByUsername(username string) (*models.User, error) {
	data := models.User{}
	err := r.db.Debug().Where("username = ? and is_delete=?", username, 1).Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[StudentRepo.GetTabunganByID] Error when query get by id with error: %w", err)
	}
	return &data, nil
}

// GetByUserRoleBYUserID ...
func (r *RegisterStruct) GetUserRoleByID(userID int) (*models.UserRole, error) {
	data := models.UserRole{}
	err := r.db.Debug().Where("user_id = ?", userID).Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[StudentRepo.GetTabunganByID] Error when query get by id with error: %w", err)
	}
	return &data, nil
}

// GetByUserRoleBYUserID ...
func (r *RegisterStruct) GetRoleByID(id int) (*models.Role, error) {
	data := models.Role{}
	err := r.db.Debug().Where("id = ? and is_delete=?", id, 1).Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[StudentRepo.GetTabunganByID] Error when query get by id with error: %w", err)
	}
	return &data, nil
}

// GetAdminByID ...
func (r *RegisterStruct) GetAdminByID(id int) (*models.User, error) {
	data := models.User{}
	err := r.db.Debug().Where("id = ? and is_delete=?", id, 1).Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[StudentRepo.GetTabunganByID] Error when query get by id with error: %w", err)
	}
	return &data, nil
}

// GetAdmin ...
func (r *RegisterStruct) GetAllAdmin() (*models.ResponseGetAllAdmin, error) {
	tx := r.db.Begin()

	var result models.ResponseGetAllAdmin
	var data models.GetAdmin
	var dataGetAllAdmin []models.GetAdmin

	rows, err := r.db.Debug().Raw("SELECT c.uuid, c.username, b.role_name, c.is_delete from tb_user_role a, tb_role b, tb_user c where a.user_id = c.id and a.role_id = b.id and c.is_delete=1;").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data.UUID, &data.Username, &data.Role, &data.IsDelete)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("[StudentRepo.GetByID] Error when query GetByID data with error: %w", err)
		}
		dataGetAllAdmin = append(dataGetAllAdmin, data)
	}

	result.Data = dataGetAllAdmin

	tx.Commit()

	return &result, nil
}

// UpdateAdmin ...
func (r *RegisterStruct) UpdateAdmin(uuid string, data *models.User) (*models.User, error) {
	tx := r.db.Begin()

	err := r.db.Debug().Model(&data).Where("uuid = ? AND is_delete=?", uuid, 1).Update(data).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[StudentRepo.Update] Error when query update data with error: %w", err)
	}

	tx.Commit()
	return data, nil
}

// DeleteAdmin ...
func (r *RegisterStruct) DeleteAdmin(uuid string) error {
	data := models.User{}
	err := r.db.Debug().Model(&data).Where("uuid = ? AND is_delete=?", uuid, 1).Update("is_delete", 0).Error
	if err != nil {
		return fmt.Errorf("[StudentRepo.Delete] Error when query delete data with error: %w", err)
	}

	return nil
}

// CreateAuth ...
func (r *RegisterStruct) CreateAuth(username string, role string) (*models.Auth, error) {
	au := &models.Auth{}
	tx := r.db.Begin()

	au.AuthUUID = uuid.NewV4().String() //generate a new UUID each time
	au.Username = username
	au.Role = role
	err := r.db.Create(&au).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	fmt.Println("Insert data to database success")
	return au, nil
}

// DeleteAuthData ...
func (r *RegisterStruct) DeleteAuthData(givenUuid string) (int, error) {
	au := &models.Auth{}
	deleted := r.db.Debug().Where("auth_uuid = ?", givenUuid).Delete(&au)
	if deleted.Error != nil {
		return 0, deleted.Error
	}
	fmt.Println("Delete data from database success")
	return 0, nil
}

//BeginTrans ...
func (r *RegisterStruct) BeginTrans() *gorm.DB {
	return r.db.Begin()
}

// CheckData ...
func (r *RegisterStruct) CheckData(username string) bool {
	var total int

	r.db.Debug().Table("tb_user").Where("username = ? and is_delete = ?", username, 1).Count(&total)
	if total > 0 {
		return true
	}
	return false
}
