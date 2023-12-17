package oasis

import (
	"fmt"
	"strings"
)

type historyHash string
type PredictionRecords map[historyHash]int
type Report struct {
	Histories             []History
	calculatedPredictions PredictionRecords
}
type History struct {
	Values []int
}

func computePrediction(r PredictionRecords, values []int) int {
	hHash := historyHash(strings.ReplaceAll(fmt.Sprint(values), " ", ","))
	pred, ok := r[hHash]
	if ok {
		return pred
	}
	nextValues := make([]int, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		nextValues[i] = values[i+1] - values[i]
	}
	allZeros := true
	for _, v := range nextValues {
		if v != 0 {
			allZeros = false
		}
	}
	firstValue := values[0]
	var prediction int
	if allZeros {
		prediction = firstValue
	} else {
		prediction = firstValue - computePrediction(r, nextValues)
	}
	r[hHash] = prediction
	return prediction
}

func (r Report) ComputePredictions() []int {
	if r.calculatedPredictions == nil {
		r.calculatedPredictions = make(PredictionRecords)
	}
	var preds []int
	for _, v := range r.Histories {
		preds = append(preds, computePrediction(r.calculatedPredictions, v.Values))
	}
	return preds
}
