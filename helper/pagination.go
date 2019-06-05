package helper

// PaginationPageMax defines max value for page
const PaginationPageMax = 100

// PaginationLimitMax defines max value for limit
const PaginationLimitMax = 100

// Pagination struct
type Pagination struct {
	Page  uint `form:"p"`
	Limit uint `form:"l"`
}

// GetPage returns current page
func (p Pagination) GetPage() uint {
	if p.Page == 0 {
		return 1
	}
	if p.Page > PaginationPageMax {
		return PaginationPageMax
	}
	return p.Page
}

// GetOffset returns current offset
func (p Pagination) GetOffset() uint {
	page := p.GetPage()
	limit := p.GetLimit()
	offset := (page - 1) * limit
	return offset
}

// GetLimit returns limit of pagination
func (p *Pagination) GetLimit() uint {
	if p.Limit == 0 || p.Limit > PaginationLimitMax {
		return PaginationLimitMax
	}
	return p.Limit
}
