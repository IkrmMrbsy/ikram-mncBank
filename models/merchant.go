package models

type Merchant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Account string `json:"account"` // Nomor akun merchant
}
