package shared

type Product struct {
	Product_type_id string `json:"product_type_id"`
	Pid             string `json:"pid"`
	Poster          string `json:"poster"`
	Name            string `json:"name"`
	SKU             string `json:"sku"`
	Price           int    `json:"price"`
	Quantity        int    `json:"quantity"`
	Brand           string `json:"brand"`
	Category        string `json:"category"`
	Color           string `json:"color"`
	Material        string `json:"material"`
	Weight          string `json:"weight"`
	Size            string `json:"size"`
	OriginalPrice   string `json:"originalPrice"`
	Sale            bool   `json:"sale"`
	Discount        int    `json:"discount"`
	IsProduct       bool   `json:"isproduct"`
}

type Category struct {
	Product_type_id string `json:"product_type_id"`
	Poster          string `json:"poster"`
	IsProduct       bool   `json:"isproduct"`
}

type CartProducts struct {
	Name     string `json:"name"`
	Pid      string `json:"pid"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Poster   string `json:"poster"`
}
