/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package fixture

// DefaultUser
func DefaultUser() map[string]string {
	// gofakeit.Seed(time.Now().UnixNano())

	return map[string]string{
		"name":     "example",
		"email":    "example@test.com",
		"password": "12345678",
	}
}
