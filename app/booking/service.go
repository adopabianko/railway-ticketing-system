package booking

import (
	"github.com/adopabianko/train-ticketing/database"
)

type IBookingService interface {
	BookingService(id string, depDate string, qty uint16, csCode string) (int, string, interface{})
}

type BookingService struct {
	Repository IBookingRepository
}

func InitBookingService() *BookingService {
	bookingRepository := new(BookingRepository)
	bookingRepository.MySQL = &database.MySQLConnection{}
	bookingRepository.Redis = &database.RedisConnection{}

	bookingService := new(BookingService)
	bookingService.Repository = bookingRepository

	return bookingService
}

func (s *BookingService) BookingService(id string, depDate string, qty uint16, csCode string) (httpCode int, message string, result interface{}) {
	// Get schedule
	resultSC, statusSC := s.Repository.FindScheduleByIdRepo(id)

	// Schedule validation
	if !statusSC {
		return 404, "Schedule is not available", nil
	} else if resultSC.Balance == 0 {
		return 422, "Quota has run out", nil
	} else if resultSC.Balance < qty {
		return 422, "Insufficient quota", nil
	} else if depDate >= resultSC.EndDate {
		return 404, "Schedule is not available", nil
	}

	// Get data customer
	resultCS, statusCs := s.Repository.FindCustomerByCustomercodeRepo(csCode)

	// Customer validation
	if !statusCs {
		return 404, "User is not found", nil
	}

	// Update balance quota
	var balance uint16
	balance = resultSC.Balance - qty

	s.Repository.UpdateBalanceQuotaRepo(id, balance)

	// Insert data booking
	var price, total float32
	price = resultSC.Price
	total = float32(qty) * price

	booking := Booking{
		ScheduleId:      id,
		DepartureDate:   depDate,
		Qty:             qty,
		CustCode:        resultCS.CustomerCode,
		CustFirstName:   resultCS.FirstName,
		CustLastName:    resultCS.LastName,
		CustEmail:       resultCS.Email,
		CustPhoneNumber: resultCS.PhoneNumber,
		Price:           price,
		Total:           total,
	}
	
	bookingCode := s.Repository.SaveBookingRepo(&booking)
	return 200, "Booking success", bookingCode
}
