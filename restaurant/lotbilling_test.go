package restaurant_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/lot-billing/restaurant"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/lot-billing/errors"
	"github.com/stretchr/testify/assert"
)

func TestLotBilling(t *testing.T) {
	t.Run("should return 3000 when input motorcycle with 1 hour", func(t *testing.T) {
		a, b := "motorcycle", 1
		expected := 3000

		result, _ := restaurant.LotBilling(a, b)

		assert.Equal(t, expected, result)
	})

	t.Run("should return 5000 when input motorcycle with 2 hour", func(t *testing.T) {
		a, b := "motorcycle", 2
		expected := 5000

		result, _ := restaurant.LotBilling(a, b)

		assert.Equal(t, expected, result)
	})

	t.Run("should return 177000 when input car with 25 hour", func(t *testing.T) {
		a, b := "car", 25
		expected := 177000

		result, _ := restaurant.LotBilling(a, b)

		assert.Equal(t, expected, result)
	})

	t.Run("should return 'vehicle is not available' when input unknown vehicle", func(t *testing.T) {
		a, b := "bus", 25
		expected := errors.ErrVehicleIsNotAvailable

		_, err := restaurant.LotBilling(a, b)

		assert.EqualError(t, expected, err.Error())
	})
}
