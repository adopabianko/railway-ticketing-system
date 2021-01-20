package payment

import (
	"github.com/adopabianko/train-ticketing/database"
	"log"
	"time"
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
	`, dateNow(), bookingCode)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func dateNow() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
