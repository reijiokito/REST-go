package restaurantmodel

import (
	"errors"
	"food-delivery/common"
	"strings"
)

type RestaurantType string

const (
	TypeNormal  RestaurantType = "normal"
	TypePremium RestaurantType = "premium"
)

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" grom:"column:name;"`
	Addr            string         `json:"addr" grom:"column:addr;"`
	Status          int            `json:"status" gorm:"column:status;"`
	Type            RestaurantType `json:"type" gorm:"column:type"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Id              int    `json:"id" grom:"column:id;"`
	Name            string `json:"name" grom:"column:name;"`
	Addr            string `json:"addr" grom:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

type RestaurantUpdate struct {
	Name string `json:"name" grom:"column:name;"`
	Addr string `json:"addr" grom:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name cannot be empty")
)
