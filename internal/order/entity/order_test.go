package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T){
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")	
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T){
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")	
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T){
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")	
}

func TestGivenAaValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order:= Order{ID: "123", Price: 10.0, Tax: 2.0}

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t ,order.IsValid())
}

func TestGivenAaValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)

	assert.Nil(t, err)

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t ,order.IsValid())
}

func TestGivenAPriceAndTax_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	order.CalculatePrice()

	assert.Nil(t, order.CalculatePrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}