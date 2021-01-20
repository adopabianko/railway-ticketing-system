package auth

import (
	"log"

	"github.com/adopabianko/train-ticketing/database"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	RegisterService(customer *Customer) (int, string, Customer)
	ActivationService(customerCode, activationCode string) (int, string)
	LoginService(email, password string) (int, string, Customer)
}

type AuthService struct {
	Repository IAuthRepository
}

func InitAuthService() *AuthService {
	authRepository := new(AuthRepository)
	authRepository.MySQL = &database.MySQLConnection{}

	authService := new(AuthService)
	authService.Repository = authRepository

	return authService
}

func (s *AuthService) RegisterService(customer *Customer) (httpCode int, message string, result Customer) {
	// Validate user is exist
	customerExists, _ := s.Repository.FindCustomerExitsByEmailRepo(customer.Email)
	if customerExists {
		return 422, "User is exists", result
	}

	// Insert data register
	insert, id := s.Repository.SaveRepo(customer)
	if !insert {
		return 500, "Register failed", result
	}

	// Find user data registering
	result, err := s.Repository.FindCustomerRegisteringRepo(id)
	if err != nil {
		log.Fatal(err)
	}

	return 200, "Register success", result
}

func (s *AuthService) ActivationService(customerCode, activationCode string) (httpCode int, message string) {
	// Check user is exist
	customerExists := s.Repository.FindCustomerActivationRepo(customerCode, activationCode)
	if !customerExists {
		return 404, "User is not found"
	}

	// Update status active user
	actived := s.Repository.UpdateStatusActive(customerCode, activationCode)
	if !actived {
		return 500, "Activation failed"
	}

	return 200, "Activation success"
}

func (s *AuthService) LoginService(email, password string) (httpCode int, message string, customer Customer) {
	// Check user is exist
	customerExists, customer := s.Repository.FindCustomerExitsByEmailRepo(email)
	if !customerExists {
		return 404, "User is not found", customer
	}

	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))

	if err != nil {
		return 404, "User is not found", customer
	}

	return 200, "Login success", customer
}
