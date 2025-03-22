package validate

import (
	"net/http"
	"testing"
)

// TODO: 注册的翻译没生效，不知道为什么
func TestCNMobilePhoneNumber(t *testing.T) {
	type foo struct {
		PhoneNumber string `zh_Hans_CN:"手机号" validate:"cnmobilephonenumber"`
	}

	tests := []struct {
		name    string
		foo     foo
		wantErr bool
	}{
		{
			name:    "valid_Mobile",
			foo:     foo{PhoneNumber: "17828852173"},
			wantErr: false,
		},
		{
			name:    "FIXED_LINE",
			foo:     foo{PhoneNumber: "01082716601"},
			wantErr: true,
		},
		{
			name:    "not_Mobile",
			foo:     foo{PhoneNumber: "178288521739"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validator{}.Validate((*http.Request)(nil), tt.foo)
			if (err != nil) != tt.wantErr {
				t.Errorf("CNMobilePhoneNumber() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if err != nil {
				t.Log(err)
			}
		})
	}
}
