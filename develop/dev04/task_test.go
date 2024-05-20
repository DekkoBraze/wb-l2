package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	initArr := [8]string{"пятак", "листок", "пятка", "тяпка", "столик", "СЛИТОК", "лишний", "столик"}

	pSet := FindAnagrams(initArr[:])

	correct := map[string][]string{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}}

	if !reflect.DeepEqual(*pSet, correct) {
		t.Errorf("неверный вывод: получили %v нужно %v",
		*pSet, correct)
	}
}
