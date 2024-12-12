package services

import (
	"fmt"

	"ardeolib.sapions.com/models"
	"ardeolib.sapions.com/repository"
	"ardeolib.sapions.com/utils"
)

type userService struct {
	usr   *models.User
	pswd  *string
	token *string
	jwt   utils.JWTUtil
	rps   *repository.Repository
	_T    usrSvcType
}

type usrSvcType int

var usrSvcName = map[usrSvcType]string{
	createUserT: "Criar",
	// createActsT: "CriarN",

	readUserT:   "Ler",
	updateUserT: "Atualizar",
	deleteUserT: "Deletar",
	signIn:      "Login",
	// deleteActsT: "DeletarN",
}

const (
	createUserT usrSvcType = iota
	readUserT
	updateUserT
	deleteUserT
	readUsersTa
	signIn
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
	case signIn:
		return s.signInUser()
	default:
		return fmt.Errorf("erro no processamento de usuario, código de serviço desconhecido [%v]", s._T)
	}
}

func NewCreateUserService(
	usr *models.User,
	pswd *string,
	rps *repository.Repository,
) *userService {
	return &userService{
		usr:  usr,
		pswd: pswd,
		rps:  rps,
		_T:   createUserT,
	}
}

func (s *userService) createUser() (err error) {
	if !utils.ValidateUserEmailFormat(s.usr.Email) &&
		utils.ValidateUserPswd(s.pswd) {
		return fmt.Errorf("usuario inválido, verifique o email e a senha")
	}
	if s.usr.PasswordHash, err = utils.HashPassword(*s.pswd); err != nil {
		return err
	}

	return s.rps.InsertUser(s.usr)
}

func NewReadUserService(
	usr *models.User,
	rps *repository.Repository,
) *userService {
	return &userService{
		usr: usr,
		rps: rps,
	}
}

func (s *userService) readUser() error {
	return s.rps.GetUserByEmail(s.usr)
}

func NewUpdateUserService(
	usr *models.User,
	rps *repository.Repository,
) *userService {
	return &userService{
		usr: usr,
		rps: rps,
	}
}

func (s *userService) updateUser() error {
	return s.rps.UpdateUser(s.usr)
}

func NewDeleteUserService(
	usr *models.User,
	rps *repository.Repository,
) *userService {
	return &userService{
		usr: usr,
		rps: rps,
	}
}

func (s *userService) deleteUser() error {
	return s.rps.DeleteUser(s.usr)
}

func NewSignInUserService(
	usr *models.User,
	pswd *string,
	token *string,
	jwt utils.JWTUtil,
	rps *repository.Repository,
) *userService {
	return &userService{
		usr:   usr,
		pswd:  pswd,
		token: token,
		jwt:   jwt,
		rps:   rps,
		_T:    signIn,
	}
}

func (s *userService) signInUser() (err error) {
	if err = s.rps.GetUserByEmail(s.usr); err != nil {
		return nil
	}

	if !utils.CheckPassword(*s.pswd, s.usr.PasswordHash) {
		return fmt.Errorf("mismatch password")
	}

	*s.token, err = s.
		jwt.
		GenerateToken(&utils.UserClaims{User: *s.usr.CopyUserWithoutHashPassword()})

	if err != nil {
		return fmt.Errorf("failed to generate token: %v", err)
	}

	return nil
}
