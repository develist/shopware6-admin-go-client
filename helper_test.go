package shopware6_admin_go_client

import "testing"

func TestToKebabCase(t *testing.T) {
	wanted := "ab-cd-ef-gh-ij"
	result := toKebabCase("AbCdEfGhIj")
	if result != wanted {
		t.Fatalf(`Wanted was %v but return value is %v`, wanted, result)
	}
}
