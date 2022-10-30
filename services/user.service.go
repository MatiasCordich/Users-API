package services

import "github.com/MatiasCordich/Golang-api/models"

// Creo mis acciones que va a hacer la Api

/*
- CreateUser()
Va a recibir por parametro un objeto de tipo User

- GetUser()
Va a recibir por parametro un string
Me va a devolver un objeto de tipo User que es el usuario pasado por parametro

- GetAll
Me va a devolver un Slice de todos los objetos User

- UpdateUser()
Va a recebir por parametro un objeto de tipo User para que sea modificado

- DeleteUser()
Va a recibir un string como parametro
Me va a eliminar el usuario que yo pase por parametro
*/

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
