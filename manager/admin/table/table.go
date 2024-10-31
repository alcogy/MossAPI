package table

type Table struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetAllTables() []Table {
	return []Table{
		{
			Name:        "customer",
			Description: "Management customer.",
		},
		{
			Name:        "product",
			Description: "Management product.",
		},
		{
			Name:        "customer_product",
			Description: "Relation customer and product.",
		},
	}
}