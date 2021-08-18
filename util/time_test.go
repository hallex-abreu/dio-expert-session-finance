package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2021-02-12T10:10:10")

	if convertedTime.Year() != 2021 {
		t.Errorf("Espera que o ano seja 2021")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Espera que o mês seja 02")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("Espera que a hora seja 10")
	}
}
