package helper

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func ApplyFilter(tx *gorm.DB, filters *map[string]string) error {
	for key, value := range *filters {
		keySplitted := strings.Split(key, ".")
		var column string
		if len(keySplitted) > 2 {
			column = strings.Join(keySplitted[:len(keySplitted)-1], ".")
		} else {
			column = keySplitted[0]
		}
		operator, err := OperatorQuery(keySplitted[len(keySplitted)-1])
		if err != nil {
			return err
		}
		query := fmt.Sprintf("%s %s ?", column, operator)
		switch operator {
		case "like":
			value = "%" + value + "%"
			tx = tx.Where(query, value)
		case "IN":
			valueSplitted := strings.Split(value, ",")
			tx = tx.Where(query, valueSplitted)
		default:
			tx = tx.Where(query, value)
		}
	}
	return nil
}
