package handler_test

import (
	"ardeolib.sapions.com/models"
	"github.com/gocql/gocql"
)

var (
	Roadmap1 = models.RoadMap{
		UUID:        gocql.UUID{0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		UserUUID:    gocql.UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		Title:       "Roadmap 1",
		Description: "Description 1",
	}
	Roadmap2 = models.RoadMap{
		UUID:        gocql.UUID{0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02},
		UserUUID:    gocql.UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		Title:       "Roadmap 2",
		Description: "Description 2",
	}
)