package jwt_test

import (
	"encoding/json"
	"fmt"

	"github.com/cristalhq/jwt"
)

func Example_Validate() {
	t := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZG1pbiIsImp0aSI6InJhbmRvbS11bmlxdWUtc3RyaW5nIn0.dv9-XpY9P8ypm1uWQwB6eKvq3jeyodLA7brhjsf4JVs`

	token, err := jwt.Parse([]byte(t))
	if err != nil {
		fmt.Printf("parse err: %q", err)
		return
	}

	fmt.Printf("Algorithm %v\n", token.Header().Algorithm)
	fmt.Printf("Type      %v\n", token.Header().Type)
	fmt.Printf("Claims    %v\n", string(token.RawClaims()))
	fmt.Printf("Payload   %v\n", string(token.Payload()))
	fmt.Printf("Token     %v\n", string(token.Raw()))

	claims := &jwt.StandardClaims{}
	_ = json.Unmarshal(token.RawClaims(), claims)

	validator := jwt.NewValidator(
		jwt.AudienceChecker([]string{"admin"}),
		jwt.IDChecker("random-unique-string"),
	)

	err = validator.Validate(claims)
	if err != nil {
		fmt.Printf("token is invalid: %#v", err)
	}

	// Output:
	// Algorithm HS256
	// Type      JWT
	// Claims    {"aud":"admin","jti":"random-unique-string"}
	// Payload   eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZG1pbiIsImp0aSI6InJhbmRvbS11bmlxdWUtc3RyaW5nIn0
	// Token     eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZG1pbiIsImp0aSI6InJhbmRvbS11bmlxdWUtc3RyaW5nIn0.dv9-XpY9P8ypm1uWQwB6eKvq3jeyodLA7brhjsf4JVs
}
