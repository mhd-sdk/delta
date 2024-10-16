package orders

import "context"

type OrdersService interface {
	Init(ctx context.Context) error
}
