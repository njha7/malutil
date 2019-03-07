package malutil

import (
	"os"
	"testing"

	"gotest.tools/assert"
)

var (
	testUsers = []string{
		"/profile/Skwizz",
		"/profile/Ahmedlzs",
		"/profile/VHMD",
		"/profile/aPPPLL",
		"/profile/KittyLilly",
		"/profile/Marfizawa",
		"/profile/riot-",
		"/profile/nguyenquanghuy2",
		"/profile/Sorekun",
		"/profile/Saku1",
		"/profile/luiscool100",
		"/profile/TitaniumTriforce",
		"/profile/SharpEdge",
		"/profile/NuclearDragon",
		"/profile/syllomex",
		"/profile/Illiyia",
		"/profile/Sandshrew2",
		"/profile/attorrep",
		"/profile/destroyer068",
		"/profile/Eternal-Destiny",
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
