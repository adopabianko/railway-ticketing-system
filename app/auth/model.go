package auth

type Customer struct {
	ID   		   string `json:"id,omitempty"`
	CustomerCode   string `json:"customer_code"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Gender         string `json:"gender"`
	BirthDate      string `json:"birth_date"`
	ActivationCode string `json:"activation_code,omitempty"`
	Password       string `json:"password,omitempty"`
	RepeatPassword string `json:"repeat_password,omitempty"`
}
