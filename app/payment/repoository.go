package payment

import (
	"github.com/adopabianko/train-ticketing/database"
	"github.com/adopabianko/train-ticketing/utils"
	"log"
)

type IPaymentRepository interface {
	UpdatePaymentStatusRepo(bookingCode string)
}

type PaymentRepository struct {
	MySQL database.IMySQLConnection
}

func (r *PaymentRepository)UpdatePaymentStatusRepo(bookingCode string){
	db := r.MySQL.CreateConnection()
	defer db.Close()

	_, err := db.Exec(`
		UPDATE booking SET booking_status = 'Paid', paid_at = ?
		WHERE booking_code = ?
	`, utils.DateNow(), bookingCode)

	if err != nil {
		log.Fatal(err.Error())
	}
}
