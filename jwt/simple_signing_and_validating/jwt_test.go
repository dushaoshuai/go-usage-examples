package simple_signing_and_validating

import (
	"testing"
	"time"

	"github.com/goccy/go-json"
)

func jsonMarshal(v any) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

func Test_jwt(t *testing.T) {
	token, err := jwtNew()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("token: %s", token)

	ticker := time.Tick(time.Second)
	for range ticker {
		claims, err := parse(token)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("claims: %+v", jsonMarshal(claims))
	}

	// Output:
	// $ go test -run Test_jwt
	// --- FAIL: Test_jwt (5.00s)
	//    jwt_test.go:20: token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnby11c2FnZS1leGFtcGxlcyIsInN1YiI6ImdvLXVzYWdlLWV4YW1wbGVzLXRlc3RpbmciLCJhdWQiOlsiZ28tdXNhZ2UtZXhhbXBsZXMiXSwiZXhwIjoxNzQyMTA3NjMyLCJuYmYiOjE3NDIxMDc2MjcsImlhdCI6MTc0MjEwNzYyNywianRpIjoidW5pcXVlX2p0aSIsInVzZXJfaWQiOjF9.ZswToK26Qeb-eqAng163Ilau8AN9dSlYL2fB28aPWiM
	//    jwt_test.go:28: claims: {"iss":"go-usage-examples","sub":"go-usage-examples-testing","aud":["go-usage-examples"],"exp":1742107632,"nbf":1742107627,"iat":1742107627,"jti":"unique_jti","user_id":1}
	//    jwt_test.go:28: claims: {"iss":"go-usage-examples","sub":"go-usage-examples-testing","aud":["go-usage-examples"],"exp":1742107632,"nbf":1742107627,"iat":1742107627,"jti":"unique_jti","user_id":1}
	//    jwt_test.go:28: claims: {"iss":"go-usage-examples","sub":"go-usage-examples-testing","aud":["go-usage-examples"],"exp":1742107632,"nbf":1742107627,"iat":1742107627,"jti":"unique_jti","user_id":1}
	//    jwt_test.go:28: claims: {"iss":"go-usage-examples","sub":"go-usage-examples-testing","aud":["go-usage-examples"],"exp":1742107632,"nbf":1742107627,"iat":1742107627,"jti":"unique_jti","user_id":1}
	//    jwt_test.go:26: token has invalid claims: token is expired
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/jwt/simple_signing_and_validating      5.005s
}
