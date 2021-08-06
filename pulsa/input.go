package pulsa

type FindProductInput struct {
	Brand    string `json:"brand" binding:"required"`
	Category string `json:"category" binding:"required"`
}

type PulsaTransactionInput struct {
	ProductName         string `json:"product_name" binding:"required"`
	Category            string `json:"category" binding:"required"`
	Brand               string `json:"brand" binding:"required"`
	Type                string `json:"type" binding:"required"`
	SelerName           string `json:"seller_name" binding:"required"`
	SkuCode             string `json:"skucode" binding:"required"`
	BuyerProductStatus  bool   `json:"buyer_product_status" binding:"required"`
	SellerProductStatus bool   `json:"seller_product_status" binding:"required"`
	Desc                string `json:"desc" binding:"required"`
	SupplierPrice       int    `json:"supplier_price" binding:"required"`
	Price               int    `json:"price" binding:"required"`
}
