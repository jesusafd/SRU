package models

type User struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	YearBirth  int    `json:"yearBirth"`
	MonthBirth int    `json:"monthBirth"`
	DayBirth   int    `json:"dayBirth"`
}
