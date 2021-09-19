package companies

import "gorm.io/gorm"

type Storage struct {
	DB *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

func (s *Storage) Create(c *Company) (uint, error) {
	err := s.DB.Create(c).Error
	if err != nil {
		return 0, err
	}

	return c.ID, nil
}

func (s *Storage) Get(id uint) (*Company, error) {
	var dest Company
	err := s.DB.Model(&Company{}).First(&dest, "id = ?", id).Error

	return &dest, err
}

func (s *Storage) GetRanking() ([]Company, error) {
	var dest []Company
	err := s.DB.First(&dest).Order("").Limit(20).Error
	return dest, err
}

// CalculateValue calculate and save the valuation generated for the given company
func (s *Storage) CalculateValue(id uint) (*Valuation, error) {
	company, err := s.Get(id)

	if err != nil {
		return nil, err
	}

	valuation := company.CalculateValuation()

	//save the valuation
	s.DB.Create(valuation)

	return valuation, nil
}
