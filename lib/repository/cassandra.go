package repository

/*
	CREATE TABLE user (
		email TEXT,
		uuid UUID,
		is_professor BOOLEAN,
		name TEXT,
		password_hash TEXT,
		PRIMARY KEY (email)
	);

	CREATE TABLE OQNAAct (
		user_uuid UUID,
		question TEXT,
		uuid UUID,
		PRIMARY KEY ((user_uuid), question)
	);

	CREATE TABLE answer (
		o_q_n_a_act_uuid UUID,
		answer TEXT,
		uuid UUID,
		is_correct BOOLEAN,
		PRIMARY KEY ((o_q_n_a_act_uuid), answer, uuid)
	);

	CREATE TABLE roadmap (
		user_uuid UUID,
		title TEXT,
		description TEXT,
		uuid UUID,
		PRIMARY KEY ((user_uuid), title)
	);

	CREATE TABLE act (
		road_map_uuid UUID,
		act_uuid UUID,
		act_type SMALLINT, -- Assumindo que ActType é uma int16
		stage INT,
		PRIMARY KEY ((road_map_uuid), act_uuid, act_type, stage)
	);

	CREATE TABLE student_road_map (
		student_uuid UUID,
		roadmap_uuid UUID,
		PRIMARY KEY (student_uuid, roadmap_uuid)
	);
*/

import (
	"strconv"

	"ardeolib.sapions.com/models"
	"github.com/gocql/gocql"
)

type DatabaseSession interface {
	Query(query string, args ...interface{}) *gocql.Query
}

type CassandraSession struct {
	session *gocql.Session
}

func (s *CassandraSession) Query(query string, args ...interface{}) *gocql.Query {
	return s.session.Query(query, args...)
}

type Repository struct {
	db DatabaseSession
}

func NewRepository(db DatabaseSession) *Repository {
	return &Repository{db: db}
}

/*
**************************************************************************
********************************** user **********************************
**************************************************************************
 */

func (r *Repository) InsertUser(user *models.User) (err error) {
	query := "INSERT INTO users (email, uuid, is_professor, name, password_hash) VALUES (?, ?, ?, ?, ?)"
	err = r.db.Query(query, user.Email, user.UUID, user.IsProfessor, user.Name, user.PasswordHash).Exec()
	return err
}

func (r *Repository) GetUserByEmail(user *models.User) (err error) {
	query := "SELECT email, uuid, is_professor, name, password_hash FROM users WHERE email = ?"
	err = r.db.Query(query, user.Email).Scan(&user.Email, &user.UUID, &user.IsProfessor, &user.Name, &user.PasswordHash)
	return err
}

func (r *Repository) DeleteUser(user *models.User) (err error) {
	query := "DELETE FROM users WHERE email = ?"
	err = r.db.Query(query, user.Email).Exec()
	return err
}

func (r *Repository) UpdateUser(user *models.User) (err error) {
	query := "UPDATE users SET "

	if user.Name != "" {
		query += "name = '" + user.Name + "' "
	}

	if user.PasswordHash != "" {
		query += "password_hash = '" + user.PasswordHash + "' "
	}

	query += "WHERE email = '" + user.Email + "'"
	err = r.db.Query(query).Exec()

	return err
}

/*
**************************************************************************
******************************** OQNAAct *********************************
**************************************************************************
 */

func (r *Repository) InsertOQNAAct(act *models.OQNAAct) (err error) {
	query := "INSERT INTO OQNAAct (user_uuid, question, uuid) VALUES (?, ?, ?)"
	err = r.db.Query(query, act.UserUUID, act.Question, act.UUID).Exec()
	return err
}

func (r *Repository) GetOQNAActByUserUUID(act *[]models.OQNAAct, userUuid gocql.UUID) (err error) {
	query := "SELECT user_uuid, question, uuid FROM OQNAAct WHERE user_uuid = ?"
	iter := r.db.Query(query, userUuid).Iter()
	oQNAAct := models.OQNAAct{}
	for iter.Scan(&oQNAAct.UserUUID, &oQNAAct.Question, &oQNAAct.UUID) {
		*act = append(*act, oQNAAct)
	}
	return iter.Close()
}

func (r *Repository) GetOQNAActUUIDByUserUUID(act_uuid *[]gocql.UUID, userUuid gocql.UUID) (err error) {
	query := "SELECT uuid FROM OQNAAct WHERE user_uuid = ?"
	iter := r.db.Query(query, userUuid).Iter()
	oQNAAct := gocql.UUID{}
	for iter.Scan(&oQNAAct) {
		*act_uuid = append(*act_uuid, oQNAAct)
	}
	return iter.Close()
}

func (r *Repository) DeleteOQNAAct(act *models.OQNAAct) (err error) {
	query := "DELETE FROM OQNAAct WHERE user_uuid = ? AND question = ?"
	err = r.db.Query(query, act.UserUUID, act.Question).Exec()
	return err
}

// question deve ser a antiga, a nova deve ser passada na struct
func (r *Repository) UpdateOQNAAct(act *models.OQNAAct, question string) (err error) {
	query := "UPDATE OQNAAct SET"

	if act.Question != "" {
		query += "question = '" + act.Question + "' "
	}

	query += "WHERE user_uuid = '" + act.UserUUID.String() + "' AND question = '" + question + "'"
	err = r.db.Query(query).Exec()

	return err
}

/*
**************************************************************************
******************************** answer **********************************
**************************************************************************
 */

func (r *Repository) InsertAnswer(answer *models.Answer) (err error) {
	query := "INSERT INTO answer (o_q_n_a_act_uuid, answer, uuid, is_correct) VALUES (?, ?, ?, ?)"
	err = r.db.Query(query, answer.OQNAActUUID, answer.Answer, answer.UUID, answer.IsCorrect).Exec()
	return err
}

// oQNAActUUID é a referência para a pergunta
func (r *Repository) GetAnswerByOQNAActUUID(answers *[]models.Answer, oQNAActUUID gocql.UUID) (err error) {
	query := "SELECT o_q_n_a_act_uuid, answer, uuid, is_correct FROM answer WHERE o_q_n_a_act_uuid = ?"
	iter := r.db.Query(query, oQNAActUUID).Iter()
	answer := models.Answer{}
	for iter.Scan(&answer.OQNAActUUID, &answer.Answer, &answer.UUID, &answer.IsCorrect) {
		*answers = append(*answers, answer)
	}
	return iter.Close()
}

func (r *Repository) DeleteAnswer(answer *models.Answer) (err error) {
	query := "DELETE FROM answer WHERE o_q_n_a_act_uuid = ? AND answer = ?"
	err = r.db.Query(query, answer.OQNAActUUID, answer.Answer).Exec()
	return err
}

// A nova answer vem na struct answer, e a antiga na variável answer
func (r *Repository) UpdateAnswer(answer *models.Answer, answerText string) (err error) {
	query := "UPDATE answer SET "

	if answer.Answer != "" {
		query += "answer = '" + answer.Answer + "' "
	}

	query += "WHERE o_q_n_a_act_uuid = '" + answer.OQNAActUUID.String() + "' AND answer = '" + answerText + "'"
	err = r.db.Query(query).Exec()

	return err
}

/*
**************************************************************************
******************************** roadmap *********************************
**************************************************************************
 */

func (r *Repository) InsertRoadMap(roadMap *models.RoadMap) (err error) {
	query := "INSERT INTO roadmap (user_uuid, title, description, uuid) VALUES (?, ?, ?, ?)"
	err = r.db.Query(query, roadMap.UserUUID, roadMap.Title, roadMap.Description, roadMap.UUID).Exec()
	return err
}

func (r *Repository) GetRoadMapByUserUUID(roadMaps *[]models.RoadMap, userUuid gocql.UUID) (err error) {
	query := "SELECT user_uuid, title, description, uuid FROM roadmap WHERE user_uuid = ?"
	iter := r.db.Query(query, userUuid).Iter()
	roadMap := models.RoadMap{}
	for iter.Scan(&roadMap.UserUUID, &roadMap.Title, &roadMap.Description, &roadMap.UUID) {
		*roadMaps = append(*roadMaps, roadMap)
	}
	return iter.Close()
}

func (r *Repository) GetRoadMapUUIDByUserUUID(roadmapUUIDslc *[]gocql.UUID, userUuid gocql.UUID) (err error) {
	query := "SELECT uuid FROM roadmap WHERE user_uuid = ?"
	iter := r.db.Query(query, userUuid).Iter()
	roadmapUUID := gocql.UUID{}
	for iter.Scan(&roadmapUUID) {
		*roadmapUUIDslc = append(*roadmapUUIDslc, roadmapUUID)
	}
	return iter.Close()
}

func (r *Repository) DeleteRoadMap(roadMap *models.RoadMap) (err error) {
	query := "DELETE FROM roadmap WHERE user_uuid = ? AND title = ?"
	err = r.db.Query(query, roadMap.UserUUID, roadMap.Title).Exec()
	return err
}

// title é o valor antigo para o where, o novo valor vem na roadmap
func (r *Repository) UpdateRoadMap(roadMap *models.RoadMap, title string) (err error) {
	query := "UPDATE roadmap SET "

	if roadMap.Title != "" {
		query += "title = '" + roadMap.Title + "' "
	}

	if roadMap.Description != "" {
		query += "description = '" + roadMap.Description + "' "
	}

	query += "WHERE user_uuid = '" + roadMap.UserUUID.String() + "' AND title = '" + title + "'"
	err = r.db.Query(query).Exec()

	return err
}

/*
**************************************************************************
********************************** act ***********************************
**************************************************************************
 */

func (r *Repository) InsertAct(act *models.Act) (err error) {
	query := "INSERT INTO act (road_map_uuid, act_uuid, act_type, stage) VALUES (?, ?, ?, ?)"
	err = r.db.Query(query, act.RoadMapUUID, act.ActUUID, act.ActType, act.Stage).Exec()
	return err
}

// Precisa de revisão
func (r *Repository) InsertActs(act *[]models.Act) (err error) {
	query := "INSERT INTO act (road_map_uuid, act_uuid, act_type, stage) VALUES (?, ?, ?, ?)"
	for _, a := range *act {
		err = r.db.Query(query, a.RoadMapUUID, a.ActUUID, a.ActType, a.Stage).Exec()
		if err != nil {
			return err
		}
	}
	return err
}

func (r *Repository) GetActByRoadMapUUID(acts *[]models.Act, roadMapUUID gocql.UUID) (err error) {
	query := "SELECT road_map_uuid, act_uuid, act_type, stage FROM act WHERE road_map_uuid = ?"
	iter := r.db.Query(query, roadMapUUID).Iter()
	act := models.Act{}
	for iter.Scan(&act.RoadMapUUID, &act.ActUUID, &act.ActType, &act.Stage) {
		*acts = append(*acts, act)
	}
	return iter.Close()
}

func (r *Repository) DeleteAct(act *models.Act) (err error) {
	query := "DELETE FROM act WHERE road_map_uuid = ? AND act_uuid = ? AND act_type = ?"
	err = r.db.Query(query, act.RoadMapUUID, act.ActUUID, act.ActType).Exec()
	return err
}

func (r *Repository) DeleteActs(act *[]models.Act, roadMapUUID gocql.UUID) (err error) {
	query := "DELETE FROM act WHERE road_map_uuid = ? "
	err = r.db.Query(query, roadMapUUID).Exec()
	return err
}

func (r *Repository) UpdateAct(act *models.Act) (err error) {
	query := "UPDATE act SET "

	if act.Stage != 0 {
		query += "stage = " + strconv.FormatInt(int64(act.Stage), 10) + " "
	}

	query += "WHERE road_map_uuid = '" + act.RoadMapUUID.String() + "' AND act_uuid = '" + act.ActUUID.String() + "' AND act_type = '" + strconv.FormatInt(int64(act.ActType), 10) + "'"
	err = r.db.Query(query).Exec()
	return err
}
