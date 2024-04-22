package ports

// import (
// 	"context"
// )

type GRPCPort interface {
	RUN()
	GetAddition()
	GetSubtraction()
	GetMultiplication()
	GetDivision()
}
