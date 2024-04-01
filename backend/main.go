package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Make an HTTP POST request in Go
// https://www.kirandev.com/http-post-golang

// After you're done with some basic HTTP requests, set up a front-end for the myDATA

// XML
// https://gobyexample.com/xml

type InvoicesDoc struct {
	Invoice Invoice `json: "invoice"`
}

type Invoice struct {
	Ιssuer         Ιssuer         `json: "issuer"`
	Counterpart    Counterpart    `json: "counterpart"`
	InvoiceHeader  InvoiceHeader  `json: "invoiceHeader"`
	PaymentMethods PaymentMethods `json: "paymentMethods"`
	InvoiceDetails InvoiceDetails `json: "invoiceDetails"`
	InvoiceSummary InvoiceSummary `json: "invoiceSummary"`
}

type Ιssuer struct {
	VatNumber string `json: "vatNumber"`
	Country   string `json: "country"`
	Branch    string `json: "branch"`
}

type Counterpart struct {
	VatNumber string  `json: "vatNumber"`
	Country   string  `json: "country"`
	Branch    string  `json: "branch"`
	Address   Address `json: "address"`
}

type Address struct {
	City       string `json: "city"`
	PostalCode string `json: "postalCode"`
}

type InvoiceHeader struct {
	Series      string `json: "series"`
	AA          string `json: "aa"`
	IssueDate   string `json: "issueDate"`
	InvoiceType string `json: "invoiceType"`
	Currency    string `json: "postalCode"`
}

type PaymentMethods struct {
	PaymentMethodDetails PaymentMethodDetails `json: "PaymentMethodDetails"`
}

type PaymentMethodDetails struct {
	Type   string  `json: "type"`
	Amount float64 `json: "amount"`
}

type InvoiceDetails struct {
	LineNumber           string               `json: "lineNumber"`
	NetValue             float64              `json: "netValue"`
	VatCategory          string               `json: "vatCategory"`
	VatAmount            float64              `json: "vatAmount"`
	IncomeClassification IncomeClassification `json: "incomeClassification"`
}

type IncomeClassification struct {
	ClassificationType     string  `json: "N1:classificationType"`
	ClassificationCategory string  `json: "N1:classificationCategory"`
	Amount                 float64 `json: "N1:amount"`
}

type InvoiceSummary struct {
	TotalNetValue         float64              `json: "totalNetValue"`
	TotalVatAmount        float64              `json: "totalVatAmount"`
	TotalWithheldAmount   float64              `json: "totalWithheldAmount"`
	TotalFeesAmount       float64              `json: "totalFeesAmount"`
	TotalStampDutyAmount  float64              `json: "totalStampDutyAmount"`
	TotalOtherTaxesAmount float64              `json: "totalOtherTaxesAmount"`
	TotalDeductionsAmount float64              `json: "totalDeductionsAmount"`
	TotalGrossValue       float64              `json: "totalGrossValue"`
	IncomeClassification  IncomeClassification `json: "incomeClassification"`
}

// project goals (do at least 2/4 to upload it)
// IMPERATIVE: SET (TEST/PRODUCTION MODE, aka the respective links per mode)
// delete invoices
// add invoice
// list invoices
// choice to save the data in an .xml file

// MYDATA test account
// USER: mydata_test_user
// OCP_SUBSCRIPTION_KEY: ccc23ee3881d600a3814d4ae441dd160

// test urls - mydata rest apis
// https://mydataapidev.aade.gr/SendInvoices
// https://mydataapidev.aade.gr/CancelInvoice
// https://mydataapidev.aade.gr/RequestDocs (έχουν υποβάλλει άλλοι για εσένα)
// https://mydataapidev.aade.gr/RequestTransmittedDocs (έχεις υποβάλλει ο ίδιος για σένα)

func godotEnvironmentVariable(value string, print bool) string {

	// TEST CODE
	// envVals := [...]string{"HOST", "USER_ID", "OCP_SUBSCRIPTION_KEY"}
	// envs := make([]string, 0)
	// for i := 0; i < 3; i++ {
	// 	envs = append(envs, godotEnvironmentVariable(envVals[i], false))
	// }

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	if print {
		fmt.Print(" [" + os.Getenv(value) + "] ")
	}
	return os.Getenv(value)
}

func runInvoiceFunction(functionArgument string) {

	fmt.Print("Function: " + functionArgument)

}

func SendInvoices() {
	// TODO

	//add invoice data as args
	//&invoiceDoc
	//marshall it
	//make it a json at first
	//try read data from it
	//make it an xml
	//double-check everything works
	//send it over

	// test_XML_data, err := os.Open("test_xml_file.xml")
	// if err != err {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Created invoice handle.")

	// request, _ := http.NewRequest(http.MethodPost, "https://mydataapidev.aade.gr/SendInvoices", bytes.NewBuffer(test_XML_data))
	// response, _ := http.DefaultClient.Do(request)
	// body, _ := io.ReadAll(response.Body)
	// defer response.Body.Close()

	// fmt.Println(string(body))
}

func CancelInvoice() {
	// TODO

	// request, _ := http.NewRequest("GET", "https://api.sampleapis.com/beers/ale", nil)
	// response, _ := http.DefaultClient.Do(request)

	// body, _ := io.ReadAll(response.Body)
	// defer response.Body.Close()
	// fmt.Println(string(body))
}

func RequestDocs() {
	// TODO
}

func RequestTransmittedDocs() {
	// TODO
}

func main() {
	fmt.Println("myDATA solution!")

	// SET DATA PASS-THROUGH VIA (PREFERABLY) WEBSOCKET FROM FRONTEND

	args := os.Args
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: No argument specified.\n")
		os.Exit(1)
	}

	runInvoiceFunction(args[1])
	// SendInvoices()
}
