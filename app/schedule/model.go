package schedule

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
