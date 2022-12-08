package application

type ProductInterface interface {
	isValid() (bool, error)
	Enable() error
	Disable() error
	getID() string
	getName() string
	getStatus() string
	getPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

// func (p *Product) isValid() (bool, error) {

// }

// func (p *Product) Enable() error {

// }

// func (p *Product) Disable() error {

// }

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
