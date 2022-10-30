package services

import (
	"context"
	"errors"

	"github.com/MatiasCordich/Golang-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

/*
Creo mi constructor para UserServiceImpl
Le vamos a pasar la coleccion de Usuarios
Le vamos a pasar el context object

Lo que vamos a retornar es la interfaz de usuario pero vamos a crear la clase implementada
la cual es UserServiceImpl el cual le pasamos el userColecction y el context (ctx)
*/

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

// Creo mis funciones para mi Api

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(dni *string) (*models.User, error) {

	// Creamos una variable usser que es un struct de User

	var user *models.User

	// Cremoa la variable query que es en base a "dni"

	query := bson.D{bson.E{Key: "user_dni", Value: dni}}

	// La logica para obtener el usuario
	// De userCollection utilizo el metodo FindOnne(), como parametro va el context y el query
	// Y utlizamos Decode() que le pasamos por parametro la variable user

	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)

	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {

	// Creo la variable users que va a ser un slice de objetos usuarios

	var users []*models.User

	// Creo las variables cursor y err
	// el valor de esas variables va a ser la coleccion de usuarios y que me encuentre todo lo que tiene la collecion

	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})

	// Si hay algun error entonces no me hagas nada

	if err != nil {
		return nil, err
	}

	// Si no hay error iterame el cursor

	for cursor.Next(u.ctx) {

		// Creamos la variable usuario

		var user models.User

		// Creoma la variable err que es lo que contiene el cursor y usamos el metodo Decode()

		err := cursor.Decode(&user)

		// Si hay algun error en el proceso detene el proceso y no hagas nada

		if err != nil {
			return nil, err
		}

		// Si no hay error quiro que me agregues cada user al slice de users

		users = append(users, &user)
	}

	// Ahora tenemos que cerrar el cursor

	// Si hay algun error en el proceso cerrame el cursor y que me devuelva un nil y el error

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Caso contrario, termina el cursor con el metodo Close()

	cursor.Close(u.ctx)

	// Si no hay usuarios en el slice de usuarios mostrame un error

	if len(users) == 0 {
		return nil, errors.New("There is no users in the DB")
	}

	// Una vez hecho todo el proceso devolveme lo que obtuve del slice users

	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {

	// Creamos la variable filter que me va a encontrar el usuario filtrado por su user.Dni

	filter := bson.D{bson.E{Key: "user_dni", Value: user.Dni}}

	// Creamos la variable update que me va a permitir actualizar los datos del User

	update := bson.D{
		bson.E{
			Key: "$set", Value: bson.D{
				bson.E{Key: "user_name", Value: user.Name},
				bson.E{Key: "user_surname", Value: user.Surname},
				bson.E{Key: "user_age", Value: user.Age},
				bson.E{Key: "user_address", Value: user.Address}}}}

	// Creo la variable result que va a tener a la collecion de Usauarios le paso por parametros el context, y las variables filter y update

	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)

	// Valido si el el usuario a modificar coincide con el result

	if result.MatchedCount != 1 {
		return errors.New("No matched users found for update")
	}

	return nil
}

func (u *UserServiceImpl) DeleteUser(dni *string) error {

	// Creamos la variable filter que me va a encontrar el usuario filtrado por su name

	filter := bson.D{bson.E{Key: "user_dni", Value: dni}}

	// Creo mi varialbe result que es el resultado de llamar a la coleccion de usarios y que me elimine un usario
	// le voy a pasar por parametro el context y lo que me devuelve la variable filter

	result, _ := u.usercollection.DeleteOne(u.ctx, filter)

	// Valido si el el usuario a eliminar coincide con el result

	if result.DeletedCount != 1 {
		return errors.New("No matched users found for deleted")
	}
	return nil
}
