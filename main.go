package main

import (
	"fmt"

	"github.com/nielcasafranca/composition/pkg/customer"
	"github.com/nielcasafranca/composition/pkg/invoice"
	"github.com/nielcasafranca/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Popaya",
		customer.New("Alejandro", "Cl 123 #23-4", "12345"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Curso de Go", 12.34),
			invoiceitem.New(2, "Curso de POO con Go", 54.23),
			invoiceitem.New(3, "Curso de testing con Go", 90.00),
		},
	)

	i.SetTotal()
	fmt.Printf("%+v", i)
}
