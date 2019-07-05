package creational

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue = 3
)

type ShirtsCache struct {}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whiteProtoType
		return &newItem, nil
	case Black:
		newItem := *blackProtoType
		return &newItem, nil
	case Blue:
		newItem := *blueProtoType
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

var whiteProtoType *Shirt = &Shirt{
	Price: 15.00,
	SKU: "empty",
	Color: White,
}

var blackProtoType *Shirt = &Shirt{
	Price: 15.00,
	SKU: "empty",
	Color: Black,
}

var blueProtoType *Shirt = &Shirt{
	Price: 15.00,
	SKU: "empty",
	Color: Blue,
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

