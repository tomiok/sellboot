package companies

type CompanyGateway interface {
	Create(c *Company) (uint, error)
	Get(id uint) (*Company, error)
	GetRanking() ([]Company, error)
	CalculateValue(id uint) (*Valuation, error)
}

type CompanyService struct {
	Gateway CompanyGateway
}

func NewGateway(g CompanyGateway) *CompanyService {
	return &CompanyService{Gateway: g}
}

func (svc *CompanyService) CreateCompany(c *Company) error {
	_, err := svc.Gateway.Create(c)
	return err
}

func (svc *CompanyService) GetCompany(id uint) (*Company, error) {
	return svc.Gateway.Get(id)
}

func (svc *CompanyService) GetRanking() ([]Company, error) {
	return svc.Gateway.GetRanking()
}

func (svc *CompanyService) GetValuation(id uint) (*Valuation, error) {
	return svc.Gateway.CalculateValue(id)
}
