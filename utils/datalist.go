package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Paginate struct {
	Data       any          `json:"data"`
	Pagination PaginateType `json:"pagination"`
}
type PaginateType struct {
	Limit    int `json:"limit"`
	Offset   int `json:"offset"`
	Total    int `json:"total"`
	Filtered int `json:"filtered"`
}

func Pagination[T any](data []T, c *fiber.Ctx) Paginate {
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)
	orderBy := c.Query("order_by", "")
	orderDir := c.Query("order_dir", "asc")
	q := c.Queries()

	filtered, fnum := FilterByQuery(data, q)

	if orderBy != "" {
		fmt.Printf("orderby: %v, orderdir: %v", orderBy, orderDir)
	}

	totalFiltered := 0
	if fnum > 0 {
		totalFiltered = len(filtered)
	}

	paginate := PaginateType{
		Limit:    limit,
		Offset:   offset,
		Total:    len(data),
		Filtered: totalFiltered,
	}
	end := offset + limit
	if end > len(filtered) {
		end = len(filtered)
	}
	if offset > len(filtered) {
		offset = len(filtered)
	}

	res := Paginate{
		Pagination: paginate,
		Data:       filtered[offset:end],
	}
	return res
}

func FilterByQuery[T any](data []T, q map[string]string) ([]T, int) {
	var filtered []T
	fnum := 0

	for idx, item := range data {
		v := reflect.ValueOf(item)
		if v.Kind() == reflect.Pointer {
			v = v.Elem()
		}
		match := true
		for k, val := range q {
			if k == "limit" || k == "offset" || k == "sortby" || k == "sortdir" {
				if idx == 0 {
					fnum++
				}
				continue
			}

			field := v.FieldByNameFunc(func(name string) bool {
				return strings.EqualFold(name, k) // case-insensitive match
			})

			if !field.IsValid() {
				continue
			}

			// Convert field value to string for comparison
			fieldVal := fmt.Sprint(field.Interface())
			if fieldVal != val {
				match = false
				break
			}
		}

		if match {
			filtered = append(filtered, item)
		}
	}

	return filtered, len(q) - fnum
}
