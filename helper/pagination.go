package helper

// Pagination struct
type Pagination struct {
	Page  uint `form:"p"`
	Limit uint `form:"l"`
}

// GetPage to get current page
func (p Pagination) GetPage() uint {
	if p.Page == 0 {
		return 1
	}
	if p.Page > 100 {
		return 100
	}
	return p.Page
}

// GetOffset to get current offset
func (p Pagination) GetOffset() uint {
	page := p.GetPage()
	limit := p.GetLimit()
	offset := (page - 1) * limit
	return offset
}

// GetLimit to get limit
func (p *Pagination) GetLimit() uint {
	if p.Limit == 0 || p.Limit > 5 {
		return 5
	}
	return p.Limit
}
