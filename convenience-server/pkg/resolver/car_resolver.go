package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/Convenience-Tools/convenience-server/pkg/middleware"
)

type Car struct {
	Name      string `json:"name"`
	Price     int32  `json:"price"`
	Create_at string `json:"created_at"`
}

type AddCarArgs struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

func (*RootResolver) AddCar(ctx context.Context, args AddCarArgs) (Car, error) {
	user, _ := ctx.Value(middleware.UserCtxKey).(*middleware.User)
	fmt.Println(user)

	newCar := Car{
		Name:      args.Name,
		Price:     args.Price,
		Create_at: time.Now().String(),
	}

	return newCar, nil
}

func (*RootResolver) Cars(ctx context.Context) ([]Car, error) {
	user, _ := ctx.Value(middleware.UserCtxKey).(*middleware.User)
	fmt.Println(user)

	cars := []Car{
		{
			Name:      "Tesla Model 3",
			Price:     110,
			Create_at: time.Now().String(),
		},
		{
			Name:      "Tesla Y",
			Price:     90,
			Create_at: time.Now().String(),
		},
	}

	return cars, nil
}
