package booking

type BookingRequest struct {
	ScheduleId string `json:"schedule_id"`
	CustomerCode string `json:"customer_code"`
	Qty uint16 `json:"qty"`
	DepartureDate string `json:"departure_date"`
}

type Booking struct {
	ID              string  `json:"id"`
	BookingCode     string  `json:"booking_code"`
	ScheduleId      string  `json:"schedule_id"`
	DepartureDate   string  `json:"departure_date"`
	Qty             uint16  `json:"qty"`
	CustCode        string  `json:"cust_code"`
	CustFirstName   string  `json:"cust_first_name"`
	CustLastName    string  `json:"cust_last_name"`
	CustEmail       string  `json:"cust_email"`
	CustPhoneNumber string  `json:"cust_phone_number"`
	Price           float32 `json:"price"`
	Total           float32 `json:"total"`
	ExpiredDate     string  `json:"expired_date"`
}

type Schedule struct {
	ID          string  `json:"id"`
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	TrainCode   string  `json:"train_code"`
	Time        string  `json:"time"`
	Quota       uint16  `json:"quota"`
	Balance     uint16  `json:"balance"`
	Price       float32 `json:"price"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
}

type Customer struct {
	CustomerCode string `json:"customer_code"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
}
