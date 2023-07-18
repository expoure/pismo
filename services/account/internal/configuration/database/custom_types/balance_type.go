package custom_types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Balance struct {
	Amount   int64
	Currency string
}

func (a *Balance) Scan(value interface{}) error {
	byteData, ok := value.([]byte)
	if !ok {
		fmt.Println("Asseertion error")
		return errors.New("Assertion error")
	}

	// Remove "()" from composite type
	dataString := strings.Trim(string(byteData), "()")
	values := strings.Split(dataString, ",")

	amount, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return errors.New(err.Error())
	}

	currency := strings.Trim(values[1], "'")

	var res Balance
	res = Balance{
		Amount:   amount,
		Currency: currency,
	}

	*a = res
	return nil
}