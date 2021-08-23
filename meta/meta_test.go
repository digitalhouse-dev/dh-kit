package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		tag                                                                    string
		page, perPage, total                                                   int
		expectedPage, expectedPerPage, expectedTotal, pageCount, offset, limit int
	}{
		// varying page
		{"t1", 1, 20, 50, 1, 20, 50, 3, 0, 20},
		{"t2", 2, 20, 50, 2, 20, 50, 3, 20, 20},
		{"t3", 3, 20, 50, 3, 20, 50, 3, 40, 20},
		{"t4", 4, 20, 50, 3, 20, 50, 3, 40, 20},
		{"t5", 0, 20, 50, 1, 20, 50, 3, 0, 20},

		// varying perPage
		{"t6", 1, 0, 50, 1, 25, 50, 2, 0, 25},
		{"t7", 1, -1, 50, 1, 25, 50, 2, 0, 25},
		{"t8", 1, 100, 50, 1, 100, 50, 1, 0, 100},
		{"t9", 1, 1001, 50, 1, 1000, 50, 1, 0, 1000},

		// varying total
		{"t10", 1, 20, 0, 1, 20, 0, 0, 0, 20},
		{"t11", 1, 20, -1, 1, 20, -1, -1, 0, 20},
	}

	for _, test := range tests {
		p := New(test.page, test.perPage, test.total)
		assert.Equal(t, test.expectedPage, p.Page, test.tag)
		assert.Equal(t, test.expectedPerPage, p.PerPage, test.tag)
		assert.Equal(t, test.expectedTotal, p.TotalCount, test.tag)
		assert.Equal(t, test.pageCount, p.PageCount, test.tag)
		assert.Equal(t, test.offset, p.Offset(), test.tag)
		assert.Equal(t, test.limit, p.Limit(), test.tag)
	}
}

func Test_parseInt(t *testing.T) {
	type args struct {
		value        string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"123", 100}, 123},
		{"t2", args{"", 100}, 100},
		{"t3", args{"a", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInt(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("parseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFormMapt(t *testing.T) {
	req := make(map[string]string)
	req["page"] = "2"
	req["limit"] = "20"
	p := NewFormMap(req, 100)
	assert.Equal(t, 2, p.Page)
	assert.Equal(t, 20, p.PerPage)
	assert.Equal(t, 100, p.TotalCount)
	assert.Equal(t, 5, p.PageCount)
}
