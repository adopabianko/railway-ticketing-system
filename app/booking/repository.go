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
	SaveBookingRepo(booking *Booking) (uuid.UUID, string)
	SavePassengerRepo(bookingUuid uuid.UUID, ticketNumber string, passenger *Passenger)
	UpdateBalanceQuotaRepo(id string, balance uint16)
	FindBookingDetailRepo(bookingCode string) (Booking, bool)
	FindPassengerRepo(bookingId string) ([]Passenger, bool)
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
		SELECT id, first_name, last_name, email, phone_number
		FROM customer
		WHERE customer_code = ?
	`, csCode).Scan(
		&customer.ID,
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

func (r *BookingRepository) SaveBookingRepo(booking *Booking) (uuid.UUID, string) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	uuid := uuid.New()
	bookingCode := generateBookingCode()
	expiredDate := expiredDate()

	_, err := db.Exec(`
		INSERT INTO booking(
			id,
			schedule_id,
			customer_id,
			booking_code,
			departure_date,
			qty,
			price,
			total,
			expired_date
		) VALUE(?,?,?,?,?,?,?,?,?)
	`,
		uuid,
		booking.ScheduleId,
		booking.CustomerId,
		bookingCode,
		booking.DepartureDate,
		booking.Qty,
		booking.Price,
		booking.Total,
		expiredDate,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return uuid, bookingCode
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

func (r *BookingRepository) SavePassengerRepo(bookingUuid uuid.UUID, ticketNumber string, passenger *Passenger) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	uuid := uuid.New()

	_, err := db.Exec(`
		INSERT INTO passenger(
			id,
			booking_id,
			ticket_number,
			first_name,
			last_name,
			email,
			phone_number
		) VALUE(?,?,?,?,?,?,?)
	`,
		uuid,
		bookingUuid,
		ticketNumber,
		passenger.FirstName,
		passenger.LastName,
		passenger.Email,
		passenger.PhoneNumber,
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func (r *BookingRepository)FindBookingDetailRepo(bookingCode string) (booking Booking, status bool) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	// Get data booking
	err := db.QueryRow(`
		SELECT
		    aa.id,
			bb.customer_code,
			bb.first_name,
			bb.last_name,
			bb.email,
			bb.phone_number,
			aa.booking_code,
			aa.departure_date,
			cc.origin,
			cc.destination,
			cc.train_code,
			cc.time,
			aa.qty,
			aa.price,
			aa.total,
			aa.booking_status,
			aa.booked_at
		FROM booking as aa
		JOIN customer as bb on aa.customer_id = bb.id
		JOIN schedule as cc on aa.schedule_id = cc.id
		AND aa.booking_code = ?
	`, bookingCode).Scan(
		&booking.ID,
		&booking.Customer.CustomerCode,
		&booking.Customer.FirstName,
		&booking.Customer.LastName,
		&booking.Customer.Email,
		&booking.Customer.PhoneNumber,
		&booking.BookingCode,
		&booking.DepartureDate,
		&booking.Schedule.Origin,
		&booking.Schedule.Destination,
		&booking.Schedule.TrainCode,
		&booking.Schedule.Time,
		&booking.Qty,
		&booking.Price,
		&booking.Total,
		&booking.BookingStatus,
		&booking.BookedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return booking, false
		}

		log.Fatal(err.Error())
	}

	return booking, true
}

func (r *BookingRepository)FindPassengerRepo(bookingId string) (passengers []Passenger, status bool) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	// Get data passenger
	rows, err := db.Query(`
		SELECT 
		    ticket_number,
			first_name,
		    last_name,
		    email,
		    phone_number
		FROM passenger
		where booking_id = ?
		ORDER BY ticket_number ASC
	`, bookingId)

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var ps Passenger
		err = rows.Scan(
			&ps.TicketNumber,
			&ps.FirstName,
			&ps.LastName,
			&ps.Email,
			&ps.PhoneNumber,
		)

		if err != nil {
			log.Fatal(err.Error())
		}

		passengers = append(passengers, ps)
	}

	if len(passengers) == 0 {
		return passengers, false
	}

	return passengers, true
}

func expiredDate() string {
	t := time.Now()
	return t.Add(time.Hour * 3).Format("2006-01-02 15:04:05")
}

func generateBookingCode() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := 10

	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}
