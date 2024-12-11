package models

import "github.com/gocql/gocql"

type ActType int16

const (
	OQNAActT ActType = iota
)

type User struct {
	/*
		Email é um único identificador do usuário
	*/
	Email string `cql:"email" json:"email"`
	/*
		UUID é o identificador único do usuário utilizado para organizar os dados relacionados a um usuário
	*/
	UUID gocql.UUID `cql:"uuid"  json:"uuid"`
	/*
		isProfessor é utilizado para saber se é um professor ou um estudante
	*/
	IsProfessor bool `cql:"is_professor" json:"isProfessor"`
	/*
		Name é o nome do usuário
	*/
	Name string `cql:"name" json:"name"`
	/*
		PasswordHash armazena o hash gerado através da senha do usuário
	*/
	PasswordHash string `cql:"passwordHash" json:"passwordHash"`
}

type OQNAAct struct { // One question N Answer
	/*
		UserUUID é o identificador de usuário para relacionar a atividade com o usuário
	*/
	UserUUID gocql.UUID `cql:"user_uuid" json:"userUUID"`
	/*
		Question é uma pergunta armazenada
	*/
	Question string `cql:"question" json:"question"`
	/*
		UUID é utilizado para referenciar unicamente uma OQNAAct
	*/
	UUID gocql.UUID `cql:"uuid" json:"uuid"`
}

type Answer struct {
	/*
		OQNAActUUID é o UUID de um OQNAAct
	*/
	OQNAActUUID gocql.UUID `cql:"o_q_n_a_act_uuid" json:"oQNAActUUID"`
	/*
		UUID é o identificador único do Answer
	*/
	UUID gocql.UUID `cql:"uuid" json:"uuid"`
	/*
		Answer é o texto de uma resposta
	*/
	Answer string `cql:"answer" json:"answer"`
	/*
		IsCorrect é um valor entre 1 e 0 que no geral ao ser inserido deve ser normatizado
	*/
	IsCorrect bool `cql:"is_correct" json:"isCorrect"`
}

type RoadMap struct {
	/*
		O Road Map é utilizado para criar um roteiro de estudos, UUID é o identificador unico
	*/
	UUID gocql.UUID `cql:"uuid" json:"uuid"`
	/*
		UserUUID é o Identificador do usuário que criou
	*/
	UserUUID gocql.UUID `cql:"user_uuid" json:"userUUID"`
	/*
		Title é o título do Road Map
	*/
	Title string `cql:"title" json:"title"`
	/*
		Description é uma descrição breve do Road Map
	*/
	Description string `cql:"description" json:"description"`
}

type Act struct {
	/*
		Descreve a qual Road Map estará associada a atividade
	*/
	RoadMapUUID gocql.UUID `cql:"road_map_uuid" json:"roadMapUUID"`
	/*
		Descreve o identificador da Actividade
	*/
	ActUUID gocql.UUID `cql:"act_uuid" json:"actUUID"`
	/*
		Descreve o tipo da atividade
	*/
	ActType ActType `cql:"act_type" json:"actType"`
	/*
		Descreve a qual etapa do Road Map a atividade estará associada
	*/
	Stage int `cql:"stage" json:"stage"`
}
