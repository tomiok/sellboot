package companies

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	datastorage "sellboot/storage"
	"strings"
	"testing"
)

var web = Web{Svc: NewGateway(NewStorage(datastorage.GetTestDB(Entities()...)))}
var app = fiber.New()
var companyDTO = CompanyDTO{
	Name:              "LinkedIn",
	Web:               "linkedin.com",
	Linkedin:          "linkedin.com/linkedin",
	Started:           2002,
	AnnualIncome:      10000000,
	AnnualOutcome:     80000,
	Profit:            250000,
	Taxes:             182365,
	ProfitMargin:      123,
	AvgGrowthMinFive:  12,
	AvgGrowthPlusFive: 26,
	Employees:         5000,
	ReturnMargin:      "",
	MinInvestment:     "500000",
	Liquidity:         "extreme",
	Stage:             "fortune5000",
	Pitch:             "",
	MainHits:          "a lot",
	Services:          "social media",
	ServiceBenefits:   "",
	Target:            "all the people",
	Clients:           "everyone",
	Investors:         "Chase, JP Morgan, Bank of Manhattan",
	Competitors:       "jobmas,zonajobs",
}
func TestCreateCompanyHandler(t *testing.T) {
	app.Post("/companies", web.SaveCompanyHandler)

	b, _ := json.Marshal(&companyDTO)
	req := httptest.NewRequest("POST", "/companies", strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, 200)


	if err != nil {
		t.Error(err.Error())
	}
	//TODO finish this and add a json return to save company
	fmt.Println(res.Body)
}