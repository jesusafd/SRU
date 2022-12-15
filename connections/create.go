package connections

import (
	"context"
	"log"
	"time"

	"github.com/jesusafd/SRU/models"
)

// Create es la funcion encargada de insertar los usuarios en la base de datos
func Create(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := "INSERT INTO user (email,password,name,last_name,year_birth,month_birth,day_birth) VALUES (?,?,?,?,?,?,?)"
	var err error
	// Encriptamos la contraseña
	u.Password, err = Encrypt(u.Password)
	if err != nil {
		log.Println("Error al encriptar la contraseña")
	}
	_, err = DBConn.ExecContext(
		ctx,
		query,
		u.Email,
		u.Password,
		u.Name,
		u.LastName,
		u.YearBirth,
		u.MonthBirth,
		u.DayBirth,
	)
	if err != nil {
		log.Println("Error al insertar el registro en la base de datos : " + err.Error())
	}

	return err
}
