package enums

import "errors"

type CategoryType int

const (
	Electronics CategoryType = 1
	Fashion     CategoryType = 2
	Beauty      CategoryType = 3
	Sports      CategoryType = 4
	Automotive  CategoryType = 5
	Health      CategoryType = 6
	Beverages   CategoryType = 7
	Books       CategoryType = 8
	Toys        CategoryType = 9
	Home        CategoryType = 10
)

func (c CategoryType) GetCategoryName() (string, error) {
	switch c {
	case Electronics:
		return "Electronics", nil
	case Fashion:
		return "Fashion", nil
	case Beauty:
		return "Beauty", nil
	case Sports:
		return "Sports", nil
	case Automotive:
		return "Automotive", nil
	case Health:
		return "Health", nil
	case Beverages:
		return "Beverages", nil
	case Books:
		return "Books", nil
	case Toys:
		return "Toys", nil
	case Home:
		return "Home", nil
	default:
		return "", errors.New("no category type found")
	}

}

