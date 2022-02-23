package service

import (
	"go-api/src/credit/entities"
	"log"
)

// Response from client api CreditLine
type Response struct {
	Status string `json:"status"`
}

// CreditService
type CreditService interface {
	CreditLine(credit entities.Credit) Response
}

type creditService struct{}

// NewCreditService create an instance
func NewCreditService() CreditService {
	return &creditService{}
}

// oneThirdCashBalance for business logic pdf
func oneThirdCashBalance(cachBalance int64) int64 {
	return cachBalance / 3
}

// oneFifthOfMonthlyRevenue for business logic pdf
func oneFifthOfMonthlyRevenue(monthlyRevenue int64) int64 {
	return monthlyRevenue / 5
}

// isRecommendedCreditLine determine if the credit request will be authorized or not
func isRecommendedCreditLine(credit entities.Credit) entities.Status {
	var creditRevenue int64

	if credit.FoundingType == entities.Sme {
		creditRevenue = oneFifthOfMonthlyRevenue(int64(credit.MonthlyRevenue))
	} else if credit.FoundingType == entities.Startup {
		creditRevenue = oneThirdCashBalance(int64(credit.CashBalance))
		monthlyRevenue := oneFifthOfMonthlyRevenue(int64(credit.MonthlyRevenue))

		if monthlyRevenue > creditRevenue {
			creditRevenue = monthlyRevenue
		}
	}

	if creditRevenue > int64(credit.RequestedCreditLine) {
		return entities.Accepted
	}

	return entities.Rejected
}

// CreditLine service
func (service *creditService) CreditLine(credit entities.Credit) Response {
	log.Println("Service CreditLine ....")
	status := isRecommendedCreditLine(credit)
	return Response{
		Status: status.String(),
	}
}
