package dbmodel

type User struct {
	UserID   int `gorm:"primarykey;not null"`
	UserName int
	Email    int
	Gender   string
	Age      int
	Mobile   int
}
