package booking

type BookingRequest struct {
	ScheduleId    string      `json:"schedule_id"`
	CustomerCode  string      `json:"customer_code"`
	Qty           uint16      `json:"qty"`
	DepartureDate string      `json:"departure_date"`
	Passengers    []Passenger `json:"passengers"`
}

type Booking struct {
	ID            string  `json:"id"`
	ScheduleId    string  `json:"schedule_id,omitempty"`
	CustomerId    string  `json:"customer_id,omitempty"`
	BookingCode   string  `json:"booking_code"`
	DepartureDate string  `json:"departure_date"`
	Qty           uint16  `json:"qty"`
	Price         float32 `json:"price"`
	Total         float32 `json:"total"`
	ExpiredDate   string  `json:"expired_date,omitempty"`
	BookingStatus string  `json:"booking_status"`
	BookedAt      string  `json:"booked_at"`
	Customer      Customer  `json:"customer"`
	Schedule      Schedule  `json:"schedule"`
}

type Schedule struct {
	ID          string  `json:"id,omitempty"`
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	TrainCode   string  `json:"train_code"`
	Time        string  `json:"time"`
	Quota       uint16  `json:"quota"`
	Balance     uint16  `json:"balance"`
	Price       float32 `json:"price"`
	StartDate   string  `json:"start_date,omitempty"`
	EndDate     string  `json:"end_date,omitempty"`
}

type Customer struct {
	ID           string `json:"id,omitempty"`
	CustomerCode string `json:"customer_code"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
}

type Passenger struct {
	ID           string `json:"id,omitempty"`
	BookingID    string `json:"booking_id,omitempty"`
	TicketNumber string `json:"ticket_number"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
}
