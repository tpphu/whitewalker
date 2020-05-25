package helper

import (
	"testing"
)

func TestPagination_GetPage(t *testing.T) {
	type fields struct {
		Page  uint
		Limit uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			"Get Page at 0",
			fields{0, 10},
			1,
		},
		{
			"Get Page with valid data",
			fields{10, 10},
			10,
		},
		{
			"Get Page with valid data",
			fields{100, 10},
			100,
		},
		{
			"Get Page with invalid data",
			fields{101, 10},
			PaginationPageMax,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pagination{
				Page:  tt.fields.Page,
				Limit: tt.fields.Limit,
			}
			if got := p.GetPage(); got != tt.want {
				t.Errorf("Pagination.GetPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_GetLimit(t *testing.T) {
	type fields struct {
		Page  uint
		Limit uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			"Get Limit with valid data",
			fields{10, 10},
			10,
		},
		{
			"Get Limit with valid data",
			fields{10, 100},
			100,
		},
		{
			"Get Limit with invalid data",
			fields{10, 101},
			PaginationLimitMax,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				Page:  tt.fields.Page,
				Limit: tt.fields.Limit,
			}
			if got := p.GetLimit(); got != tt.want {
				t.Errorf("Pagination.GetLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_GetOffset(t *testing.T) {
	type fields struct {
		Page  uint
		Limit uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			"Get Offset with valid data",
			fields{10, 10},
			90,
		},
		{
			"Get Offset with valid data",
			fields{10, 100},
			900,
		},
		{
			"Get Offset with invalid data",
			fields{10, 101},
			(10 - 1) * PaginationLimitMax,
		},
		{
			"Get Offset with invalid data",
			fields{101, 10},
			(PaginationLimitMax - 1) * 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pagination{
				Page:  tt.fields.Page,
				Limit: tt.fields.Limit,
			}
			if got := p.GetOffset(); got != tt.want {
				t.Errorf("Pagination.GetOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
