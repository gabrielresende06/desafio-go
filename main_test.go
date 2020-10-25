package main

import "testing"

func TestSum(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Função soma está incorreta, o resultado deve ser: %d, e o valor retornado foi: %d", 10, total);
	}
}