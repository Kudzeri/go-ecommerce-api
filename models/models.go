package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      *string            `json:"first_name" bson:"first_name" validate:"required,min=3, max=20"`
	LastName       *string            `json:"last_name" bson:"last_name" validate:"required,min=3, max=30"`
	Email          *string            `json:"email" bson:"email" validate:"required,email"`
	Password       *string            `json:"password" bson:"password" validate:"required,min=8"`
	Phone          *string            `json:"phone" bson:"phone" validate:"required"`
	RefreshToken   *string            `json:"refresh_token" bson:"refresh_token"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
	UserID         string             `json:"user_id" bson:"user_id"`
	UserCart       []ProductUser      `json:"user_cart" bson:"user_cart"`
	AddressDetails []Address          `json:"address_details" bson:"address_details"`
	OrderStatus    []Order            `json:"order_status" bson:"order_status"x`
}

type Product struct {
	ProductID   primitive.ObjectID `json:"_id" bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name"`
	Price       *uint64            `json:"price" bson:"price"`
	Rating      *uint8             `json:"rating" bson:"rating"`
	ImageURL    *string            `json:"image_url" bson:"image_url"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `json:"_id" bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Rating      *uint              `json:"rating" bson:"rating"`
	ImageURL    *string            `json:"image_url" bson:"image_url"`
}

type Address struct {
	AddressID primitive.ObjectID `json:"_id" bson:"_id"`
	House     *string            `json:"house" bson:"house"`
	Street    *string            `json:"street" bson:"street"`
	City      *string            `json:"city" bson:"city"`
	Pincode   *string            `json:"pincode" bson:"pincode"`
}

type Order struct {
	OrderID       primitive.ObjectID `json:"_id" bson:"_id"`
	OrderCart     []ProductUser      `json:"order_cart" bson:"order_cart"`
	OrderedAt     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price         int                `json:"price" bson:"price"`
	Discount      *int               `json:"discount" bson:"discount"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}
type Payment struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod" bson:"cod"`
}
