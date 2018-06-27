package user

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/daniel/PruebaBackend/UserMicroservice/models"

	"github.com/labstack/gommon/log"
	"github.com/daniel/PruebaBackend/UserMicroservice/commons"
	"crypto/sha256"
	"fmt"
	"github.com/daniel/PruebaBackend/UserMicroservice/connections"
)

func Index(context echo.Context) error{
	users := make([]models.User, 0)
	db := connections.GetConnection()

	db.Find(&users)

	/*for _, value := range users{
		fmt.Println(value)
	}*/

	if len(users) == 0 {
		return commons.DisplayMessage(
			http.StatusOK,
			"No se encontraron registros en la base de datos",
			"no-content",
			context,
		)
	}

	return commons.DisplayMessage(
		http.StatusOK,
		"La petición fue realizada exitosamente",
		users,
		context,
	)
}

func Create(context echo.Context) error{
	user := new(models.User)
	db := connections.GetConnection()

	if err := context.Bind(user); err != nil{
		log.Printf("error: %s", err)

		return commons.DisplayMessage(
			http.StatusBadRequest,
			"La peticion no envio los parametros correctos para registrar el usuario",
			"no-content",
			context,
		)
	}

	if user.Nombre == "" || user.Email == "" || user.Password == "" || user.Telefono == ""{
		return commons.DisplayMessage(
			http.StatusUnauthorized,
			"El usuario que se desea registrar no envia datos importantes como el nombre,email, contraseña y telefono",
			"no-content",
			context,
		)
	}

	hash := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", hash )
	user.Password = pwd

	err := db.Create(&user).Error

	if err != nil {
		return commons.DisplayMessage(
			http.StatusBadRequest,
			"A ocurrido un error al momento de guardar el usuario: \n Talvéz el email o el número de telefono que se envio ya estaba registrado",
			"no-content",
			context,
		)
	}

	return commons.DisplayMessage(
		http.StatusCreated,
		"El usuario fue creado exitosamente",
		"no-content",
		context,
	)
}