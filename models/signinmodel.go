package models

import (
	"github.com/jinzhu/gorm"
)

// type FiriModel struct {
// 	ID        int        `gorm:"primary_key" json:"id"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"updated_at"`
// 	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
// }

type RegistrationData struct {
	gorm.Model
	FullName string `json:"fullname"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// type UserAuthentication struct {
// 	// FiriModel
// 	FullName string `gorm:"type:varchar(100)" json:"full_name"`
// 	Email    string `gorm:"type:varchar(100); unique" json:"email"`
// 	Password string `gorm:"type:varchar(100)" json:"password"`
// }
//LoginData takes user login information
type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//RegistrationMailObject handles registration mail object
// type RegistrationMailObject struct {
// 	Email            string `json:"email"`
// 	VerificationCode string `json:"code"`
// 	FullName         string `json:"full_name"`
// }

//ResponseObject holds response object data
type ResponseObject struct {
	Code    int         `json:"code"`
	Body    interface{} `json:"body"`
	Message string      `json:"message"`
}
