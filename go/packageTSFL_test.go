package TSFL

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func MakeSampleDataf64(rows int, cols int, random bool) [][]float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	col := make([]float64, rows)
	for i := 0; i < rows; i++ {
		if random == true {
			col[i] = rand.Float64()
		} else {
			col[i] = float64(i)
		}
		//fmt.Printf("%f\r\n", col[i])
	}
	dat := make([][]float64, cols)
	for i := 0; i < cols; i++ {
		dat[i] = col
	}
	return dat
}

//objective: To assure that sum sums
func TestSMA1D(t *testing.T) {
	data := make([]float64, 5)
	for i := 0; i < 5; i++ {
		data[i] = float64(i)
	}
	period := 3
	sma := MakeSMA()
	result, err := sma.Compute1D(data, period)
	if err != nil {
		fmt.Println("Oops")
	}
	for i := 0; i < len(result); i++ {
		fmt.Printf("%f\r\n", result[i])
	}
	if result[4] != 3.0 {
		fmt.Printf("1D FAIL: MA in last row is: %f\n", result[4])
	} else {
		fmt.Println("1D PASS: MA correct.")
	}

}

//objective: To assure that sum sums
func TestSMA2D(t *testing.T) {
	data := MakeSampleDataf64(5, 2, false)
	period := 3
	columnID := 1
	sma := MakeSMA()
	result, err := sma.Compute2D(data, period, columnID)
	if err != nil {
		fmt.Println("Oops")
	}
	for i := 0; i < len(result); i++ {
		fmt.Printf("%f\r\n", result[i])
	}
	if result[4] != 3.0 {
		fmt.Printf("2D FAIL: MA in last row is: %f\n", result[4])
	} else {
		fmt.Println("2D PASS: MA correct.")
	}

}
