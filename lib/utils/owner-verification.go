package utils

import (
	"fmt"

	"ardeolib.sapions.com/models"
	"ardeolib.sapions.com/repository"
	"github.com/gocql/gocql"
)

func CheckIfUserIsActsOwner(usr *models.User, acts *[]models.Act, rps *repository.Repository) (err error) {

	// Buscar as oQNAActUUID pertencentes ao usuário
	oQNAActUUID := make([]gocql.UUID, 100)
	err = rps.GetOQNAActUUIDByUserUUID(&oQNAActUUID, usr.UUID)

	for _, act := range *acts {
		if act.ActType == models.OQNAActT {
			_, ok := UUIDBinarySearch(&oQNAActUUID, act.ActUUID)
			if !ok {
				return fmt.Errorf("o usuário não é o dono de todas as atividades")
			}
		}
		// outros tipos
	}
	return err
}

func CheckIfUserIsRoadMapOwner(usr *models.User, roadmap models.RoadMap, rps *repository.Repository) (err error) {

	return err
}
