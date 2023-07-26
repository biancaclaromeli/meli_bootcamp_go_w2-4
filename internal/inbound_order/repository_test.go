package inboundorder_test

import (
	"context"
	"testing"
	"time"

	"github.com/mercadolibre/fury_bootcamp-go-w2-s4-4-4/internal/domain"
	inboundorder "github.com/mercadolibre/fury_bootcamp-go-w2-s4-4-4/internal/inbound_order"
	"github.com/mercadolibre/fury_bootcamp-go-w2-s4-4-4/pkg/testutil"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Creates valid inbound order", func(t *testing.T) {
		db := testutil.InitDatabase(t)
		defer db.Close()

		repo := inboundorder.NewRepository(db)

		order := domain.InboundOrder{
			OrderDate:      time.Now().AddDate(0, 0, 1),
			OrderNumber:    "123456",
			EmployeeID:     1,
			ProductBatchID: 2,
			WarehouseID:    1,
		}

		id, _ := repo.Save(context.TODO(), order)
		var receivedNumber string

		row := db.QueryRow(`SELECT order_number FROM inbound_orders WHERE id = ?;`, id)
		row.Scan(&receivedNumber)

		assert.Equal(t, order.OrderNumber, receivedNumber)
	})
	t.Run("Does not create invalid inbound order", func(t *testing.T) {
		db := testutil.InitDatabase(t)
		defer db.Close()

		repo := inboundorder.NewRepository(db)

		order := domain.InboundOrder{
			OrderDate:      time.Now().AddDate(0, 0, 1),
			OrderNumber:    "123456",
			EmployeeID:     1,
			ProductBatchID: 2,
			WarehouseID:    1,
		}

		_, err := repo.Save(context.TODO(), order)
		assert.NoError(t, err)

		_, err = repo.Save(context.TODO(), order)
		assert.Error(t, err)
	})
}
