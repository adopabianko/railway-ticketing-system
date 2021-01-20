package auth

import (
	"database/sql"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/adopabianko/train-ticketing/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IAuthRepository interface {
	FindCustomerExitsByEmailRepo(email string) (bool, Customer)
	FindCustomerRegisteringRepo(id uuid.UUID) (Customer, error)
	SaveRepo(customer *Customer) (bool, uuid.UUID)
	FindCustomerActivationRepo(customerCode, activationCode string) bool
	UpdateStatusActive(customerCode, activationCode string) bool
}

type AuthRepository struct {
	MySQL database.IMySQLConnection
}

func (r *AuthRepository) FindCustomerExitsByEmailRepo(email string) (status bool, customer Customer) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	err := db.QueryRow(`
		SELECT
			id,
		   	customer_code,
		   	first_name,
		   	last_name,
		   	email,
		   	phone_number,
		   	gender,
		   	birth_date,
		   	activation_code,
		   	password
		FROM customer 
		WHERE email = ? 
		AND status_active = 1`, email).Scan(
		&customer.ID,
		&customer.CustomerCode,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.Gender,
		&customer.BirthDate,
		&customer.ActivationCode,
		&customer.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, customer
		}

		log.Fatal(err.Error())
	}

	return true, customer
}

func (r *AuthRepository) FindCustomerRegisteringRepo(id uuid.UUID) (customer Customer, err error) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	err = db.QueryRow(`
		SELECT
			customer_code,
			first_name,
			last_name,
			email,
			phone_number,
			gender,
			birth_date,
		    activation_code
		FROM customer WHERE id = ?
	`, id).Scan(
		&customer.CustomerCode,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.Gender,
		&customer.BirthDate,
		&customer.ActivationCode,
	)

	if err != nil {
		return
	}

	return customer, nil
}

func (r *AuthRepository) SaveRepo(customer *Customer) (status bool, id uuid.UUID) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	uuid := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	activationCode := generateActivationcode()
	customerCode := generateCustomerCode()

	_, err = db.Exec(`INSERT INTO 
    		customer(
    			id,
    			customer_code, 
    			first_name,
    		    last_name,
    			email,
				phone_number,
				gender,
				birth_date,
    			activation_code,
				password
    		) VALUES(?,?,?,?,?,?,?,?,?,?)`,
		uuid,
		customerCode,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.PhoneNumber,
		customer.Gender,
		customer.BirthDate,
		activationCode,
		hashedPassword,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return true, uuid
}

func (r *AuthRepository) FindCustomerActivationRepo(customerCode, activationCode string) (status bool) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	var countCustomer uint8

	err := db.QueryRow(`
		SELECT COUNT(*) 
		FROM customer 
		WHERE customer_code = ?
	  	AND activation_code = ?
		AND status_active = 0`, customerCode, activationCode).Scan(&countCustomer)

	if err != nil {
		log.Fatal(err.Error())
	}

	if countCustomer == 0 {
		return false
	}

	return true
}

func (r *AuthRepository) UpdateStatusActive(customerCode, activationCode string) (status bool) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	_, err := db.Exec(`
			UPDATE customer SET status_active = 1
			WHERE customer_code = ?
			AND activation_code = ?
		`, customerCode, activationCode)

	if err != nil {
		log.Fatal(err.Error())
	}

	return true
}

func generateCustomerCode() string {
	t := time.Now()
	var dateNow string = t.Format("20060102")

	rand.Seed(t.UnixNano())
	chars := []rune("0123456789")
	length := 6

	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return dateNow + b.String()
}

func generateActivationcode() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 6

	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}
