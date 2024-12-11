package services

import (
	"fmt"

	"ardeolib.sapions.com/models"
	"ardeolib.sapions.com/repository"
	"ardeolib.sapions.com/utils"
)

type userService struct {
	usr  *models.User
	pswd *string
	rps  *repository.Repository
	_T   usrSvcType
}

type usrSvcType int

var usrSvcName = map[usrSvcType]string{
	createUserT: "Criar",
	// createActsT: "CriarN",

	readUserT:   "Ler",
	updateUserT: "Atualizar",
	deleteUserT: "Deletar",
	// deleteActsT: "DeletarN",
}

const (
	createUserT usrSvcType = iota
	readUserT
	updateUserT
	deleteUserT
	readUsersTa
)

func (s *userService) Run() error {
	switch s._T {
	case createUserT:
		return s.createUser()
	case readUserT:
		return s.readUser()
	case updateUserT:
		return s.updateUser()
	case deleteUserT:
		return s.deleteUser()
	default:
		return fmt.Errorf("erro no processamento de usuario, serviço de [%s]", usrSvcName[s._T])
	}
}

func NewCreateUserService(usr *models.User, pswd *string, rps *repository.Repository) *userService {

	return &userService{usr: usr, pswd: pswd, rps: rps, _T: createUserT}
}

func (s *userService) createUser() (err error) {
	ok := utils.ValidateUserEmailFormat(s.usr.Email) &&
		utils.ValidateUserPswd(s.pswd)
	if !ok {
		return fmt.Errorf("usuario inválido, verifique o email e a senha")
	}
	return s.rps.InsertUser(s.usr)
}

func NewReadUserService(usr *models.User, rps *repository.Repository) *userService {
	return &userService{usr: usr, rps: rps}
}

func (s *userService) readUser() error {
	return s.rps.GetUserByEmail(s.usr)
}

func NewUpdateUserService(usr *models.User, rps *repository.Repository) *userService {
	return &userService{usr: usr, rps: rps}
}

func (s *userService) updateUser() error {
	return s.rps.UpdateUser(s.usr)
}

func NewDeleteUserService(usr *models.User, rps *repository.Repository) *userService {
	return &userService{usr: usr, rps: rps}
}

func (s *userService) deleteUser() error {
	return s.rps.DeleteUser(s.usr)
}
