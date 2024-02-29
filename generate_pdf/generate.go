package generate_pdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"subscription_management_system/email"
	"subscription_management_system/models"
	"subscription_management_system/utility"
)

func GeneratePDF(contract models.Contract, filePath string) error {
	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

	durationUnitText := utility.GetDurationUnitText(contract.DurationUnit)
	contractStatus := utility.GetContractStatus(contract.Status)
	fmt.Println(durationUnitText, " ", contractStatus, " ", contract.Customer.Name)

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, "Contract Details")

	pdf.SetFont("Arial", "", 12)
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("ID: %d", contract.ID))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Customer: %s", contract.Customer.Name))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Subscription: %s", contract.Subscription.Name))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Billing Interval: %s", durationUnitText))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Installment Amount: %.2f", contract.InstallmentAmount))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Duration: %d", contract.Duration))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Duration Unit: %s", durationUnitText))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Status: %s", contractStatus))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Contract Start Date: %s", contract.ContractStartDate.UTC()))

	// Output the PDF to a file
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return err
	}
	to := contract.Customer.Email
	subject := "Remaining Installment Bill"
	body := "<h1>Hello!</h1><p>Please pay remaining installment bill.</p>"
	err = email.SendEmailWithAttachment(to, subject, body, filePath)
	if err != nil {
		fmt.Println("Mail sending unsuccessful. Reason: ", err)
	}

	return nil
}
