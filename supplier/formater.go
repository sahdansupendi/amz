package supplier

type SupplierFormatter struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	KodeSupplier string `json:"kode_supplier"`
	NoHp         string `json:"no_hp"`
	Email        string `json:"email"`
}

func FormatSupplier(supplier Supplier) SupplierFormatter {
	formatter := SupplierFormatter{
		ID:           supplier.ID,
		Name:         supplier.NamaSupplier,
		KodeSupplier: supplier.KodeSupplier,
		NoHp:         supplier.NoHp,
		Email:        supplier.Email,
	}

	return formatter
}
