package booking

import (
	"database/sql"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/adopabianko/train-ticketing/database"
	"github.com/google/uuid"
)

type IBookingRepository interface {
	FindScheduleByIdRepo(id string) (Schedule, bool)
	FindCustomerByCustomercodeRepo(csCode string) (Customer, bool)
	SaveBookingRepo(booking *Booking) string
	UpdateBalanceQuotaRepo(id string, balance uint16)
}

type BookingRepository struct {
	MySQL database.IMySQLConnection
	Redis database.IRedisConnection
}

func (r *BookingRepository) FindScheduleByIdRepo(id string) (schedule Schedule, status bool) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	err := db.QueryRow(`
		SELECT balance, price, start_date, end_date
		FROM schedule
		WHERE id = ?
	`, id).Scan(
		&schedule.Balance,
		&schedule.Price,
		&schedule.StartDate,
		&schedule.EndDate,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return schedule, false
		}

		log.Fatal(err.Error())
	}

	return schedule, true
}

func (r *BookingRepository) FindCustomerByCustomercodeRepo(csCode string) (customer Customer, status bool) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	err := db.QueryRow(`
		SELECT first_name, last_name, email, phone_number
		FROM customer
		WHERE customer_code = ?
	`, csCode).Scan(
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.PhoneNumber,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return customer, false
		}

		log.Fatal(err.Error())
	}

	return customer, true
}

func (s *BookingRepository) SaveBookingRepo(booking *Booking) string{
	db := s.MySQL.CreateConnection()
	defer db.Close()

	uuid := uuid.New()
	bookingCode := generateBookingCode()
	expiredDate := expiredDate()

	_, err := db.Exec(`
		INSERT INTO booking(
			id,
			booking_code,
			schedule_id,
			departure_date,
			qty,
			cust_code,
			cust_first_name,
			cust_last_name,
			cust_email,
			cust_phone_number,
			price,
			total,
			expired_date
		) VALUE(?,?,?,?,?,?,?,?,?,?,?,?,?)
	`,
		uuid,
		bookingCode,
		booking.ScheduleId,
		booking.DepartureDate,
		booking.Qty,
		booking.CustCode,
		booking.CustFirstName,
		booking.CustLastName,
		booking.CustEmail,
		booking.CustPhoneNumber,
		booking.Price,
		booking.Total,
		expiredDate,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return bookingCode
}

func (r *BookingRepository) UpdateBalanceQuotaRepo(id string, balance uint16) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	_, err := db.Exec(`
		UPDATE schedule SET balance = ? WHERE id = ?
	`, balance, id)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func expiredDate() string {
	t := time.Now()
	return t.Add(time.Hour * 3).Format("2006-01-02 15:04:05")
}

func generateBookingCode() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 10

	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}
