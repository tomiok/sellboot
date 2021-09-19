package companies

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

type Web struct {
	Svc *CompanyService
}

// CompanyFormHandler mvc for the form template
func (w *Web) CompanyFormHandler(ctx *fiber.Ctx) error {
	return ctx.Render("templates/form", nil)
}

// SaveCompanyHandler create a new company
func (w *Web) SaveCompanyHandler(ctx *fiber.Ctx) error {
	var dto CompanyDTO
	err := ctx.BodyParser(&dto)

	if err != nil {
		return err
	}

	company := dto.toCompany()
	ctx.Status(http.StatusCreated)
	return w.Svc.CreateCompany(company)
}

// GetCompaniesRankingHandler get the ranking by value
func (w *Web) GetCompaniesRankingHandler(ctx *fiber.Ctx) error {
	companies, err := w.Svc.GetRanking()

	if err != nil {
		return err
	}

	return ctx.JSON(companies)
}

func (w *Web) GetValuationHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	u, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		log.Error().Msgf("invalid id %s", id)
		return err
	}

	valuation, err := w.Svc.GetValuation(uint(u))

	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "cannot get valuation")
	}

	return ctx.JSON(valuation)
}

type CompanyDTO struct {
	Name              string  `json:"name"`
	Web               string  `json:"web"`
	Linkedin          string  `json:"linkedin"`
	Started           int     `json:"started"`
	AnnualIncome      float64 `json:"annual_income"`
	AnnualOutcome     float64 `json:"annual_outcome"`
	Profit            float64 `json:"profit"`
	Taxes             float64 `json:"taxes"`
	ProfitMargin      float64 `json:"profit_margin"`
	AvgGrowthMinFive  float64 `json:"avg_growth_min_five"`
	AvgGrowthPlusFive float64 `json:"avg_growth_plus_five"`
	Employees         int     `json:"employees"`
	ReturnMargin      string  `json:"return_margin"`
	MinInvestment     string  `json:"min_investment"`
	Liquidity         string  `json:"liquidity"`
	Stage             string  `json:"stage"`
	Pitch             string  `json:"pitch"`
	MainHits          string  `json:"main_hits"`
	Services          string  `json:"services"`
	ServiceBenefits   string  `json:"service_benefits"`
	Target            string  `json:"target"`
	Clients           string  `json:"clients"`
	Investors         string  `json:"investors"`
	Competitors       string  `json:"competitors"`
}

func (dto *CompanyDTO) toCompany() *Company {
	clients := crateClients(strings.Split(dto.Clients, ","))
	investors := createInvestors(strings.Split(dto.Investors, ","))
	competitors := crateCompetitors(strings.Split(dto.Competitors, ","))
	return &Company{
		Name:     dto.Name,
		Web:      dto.Web,
		Linkedin: dto.Linkedin,
		Started:  dto.Started,
		MeasurableAssets: MeasurableAssets{
			AnnualIncome:      dto.AnnualIncome,
			AnnualOutcome:     dto.AnnualOutcome,
			Profit:            dto.Profit,
			Taxes:             dto.Taxes,
			ProfitMargin:      dto.ProfitMargin,
			AvgGrowthMinFive:  dto.AvgGrowthMinFive,
			AvgGrowthPlusFive: dto.AvgGrowthPlusFive,
			Employees:         dto.Employees,
		},
		MidMeasurableAssets: MidMeasurableAssets{
			ReturnMargin:  dto.ReturnMargin,
			MinInvestment: dto.MinInvestment,
			Liquidity:     dto.Liquidity,
			Stage:         Stage(dto.Stage),
		},
		Pitch:           dto.Pitch,
		MainHits:        dto.MainHits,
		Services:        dto.Services,
		ServiceBenefits: dto.ServiceBenefits,
		Target:          dto.Target,
		Clients:         clients,
		Investors:       investors,
		Competitors:     competitors,
	}
}
