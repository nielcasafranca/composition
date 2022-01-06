package invoice

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
}
