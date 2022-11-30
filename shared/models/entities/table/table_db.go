package table

type Categories struct {
	Id   string `gorm:"not null;unique"`
	Name string `gorm:"not null;uniqueIndex;primaryKey"`
}

type User struct {
	UserId       string `gorm:"size:128;not null;uniqueIndex;primary_key" json:"user_id"`
	UserEmail    string `gorm:"size:128;not null;uniqueIndex" json:"user_email"`
	UserName     string `gorm:"size:128;not null" json:"user_name"`
	UserImage    string `gorm:"size:256;not null" json:"user_image"`
	UserPassword string `gorm:"size:128;not null" json:"user_password"`
	UserSession  string `gorm:"size:256;null" json:"user_session"`
	UserStatus   StatusUser
	UserStatusId uint32    `gorm:"not null;" json:"user_status"`
	CreatedDate  string    `gorm:"size:128;null" json:"created_date"`
	LastUpdate   string    `gorm:"size:128;null" json:"last_update"`
	RememberMe   bool      `gorm:"size:8;null" json:"remember_me"`
	Products     []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StatusUser struct {
	Id     int8   `gorm:"not null;uniqueIndex;primaryKey"`
	Status string `gorm:"not null;unique"`
}

type Product struct {
	UserId       string `gorm:"not null;" json:"user_id"`
	Category     Categories
	CategoryName string `gorm:"not null" json:"category_name"`
	ProductId    string `gorm:"size:128;not null;uniqueIndex;primary_key" json:"product_id"`
	ProductName  string `gorm:"size:128;not null" json:"product_name"`
	ProductImage string `gorm:"size:256;not null;default='kosong'" json:"product_image"`
	ProductInfo  string `gorm:"size:2024;not null" json:"product_info"`
	ProductStock uint32 `gorm:"not null" json:"product_stoct"`
	ProductPrice uint32 `gorm:"not null" json:"product_price"`
	ProductSell  uint32 `gorm:"not null" json:"product_sell"`
	CreatedDate  string `gorm:"size:128;null" json:"created_date"`
	LastUpdate   string `gorm:"size:128;null" json:"last_update"`
}
