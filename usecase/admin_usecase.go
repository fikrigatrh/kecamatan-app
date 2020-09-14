package usecase

import (
	"errors"
	"kecamatan_app/auth"
	"kecamatan_app/models"
	"kecamatan_app/repo"
	"kecamatan_app/utils"
	"log"
	"os"
)

type RegisterService struct {
	registerRepo repo.RegisterRepoInterface
}

type RegisterServiceInterface interface {
	AddRole(data *models.Role) (*models.Role, error)
	RegisterAdmin(data *models.User) (*models.User, error)
	DeleteAuthData(givenUuid string) (int, error)
	CreateAuth(username string, role string) (*models.Auth, error)
	GetAllAdmin() (*models.ResponseGetAllAdmin, error)
	GetAdminByUUID(uuid string) (*models.User, error)
	GetAdminByID(id int) (*models.User, error)
	GetAdminByUsername(username string) (*models.User, error)
	GetUserRoleByID(userID int) (*models.UserRole, error)
	GetRoleByID(id int) (*models.Role, error)
	UpdateAdmin(IDAdmin int, data *models.User) (*models.User, error)
	DeleteAdmin(id int) error
	SignIn(authD models.Auth) (string, error)
	LoginService(username, password string) (string, error)
	CheckData(username string) bool
}

func CreateRegisterServiceImpl(registerRepo repo.RegisterRepoInterface) RegisterServiceInterface {
	return &RegisterService{registerRepo}
}

func (r *RegisterService) AddRole(data *models.Role) (*models.Role, error) {
	user, err :=r.registerRepo.AddRole(data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *RegisterService) LoginService(username, password string) (string, error) {
	var jwt models.TokenStruct
	var tokenResult string
	var authD models.Auth
	var roleResult models.Role
	usernameSU := os.Getenv("USERNAME_SU")
	passwordSU := os.Getenv("PASSWORD_SU")

	if username != usernameSU || password != passwordSU {
		daUsername, err := r.registerRepo.GetAdminByUsername(username)
		if err != nil {
			log.Printf("[login Controller] error when get data username : %v\n", err)
			return "", err
		}

		userRole, _ := r.registerRepo.GetUserRoleByID(daUsername.ID)

		role, _ := r.registerRepo.GetRoleByID(userRole.RoleID)

		roleResult.RoleName = role.RoleName
		isCorrect := utils.ComparePassword(daUsername.Password, []byte(password))

		if isCorrect {
			authData, err := r.registerRepo.CreateAuth(username, role.RoleName)
			if err != nil {
				log.Println("[LoginController] ERRORR IN HERE THREE")
				return "", err
			}

			authD.AuthUUID = authData.AuthUUID
			authD.Username = authData.Username
			authD.Role = authData.Role
		} else {
			log.Println("[ADMIN_USECASE] ERRORR IN HERE ISCORRECT")
			return "", err
		}
	}

	if username == usernameSU && password == passwordSU {
		roleResult.RoleName = "super_admin"

		authData, err := r.registerRepo.CreateAuth(username, roleResult.RoleName)
		if err != nil {
			log.Println("[LoginController] ERRORR IN HERE THREE")
			return "", err
		}

		authD.AuthUUID = authData.AuthUUID
		authD.Username = authData.Username
		authD.Role = authData.Role
	} else {
		log.Println("[ADMIN_USECASE] ERRORR IN HERE SUPERADMIN")
		return "", nil
	}

	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	jwt.Token = token

	tokenResult = jwt.Token

	return tokenResult, err

}

func (r *RegisterService) RegisterAdmin(data *models.User) (*models.User, error) {
	var dataUser models.UserRole
	user, err :=r.registerRepo.RegisterAdmin(data)
	if err != nil {
		return nil, err
	}
	dataUser.UserID = user.ID
	dataUser.RoleID = data.RoleID
	r.registerRepo.CreateUserRole(&dataUser)
	return user, nil
}

func (r *RegisterService) GetAllAdmin() (*models.ResponseGetAllAdmin, error) {
	data, err := r.registerRepo.GetAllAdmin()
	if err != nil {
		return nil, errors.New("Oppss, something error")
	}

	return data, nil
}

func (r *RegisterService) GetAdminByUUID(uuid string) (*models.User, error) {
	admin, err := r.registerRepo.GetAdminByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *RegisterService) GetAdminByUsername(username string) (*models.User, error) {
	admin, err := r.registerRepo.GetAdminByUsername(username)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *RegisterService) GetUserRoleByID(userID int) (*models.UserRole, error) {
	admin, err := r.registerRepo.GetUserRoleByID(userID)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *RegisterService) GetRoleByID(id int) (*models.Role, error) {
	admin, err := r.registerRepo.GetRoleByID(id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *RegisterService) GetAdminByID(id int) (*models.User, error) {
	admin, err := r.registerRepo.GetAdminByID(id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *RegisterService) UpdateAdmin(IDAdmin int, data *models.User) (*models.User, error) {
	firstData, err := r.registerRepo.GetAdminByID(IDAdmin)
	if err != nil {
		return nil, errors.New("studentID does not exist")
	}

	if data.Username == "" {
		data.Username = firstData.Username
	}
	if data.Password == "" {
		data.Password = firstData.Password
	}

	student, err := r.registerRepo.UpdateAdmin(IDAdmin, data)
	if err != nil {
		return nil, err
	}

	return student, nil

}

func (r *RegisterService) DeleteAdmin(id int) error {
	_, err := r.registerRepo.GetAdminByID(id)
	if err != nil {
		return errors.New("UUID UserDB does not exist")
	}

	err = r.registerRepo.DeleteAdmin(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *RegisterService) CreateAuth(username string, role string) (*models.Auth, error) {
	dataAuth, err := r.registerRepo.CreateAuth(username, role)
	if err != nil {
		return nil, err
	}

	return dataAuth, nil
}

func (r *RegisterService) SignIn(authD models.Auth) (string, error) {
	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *RegisterService) DeleteAuthData(givenUuid string) (int, error) {
	return r.registerRepo.DeleteAuthData(givenUuid)
}

func (r *RegisterService) CheckData(username string) bool {
	data := r.registerRepo.CheckData(username)
	return data
}

