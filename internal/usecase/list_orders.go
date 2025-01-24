package usecase

import (
	"fmt"

	"github.com/maxnet04/CleanArch/internal/entity"
)

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute() ([]ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return nil, fmt.Errorf(
			"error listing orders: %w",
			err)
	}
	dto := make([]ListOrderOutputDTO, len(orders))
	for i := range orders {
		result := ListOrderOutputDTO{
			ID:         orders[i].ID,
			Price:      orders[i].Price,
			Tax:        orders[i].Tax,
			FinalPrice: orders[i].FinalPrice,
		}
		dto[i] = result
	}

	return dto, nil
}
