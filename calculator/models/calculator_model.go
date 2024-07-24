package models

type AddCal struct {
	FirstValue int `json:"Firstvalue"`
	SecondValue  int `json:"Secondvalue"`
	Operation  string `json:"Operation" `
}

type Calculation struct {
	ID        string    `json:"id" db:"id"`
	FirstValue string    `json:"Firstvalue" db:"Firstvalue"`
	SecondValue  string    `json:"Secondvalue" db:"Secondvalue"`
	Operation     string    `json:"Operation" db:"operation"`
	Result int `json:"result" db:"result"`
}

type Responsed struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
