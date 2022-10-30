package controllers

import (
	"net/http"

	"github.com/MatiasCordich/Golang-api/models"
	"github.com/MatiasCordich/Golang-api/services"
	"github.com/gin-gonic/gin"
)

/*
Vamos a crear el struct de tipo UserController que va a contener
el struct de UserService que vamos a importar
*/

type UserController struct {
	UserService services.UserService
}

/*
Vamos a crear el constructor UserController
Le voy a pasar por parametro el struct de UserService
*/

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

/*
Estas funciones son parecidas a las funcioes de services pero no,
ya que van a tener como funcion la de manipular la informacion dependiendo de lo que
queramos hacer en las rutas.

Voy a recibir por parametro un objeto (context) cuya informacion va a ser guardado
mediante el gin.Context, dicha informacion la enviaremos como una request
y la respuesta va a ser el ctx parseado a JSON
*/

func (uc *UserController) CreateUser(ctx *gin.Context) {

	// Creamos nuestra variable user que comprende el objeto de tipo User

	var user models.User

	// Validamos

	/*
		Si hay un error mientras bindeo el context con el usuario mostrame un msg de error
	*/

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//Si no hay error entonces llamamos de userService el metodo CreateUser y la pasamos la variable user que creamos

	err := uc.UserService.CreateUser(&user)

	// Validamos, si el error es diferente a null, es decir hubo un error en la creacion del usuario, enviame un mensaje de error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Si no hay error alguno enviame un mesaje de que el usuario fue creado

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {

	// Creamos la variable username, va a ser igual el param del ctx que es "name"

	username := ctx.Param("dni")

	// Llamame a UserService y que utilice la funcion GetUser que le voy a pasar
	// la variable username

	user, err := uc.UserService.GetUser(&username)

	// Valido si hay algun error en el proceso

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Si todo esta bien devolve el usuario encontrado

	ctx.JSON(http.StatusOK, user)

}

func (uc *UserController) GetAll(ctx *gin.Context) {

	// Creo mis variables users y err cuyo valor es al llamar userService y a la funcion getAll()

	users, err := uc.UserService.GetAll()

	// Valido si hay algun error en el proceso

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Si todo esta bien devolve el usuario encontrado

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

	var user models.User

	// Si hay un error mientras bindeo el context con el usuario mostrame un msg de error

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.UpdateUser(&user)

	// Valido si hay algun error en el proceso

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Si todo esta bien devolve el usuario encontrado

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	// Creo mi variable username que va a a ser el context con parametro "name"

	username := ctx.Param("dni")

	// Creo mi variable error que va a ser el resultado del proceso de eliminar el usario

	err := uc.UserService.DeleteUser(&username)

	// Valido si hay algun error en el proceso

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Si no hay error alguno enviame un mesaje de que el usuario fue eliminado

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Vamos a crear las rutas

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {

	// Creamos la base de las rutas que van a compartir el mismo path

	userroute := rg.Group("/user")

	// Creamos las rutas con sus respectivos metodos

	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:dni", uc.GetUser)
	userroute.GET("/getAll", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete/:dni", uc.DeleteUser)
}
