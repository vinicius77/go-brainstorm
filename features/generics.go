package features

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type DataGeneric string

/* tilde ~ accepts custom types of the same type => DataGeneric of type string */
type User[T ~string | int64] struct {
	Name string
	Age  int
	Date T
}

func addGeneric[T constraints.Ordered](a T, b T) T {
	return a + b
}

func MapValues[T constraints.Float | constraints.Integer](values []T, multiply func(T) T) {
	var newValues []T

	for _, value := range values {
		newValue := multiply(value)
		newValues = append(newValues, newValue)
	}

	fmt.Println(newValues)

}

func Multiply[T constraints.Integer | constraints.Float](item T) T {
	return item * 2
}

func FeaturesCheck() {

	user := User[int64]{Name: "Vincius", Age: 99, Date: 1720625718964}
	user2 := User[DataGeneric]{Name: "Second", Age: 77, Date: "13-12-2009"}

	values := []int{1, 4, 6, 7}
	valuesFloat := []float32{12.2, 32.6, 34.2}

	MapValues(values, Multiply[int])
	MapValues(valuesFloat, Multiply[float32])

	res := addGeneric(1, 4)
	floatRes := addGeneric(3.5, 8.9)
	stringRes := addGeneric("Vinicius", " hello")

	fmt.Println("Feature Name", res, floatRes, user, user2, stringRes)

}
