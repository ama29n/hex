package api

import (
	"github.com/ama29n/hex/internal/ports"
)

type Adapter struct {
	arith ports.ArithmaticPort
	db    ports.DBPort
}

// dependency injection
// this is part of application layer
// benefit of this is, we can make changes to core layer without needing to change application layer

func NewAdapter(db ports.DBPort, arith ports.ArithmaticPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}

// application layer can access both core and framewrok layer
// but since dependencies points inwards, we have to give application layer access to db through dependency injection
// and application layer is given access to core through dependecy injection, however it can be imported directly

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	res, error := apia.arith.Addition(a, b)

	if error != nil {
		return 0, error
	}

	// add result to DB
	error = apia.db.AddToHistory(res, "addition")
	if error != nil {
		return 0, error
	}

	// return result
	return res, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	res, error := apia.arith.Subtraction(a, b)
	if error != nil {
		return 0, error
	}

	// add result to DB
	error = apia.db.AddToHistory(res, "subtraction")
	if error != nil {
		return 0, error
	}

	return res, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	res, error := apia.arith.Multiplication(a, b)
	if error != nil {
		return 0, error
	}

	// add result to DB
	error = apia.db.AddToHistory(res, "multiplication")
	if error != nil {
		return 0, error
	}

	return res, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	res, error := apia.arith.Division(a, b)
	if error != nil {
		return 0, error
	}

	// add result to DB
	error = apia.db.AddToHistory(res, "division")
	if error != nil {
		return 0, error
	}

	return res, nil
}
