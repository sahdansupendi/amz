package supplier

type RegisterSupllierInput struct {
	NamaSupplier string `json:"nama_supplier" binding:"required"`
	Alamat       string `json:"alamat" binding:"required"`
	NoHp         string `json:"no_hp" binding:"required"`
	Email        string `json:"email" binding:"required"`
}
