package services

import (
	"fmt"

	"ardeolib.sapions.com/models"
	"ardeolib.sapions.com/repository"
	"ardeolib.sapions.com/utils"
)

type actSvcType int

const (
	createActT actSvcType = iota
	createActsT
	readActT
	updateActT
	deleteActT
	deleteActsT
)

var actServiceName = map[actSvcType]string{
	createActT:  "Criar",
	createActsT: "CriarN",

	readActT:    "Ler",
	updateActT:  "Atualizar",
	deleteActT:  "Deletar",
	deleteActsT: "DeletarN",
}

type actService struct {
	act  *models.Act
	usr  *models.User
	rdm  *models.RoadMap
	rps  *repository.Repository
	acts *[]models.Act
	_T   actSvcType
}

func (s *actService) Run() error {
	switch s._T {
	case createActT:
		return s.createAct()
	case createActsT:
		return s.createActs()
	case readActT:
		return s.readAct()
	case updateActT:
		return s.updateAct()
	case deleteActT:
		return s.deleteAct()
	case deleteActsT:
		return s.deleteActs()
	default:
		return fmt.Errorf("erro no processamento de atividade, serviço de [%s]", actServiceName[s._T])
	}
}

func NewCreateActService(act *models.Act, usr *models.User, rps *repository.Repository) *actService {
	return &actService{act: act, usr: usr, rps: rps, _T: createActT}
}

func (s *actService) createAct() (err error) {
	err = s.rps.InsertAct(s.act)
	return err
}

func NewCreateActsService(acts *[]models.Act, usr *models.User, rps *repository.Repository) *actService {
	return &actService{acts: acts, usr: usr, rps: rps, _T: createActsT}
}

func (s *actService) createActs() (err error) {

	err = s.rps.InsertActs(s.acts)
	return err
}

func NewReadActService(acts *[]models.Act, rdm *models.RoadMap, usr *models.User, rps *repository.Repository) *actService {
	return &actService{acts: acts, usr: usr, rdm: rdm, rps: rps, _T: deleteActsT}
}

func (s *actService) readAct() (err error) {
	if err = utils.CheckIfUserIsActsOwner(s.usr, s.acts, s.rps); err != nil {
		return err
	}
	return s.rps.GetActByRoadMapUUID(s.acts, s.act.RoadMapUUID)
}

func NewUpdateActService(act *models.Act, usr *models.User, rps *repository.Repository) *actService {
	return &actService{act: act, usr: usr, rps: rps, _T: updateActT}
}

func (s *actService) updateAct() (err error) {
	if err = utils.CheckIfUserIsActsOwner(s.usr, s.acts, s.rps); err != nil {
		return err
	}
	return s.rps.UpdateAct(s.act)
}

func NewDeleteActService(act *models.Act, usr *models.User, rps *repository.Repository) *actService {
	return &actService{act: act, usr: usr, rps: rps, _T: deleteActT}
}

func (s *actService) deleteAct() (err error) {
	if err = utils.CheckIfUserIsActsOwner(s.usr, s.acts, s.rps); err != nil {
		return err
	}
	return s.rps.DeleteAct(s.act)
}

func NewDeleteActsService(rdm *models.RoadMap, usr *models.User, rps *repository.Repository) *actService {
	return &actService{rdm: rdm, usr: usr, rps: rps, _T: deleteActsT}
}

func (s *actService) deleteActs() (err error) {

	if err = utils.CheckIfUserIsActsOwner(s.usr, s.acts, s.rps); err != nil {
		return err
	}

	return s.rps.DeleteActs(s.acts, s.rdm.UUID)
}
