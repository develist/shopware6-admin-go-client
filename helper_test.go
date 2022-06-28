package shopware6_admin_go_client

import (
	"testing"
)

func TestToKebabCase(t *testing.T) {
	wanted := "ab-cde-fghi-jkl-mn"
	result := toKebabCase("AbCdeFghiJklMn")
	if result != wanted {
		t.Fatalf(`Wanted was %v but return value is %v`, wanted, result)
	}
}
