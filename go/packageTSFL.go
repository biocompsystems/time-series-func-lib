package TSFL

import (
	"errors"
	"fmt"
)

type Functioner interface {
	Compute1D(arrayIn []float64, period int) ([]float64, error)
	Compute1Dx(arrayIn []float64, period int, startRow int, endRow int) ([]float64, error)
	Compute2D(arrayIn [][]float64, period int, columnID int) ([]float64, error)
	Compute2Dx(arrayIn [][]float64, period int, startRow int, endRow int, columnID int) ([]float64, error)
}

type SMA struct {
	RequiredColumns int
	MinimumRows     int
}

func MakeSMA() *SMA {
	sma := new(SMA)
	sma.MinimumRows = 1
	sma.RequiredColumns = 1
	return sma
}
func doParameterChecks(data []float64, period int, startRow int, endRow int) error {
	//Do checks
	if data == nil {
		return errors.New("data cannot be nil")
	}
	if period <= 0 {
		return errors.New("period must be greater than zero")
	}
	if startRow < 0 {
		return errors.New("startRow cannot be negative")
	}
	if endRow < 0 {
		return errors.New("endRow cannot be negative")
	}
	if endRow < startRow {
		return errors.New("endRow must be greater or equal to startRow")
	}
	return nil
}
func doParameterChecks2D(data [][]float64, period int, startRow int, endRow int, columnID int) error {
	//Do checks
	if data == nil {
		return errors.New("data cannot be nil")
	}
	if period <= 0 {
		return errors.New("period must be greater than zero")
	}
	if startRow < 0 {
		return errors.New("startRow cannot be negative")
	}
	if endRow < 0 {
		return errors.New("endRow cannot be negative")
	}
	if endRow < startRow {
		return errors.New("endRow must be greater or equal to startRow")
	}
	if columnID <= 0 {
		return errors.New("columnID must be greater than zero")
	}
	if columnID > len(data) {
		return errors.New("columnID is beyond the bounds of the array")
	}
	return nil
}
func (f *SMA) Compute1D(data []float64, period int) ([]float64, error) {
	//merely convert to extended form, call and return the result
	return f.Compute1Dx(data, period, 0, len(data))
}
func (f *SMA) Compute1Dx(data []float64, period int, startRow int, endRow int) ([]float64, error) {

	err := doParameterChecks(data, period, startRow, endRow)
	if err != nil {
		return nil, err
	}
	//allocate returned result the same as the incoming array
	returnData := make([]float64, len(data))
	//if the period is greater than the length of the data, just return array with zeros
	if period > len(data) {
		return returnData, nil
	}
	//Figure out where to start computing.  It isn't just as simple as using start and end row
	//if the start row is less than the period, use the period, else the start row
	startPoint := startRow
	if startRow < period {
		startPoint = period
	}

	//if the endrow is bigger than or equal to the data length, use the data length
	//else use the endrow + 1 as we are going one less than the endpoint: i < endPoint
	var endPoint int
	if endRow >= len(data) {
		endPoint = len(data)
	} else {
		endPoint = endRow + 1
	}

	//Walk through the data from start to end
	for i := startPoint - 1; i < endPoint; i++ {
		//computing a moving window of size period
		value := 0.0
		for j := i; j > i-period; j += -1 {
			value += data[j]
		}
		returnData[i] = value / float64(period)
	}
	return returnData, nil
}
func (f *SMA) Compute2D(data [][]float64, period int, columnID int) ([]float64, error) {
	//merely convert to extended form, call and return the result
	fmt.Printf("Data len: %d\r\n", len(data[0]))
	return f.Compute2Dx(data, period, 0, len(data[0]), columnID)
}
func (f *SMA) Compute2Dx(data [][]float64, period int, startRow int, endRow int, columnID int) ([]float64, error) {
	err := doParameterChecks2D(data, period, startRow, endRow, columnID)
	if err != nil {
		return nil, err
	}
	//allocate returned result the same as the incoming array
	returnData := make([]float64, len(data[0]))
	//if the period is greater than the length of the data, just return array with zeros
	if period > len(data[0]) {
		return returnData, errors.New("Period is longer than data.")
	}
	//Figure out where to start computing.  It isn't just as simple as using start and end row
	//if the start row is less than the period, use the period, else the start row
	startPoint := startRow
	if startRow < period {
		startPoint = period
	}

	//if the endrow is bigger than or equal to the data length, use the data length
	//else use the endrow + 1 as we are going one less than the endpoint: i < endPoint
	var endPoint int
	if endRow >= len(data[0]) {
		endPoint = len(data[0])
	} else {
		endPoint = endRow + 1
	}

	//Walk through the data from start to end
	for i := startPoint - 1; i < endPoint; i++ {
		//computing a moving window of size period
		value := 0.0
		for j := i; j > i-period; j += -1 {
			value += data[columnID][j]
		}
		returnData[i] = value / float64(period)
	}
	return returnData, nil
}
