package domain

type Invoice struct {
	Bookings   []Booking
	CustomerId int    `json:"customerId"`
	Id         int    `json:"id"`
	Month      int    `json:"month"`
	Status     string `json:"status"`
	Year       int    `json:"year"`
}

type Booking struct {
	ActivityId int     `json:"activityId"`
	Day        int     `json:"day"`
	Hours      float32 `json:"hours"`
	Id         int     `json:"id"`
	InvoiceId  int     `json:"invoiceId"`
	ProjectId  int     `json:"projectId"`
}

type Customer struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Projects []Project `json:"projects"`
	UserId   int       `json:"userId"`
}

type Project struct {
	CustomerId int    `json:"customerId"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
}

type Activity struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"userId"`
}

type Rate struct {
	Price      float32 `json:"price"`
	ProjectId  int     `json:"projectId"`
	ActivityId int     `json:"activityId"`
}
