package constants

var Perms map[string]byte = map[string]byte{
	"PROMOTE_USER":               0,
	"DELETE_USER":                1,
	"GET_USERS":                  2,
	"GET_USER_EMAIL":             3,
	"GET_USER_PHONE":             4,
	"GET_USER_ADDRESS":           5,
	"ADD_ROLE":                   6,
	"UPDATE_ROLE_NAME":           7,
	"UPDATE_ROLE_DESCRIPTION":    8,
	"UPDATE_ROLE_PERMS":          9,
	"DELETE_ROLE":                10,
	"GET_ROLE":                   11,
	"ADD_PRODUCT":                12,
	"GET_PRODUCT":                13,
	"UPDATE_PRODUCT_NAME":        14,
	"UPDATE_PRODUCT_DESCRIPTION": 15,
	"UPDAE_PRODUCT_EXPLANATION":  16,
	"UPDATE_PRODUCT_SPECS":       17,
	"UPDATE_PRODUCT_CATEGORY":    18,
	"DELETE_PRODUCT":             19,
}
