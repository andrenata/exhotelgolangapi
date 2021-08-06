package pulsa

type PulsaTransaction struct {
	ID            int
	RefId         int
	CustomerNo    string
	SkuCode       string
	Message       string
	Status        string
	Rc            string
	Sn            string
	SupplierPrice int
	SellerPrice   int
	Tele          string
	Wa            string
}
