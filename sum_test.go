package main

import "testing"

func TestSumFivetoFive(t *testing.T) {
	result := SumFiveToFive();

	if result != 10 {
		t.Errorf("Soma esperada: %d, obtida: %d", 10, result);
	}
}