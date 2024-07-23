package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName      *string            `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName       *string            `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email          *string            `json:"email,omitempty" bson:"email,omitempty"`
	Password       *string            `json:"password,omitempty" bson:"password,omitempty"`
	Phone          *string            `json:"phone,omitempty" bson:"phone,omitempty"`
	RefreshToken   *string            `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	CreatedAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UserID         string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	UserCart       []ProductUser      `json:"user_cart,omitempty" bson:"user_cart,omitempty"`
	AddressDetails []Address          `json:"address_details,omitempty" bson:"address_details,omitempty"`
	OrderStatus    []Order            `json:"order_status,omitempty" bson:"order_status,omitempty"x`
}

type Product struct {
	ProductID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductName *string            `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Price       *uint64            `json:"price,omitempty" bson:"price,omitempty"`
	Rating      *uint8             `json:"rating,omitempty" bson:"rating,omitempty"`
	ImageURL    *string            `json:"image_url,omitempty" bson:"image_url,omitempty"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductName *string            `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Price       int                `json:"price,omitempty" bson:"price,omitempty"`
	Rating      *uint              `json:"rating,omitempty" bson:"rating,omitempty"`
	ImageURL    *string            `json:"image_url,omitempty" bson:"image_url,omitempty"`
}

type Address struct {
	AddressID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	House     *string            `json:"house,omitempty" bson:"house,omitempty"`
	Street    *string            `json:"street,omitempty" bson:"street,omitempty"`
	City      *string            `json:"city,omitempty" bson:"city,omitempty"`
	Pincode   *string            `json:"pincode,omitempty" bson:"pincode,omitempty"`
}

type Order struct {
	OrderID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrderCart     []ProductUser      `json:"order_cart,omitempty" bson:"order_cart,omitempty"`
	OrderedAt     time.Time          `json:"ordered_at,omitempty" bson:"ordered_at,omitempty"`
	Price         int                `json:"price,omitempty" bson:"price,omitempty"`
	Discount      *int               `json:"discount,omitempty" bson:"discount,omitempty"`
	PaymentMethod Payment            `json:"payment_method,omitempty" bson:"payment_method,omitempty"`
}
type Payment struct {
	Digital bool `json:"digital,omitempty" bson:"digital,omitempty"`
	COD     bool `json:"cod,omitempty" bson:"cod,omitempty"`
}
