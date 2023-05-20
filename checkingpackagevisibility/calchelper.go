package checkingpackagevisibility

// PublicGetDiscount this function can be used outside by importing checkingpackagevisibility package
func PublicGetDiscount(price float64) float64 {
	discount := 0.95
	return price * discount
}

// this function cannot be used outside by importing checkingpackagevisibility package
func privateGetDiscount(price float64) float64 {
	return price * .50
}
