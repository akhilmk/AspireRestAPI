package model

type User struct {
	UserID   int    `json:"userid,omitempty" gorm:"primarykey;not null"`
	UserName string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Age      int    `json:"age,omitempty"`
	Mobile   int    `json:"mobile" gorm:"-"`
}
