package malutil

import (
	"os"
	"testing"

	"gotest.tools/assert"
)

var (
	testUsers = []string{
		"Skwizz",
		"Ahmedlzs",
		"VHMD",
		"aPPPLL",
		"KittyLilly",
		"Marfizawa",
		"riot-",
		"nguyenquanghuy2",
		"Sorekun",
		"Saku1",
		"luiscool100",
		"TitaniumTriforce",
		"SharpEdge",
		"NuclearDragon",
		"syllomex",
		"Illiyia",
		"Sandshrew2",
		"attorrep",
		"destroyer068",
		"Eternal-Destiny",
	}
)

func TestGetUsers(t *testing.T) {
	page, err := os.Open("test_pages/users.html")
	if err != nil {
		t.Errorf("Could not load test page.")
	}
	users := GetUsersFromPage(page)
	// fmt.Println(users)
	// fmt.Println(testUsers)
	assert.DeepEqual(t, users, testUsers)
}
