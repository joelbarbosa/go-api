package entities

import (
	"time"
)

type Credit struct {
	Id                  int       `json:"id"`
	FoundingType        Founding  `json:"foundingType"`
	CashBalance         float32   `json:"cashBalance"`
	MonthlyRevenue      float32   `json:"monthlyRevenue"`
	RequestedCreditLine int       `json:"requestedCreditLine"`
	RequestedDate       time.Time `json:"requestedDate"`
}
