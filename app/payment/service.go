package payment

import "github.com/adopabianko/train-ticketing/database"

type IPaymentService interface {
	PaymentService(bookingCode string)(int, string)
}

type PaymentService struct {
	Repository IPaymentRepository
}

func InitPaymentService() *PaymentService {
	paymentRepository := new(PaymentRepository)
	paymentRepository.MySQL = &database.MySQLConnection{}

	paymentService := new(PaymentService)
	paymentService.Repository = paymentRepository

	return paymentService
}

func(s *PaymentService) PaymentService(bookingCode string)  (httpCode int, message string){
	s.Repository.UpdatePaymentStatusRepo(bookingCode)

	return 200, "Payment Success"
}
