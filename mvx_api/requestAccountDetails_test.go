package mvx_api

import "testing"

func TestRequestAccountDetail(t *testing.T) {
	test := []struct {
		address string
		wantAcc Accounts
		wantErr bool
	}{
		{
			address: "",
			wantAcc: Accounts{},
			wantErr: true,
		},
		{
			address: "0x000000",
			wantAcc: Accounts{},
			wantErr: true,
		},
		{
			address: "erd1q54q99yklpg7cp8u4tku9g99c72x7ewtulflw3khsm2uvrgz5tkskpjp0z",
			wantAcc: Accounts{
				Address: "erd1q54q99yklpg7cp8u4tku9g99c72x7ewtulflw3khsm2uvrgz5tkskpjp0z",
			},
			wantErr: false,
		},
		{
			address: "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllst77y4l",
			wantAcc: Accounts{
				Address: "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllst77y4l",
			},
			wantErr: false,
		},
	}
	for _, tt := range test {
		acc, err := RequestAccountDetail(tt.address)
		if (err != nil) != tt.wantErr && acc.Address != tt.wantAcc.Address {
			t.Errorf("RequestAccountDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}
