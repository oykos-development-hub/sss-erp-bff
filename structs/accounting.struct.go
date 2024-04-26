package structs

import "time"

type AccountingOrderForObligationsData struct {
	InvoiceID          []int     `json:"invoice_id"`
	SalaryID           []int     `json:"salary_id"`
	DateOfBooking      time.Time `json:"date_of_booking"`
	OrganizationUnitID int       `json:"organization_unit_id"`
}
