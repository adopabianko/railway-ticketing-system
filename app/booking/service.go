package booking

import (
	"fmt"
	"strconv"

	"github.com/adopabianko/train-ticketing/database"
)

type IBookingService interface {
	BookingService(bookingParam *BookingParam) (int, string, interface{})
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

func (s *BookingService) BookingService(bookingParam *BookingParam) (httpCode int, message string, result interface{}) {
	id := bookingParam.ScheduleId
	depDate := bookingParam.DepartureDate
	qty := bookingParam.Qty
	custCode := bookingParam.CustomerCode
	passengers := bookingParam.Passengers

	// Check total passengers and qty
	if uint16(len(passengers)) != qty {
		return 422, "Total passengers and qty is not same", nil
	}

	// Get schedule
	resultSC, statusSC := s.Repository.FindScheduleByIdRepo(id)

	// Schedule validation
	if !statusSC {
		return 404, "Schedule is not available", nil
	} else if resultSC.Balance == 0 {
		return 422, "Ticket is sold out", nil
	} else if resultSC.Balance < qty {
		return 422, fmt.Sprintf("The only ticket left %d", resultSC.Balance), nil
	} else if depDate >= resultSC.EndDate {
		return 404, "Schedule is not available", nil
	}

	// Get data customer
	resultCS, statusCs := s.Repository.FindCustomerByCustomercodeRepo(custCode)

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
		ScheduleId:    id,
		CustomerId:    resultCS.ID,
		DepartureDate: depDate,
		Qty:           qty,
		Price:         price,
		Total:         total,
	}

	bookingUuid, bookingCode := s.Repository.SaveBookingRepo(&booking)

	// Insert data passenger
	for i, p := range passengers {
		var increment, digit int
		increment = i + 1
		digit = countDigit(increment)

		var ticketNumber string
		if digit == 1 {
			ticketNumber = bookingCode + "00" + strconv.Itoa(increment)
		} else if digit == 2 {
			ticketNumber = bookingCode + "0" + strconv.Itoa(increment)
		} else {
			ticketNumber = bookingCode + strconv.Itoa(increment)
		}

		s.Repository.SavePassengerRepo(bookingUuid, ticketNumber, &p)
	}

	return 200, "Booking success", bookingCode
}

func countDigit(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}
