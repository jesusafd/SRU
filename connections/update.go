package connections

import (
	"context"
	"log"
	"time"

	"github.com/jesusafd/SRU/models"
)

// Update Actualiza los datos de un registro, excluyendo
// llave primaria (email) y contraseña
func Update(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := "UPDATE user SET name=?,last_name=?,year_birth=?,month_birth=?,day_birth=? WHERE email=?"
	var err error
	_, err = DBConn.ExecContext(
		ctx,
		query,
		u.Name,
		u.LastName,
		u.YearBirth,
		u.MonthBirth,
		u.DayBirth,
		u.Email,
	)
	if err != nil {
		log.Println("Error al actulizar el registro en la base de datos : " + err.Error())
	}

	return err
}
