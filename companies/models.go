package companies

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name                string              `json:"name" gorm:"column:name"`
	Web                 string              `json:"web" gorm:"column:web"`
	Linkedin            string              `json:"linkedin" gorm:"column:linkedin"`
	Started             int                 `json:"started" gorm:"column:started"`
	MeasurableAssets    MeasurableAssets    `json:"measurable_assets"`
	MidMeasurableAssets MidMeasurableAssets `json:"mid_measurable_assets"`
	Pitch               string              `json:"pitch" gorm:"column:pitch;size:15000"`
	MainHits            string              `json:"main_hits" gorm:"column:main_hits;size:15000"`
	Services            string              `json:"services" gorm:"column:services;size:15000"`
	ServiceBenefits     string              `json:"service_benefits" gorm:"column:service_benefits;size:15000"`
	Target              string              `json:"target" gorm:"column:target;size:15000"`
	Clients             []Client            `json:"clients"`
	Investors           []Investor          `json:"investors"`
	Competitors         []Competitor        `json:"competitors"`
}

type Client struct {
	gorm.Model
	Name        string `json:"name"`
	Sector      string `json:"sector"`
	ClientType  int    `json:"client_type"` //1 to 5
	IsForbes500 bool   `json:"is_forbes_500"`
	CompanyID   uint   `json:"company_id"`
}

type Investor struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name"`
	CompanyID uint   `json:"company_id"`
}

type Competitor struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name"`
	CompanyID uint   `json:"company_id"`
}

type MeasurableAssets struct {
	gorm.Model
	AnnualIncome      float64 `json:"annual_income" gorm:"column:annual_income"`
	AnnualOutcome     float64 `json:"annual_outcome" gorm:"column:annual_outcome"`
	Profit            float64 `json:"profit" gorm:"column:profit"`
	Taxes             float64 `json:"taxes" gorm:"column:taxes"` //percentage
	ProfitMargin      float64 `json:"profit_margin" gorm:"column:profit_margin"`
	AvgGrowthMinFive  float64 `json:"avg_growth_min_five" gorm:"column:avg_growth_min_five"`
	AvgGrowthPlusFive float64 `json:"avg_growth_plus_five" gorm:"column:avg_growth_plus_five"`
	Employees         int     `json:"employees" gorm:"column:employees"`
	CompanyID         uint    `json:"company_id"`
}

type Stage string

type MidMeasurableAssets struct {
	gorm.Model
	ReturnMargin  string `json:"return_margin" gorm:"column:return_margin"` // i.e "5-15"
	MinInvestment string `json:"min_investment" gorm:"column:min_investment"`
	Liquidity     string `json:"liquidity" gorm:"column:liquidity"`
	Stage         Stage  `json:"stage" gorm:"column:stage"`
	CompanyID     uint   `json:"company_id"`
}

type SoftAssets struct {
	gorm.Model
	Sector        string `json:"sector" gorm:"column:sector"`
	OfficeCountry string `json:"office_country" gorm:"column:office_country"`
}

type Valuation struct {
	gorm.Model
	Value       float64 `json:"value"`
	CompanyName string  `json:"company_name"`
}

func (c *Company) CalculateValuation() *Valuation {
	return &Valuation{
		Value:       165526,
		CompanyName: c.Name,
	}
}

func crateClients(s []string) []Client {
	var clients []Client
	for _, name := range s {
		clients = append(clients, Client{
			Name: name,
		})
	}
	return clients
}

func createInvestors(s []string) []Investor {
	var investors []Investor
	for _, name := range s {
		investors = append(investors, Investor{
			Name: name,
		})
	}
	return investors
}

func crateCompetitors(s []string) []Competitor {
	var competitors []Competitor
	for _, name := range s {
		competitors = append(competitors, Competitor{
			Name: name,
		})
	}
	return competitors
}
