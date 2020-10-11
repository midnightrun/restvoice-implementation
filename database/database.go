package database

import "github.com/midnightrun/restvoice-implementation/domain"

type Repository struct {
	invoices map[int]domain.Invoice
}

func NewRepository() Repository {
	return Repository{invoices: make(map[int]domain.Invoice)}
}

func (r *Repository) CreateInvoice(invoice domain.Invoice) domain.Invoice {
	id := r.nextId()
	invoice.Id = id
	invoice.Status = "open"
	invoice.Bookings = []domain.Booking{}
	r.invoices[invoice.Id] = invoice
	return invoice
}

func (r Repository) nextId() int {
	id := 0

	for k := range r.invoices {
		if k >= id {
			id = k + 1
		}
	}

	return id
}
