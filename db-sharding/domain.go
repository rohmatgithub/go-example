package dbsharding

import "time"

type ProductCategory struct {
	ID   int64  `gorm:"primaryKey,SERIAL"`
	Code string `gorm:"size:50;index:idx_code,unique"`
	Name string
}

func (ProductCategory) TableName() string {
	return "product_category"
}

type Product struct {
	ID          int64  `gorm:"primaryKey,SERIAL"`
	Title       string `gorm:"size:255"`
	Description string
	CategoryID  int64
	Price       float64
	Brand       string `gorm:"size:50"`
	SKU         string `gorm:"size:30"`
	Weight      float64
}

func (Product) TableName() string {
	return "product"
}

type Customer struct {
	ID      int64  `gorm:"primaryKey,SERIAL"`
	Name    string `gorm:"size:255"`
	Email   string `gorm:"size:100"`
	Phone   string `gorm:"size:20"`
	Address string
}

func (Customer) TableName() string {
	return "customer"
}

type SalesOrder struct {
	ID               int64     `gorm:"primaryKey,SERIAL"`
	SalesOrderNumber string    `gorm:"size:30;index:idx_invoice_number,unique"`
	SalesOrderDate   time.Time `gorm:"type:date"`
	CustomerID       int64
	TotalAmount      float64
	Status           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CreatedBy        string `gorm:"size:50"`
	UpdatedBy        string `gorm:"size:50"`
}

func (SalesOrder) TableName() string {
	return "sales_order"
}

type SalesOrderItem struct {
	ID           int64 `gorm:"primaryKey,SERIAL"`
	SalesOrderID int64
	ProductID    int64
	Quantity     int64
	SellingPrice float64
}

func (SalesOrderItem) TableName() string {
	return "sales_order_item"
}
