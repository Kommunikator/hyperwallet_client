package hyperwallet

import (
	"context"
	"reflect"
	"testing"
)

func NewTestPaypalAccountGateway() *PaypalAccountGateway {
	return &PaypalAccountGateway{
		NewTestClient(),
	}
}

func TestCreatePaypalAccountDataValidate(t *testing.T) {
	t.Parallel()

	type test struct {
		input    CreatePaypalAccountData
		expected string
	}

	tests := []test{
		{
			CreatePaypalAccountData{
				TransferMethodCountry:  "US",
				TransferMethodCurrency: "USD",
				Type:                   "PAYPAL_ACCOUNT",
				Email:                  "qwe@ad.comm",
			},
			"",
		},
		{
			CreatePaypalAccountData{
				TransferMethodCountry:  "US",
				TransferMethodCurrency: "USD",
				Type:                   "WIRE_ACCOUNT",
				Email:                  "asd@ad.comm",
			},
			"Bad value for Type",
		},
		{
			CreatePaypalAccountData{
				TransferMethodCurrency: "USD",
				Type:                   "WIRE_ACCOUNT",
			},
			"Key: 'CreatePaypalAccountData.TransferMethodCountry' Error:Field validation for 'TransferMethodCountry' failed on the 'required' tag\n" +
				"Key: 'CreatePaypalAccountData.Email' Error:Field validation for 'Email' failed on the 'required' tag",
		},
	}

	for _, tc := range tests {
		err := tc.input.Validate()

		// Позитивный сценарий
		if err == nil {
			if tc.expected == "" {
				continue
			} else {
				t.Errorf("Wrong data not validated, expected - %s, reciving - nil", tc.expected)
				continue
			}
		}

		// Негативный сценарий
		if err.Error() != tc.expected {
			t.Errorf("%s", err.Error())
		}
	}
}

func TestUpdatePaypalAccountDataValidate(t *testing.T) {
	t.Parallel()

	type test struct {
		input    UpdatePaypalAccountData
		expected string
	}

	tests := []test{
		{
			UpdatePaypalAccountData{
				Email: "qwe@ad.comm",
			},
			"",
		},
		{
			UpdatePaypalAccountData{
				Email: "@ad.comm",
			},
			"Key: 'UpdatePaypalAccountData.Email' Error:Field validation for 'Email' failed on the 'email' tag",
		},
		{
			UpdatePaypalAccountData{},
			"Key: 'UpdatePaypalAccountData.Email' Error:Field validation for 'Email' failed on the 'required' tag",
		},
	}

	for _, tc := range tests {
		err := tc.input.Validate()

		// Позитивный сценарий
		if err == nil {
			if tc.expected == "" {
				continue
			} else {
				t.Errorf("Wrong data not validated, expected - %s, reciving - nil", tc.expected)
				continue
			}
		}

		// Негативный сценарий
		if err.Error() != tc.expected {
			t.Errorf("%s", err.Error())
		}
	}
}

func TestGetPaypalAccountList(t *testing.T) {
	testClient := NewTestPaypalAccountGateway()

	ctx := context.Background()

	paypalAccountList, err := testClient.GetPaypalAccountList(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", GetPaypalAccountListQuery{})
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	expected := &PaypalAccountList{
		Count:  1,
		Offset: 0,
		Limit:  10,
		Data: []PaypalAccount{
			{
				Token:                  "trm-0d66f04d-4340-4820-87a3-721a5e4a2754",
				Type:                   "PAYPAL_ACCOUNT",
				Status:                 "ACTIVATED",
				CreatedOn:              "2021-10-21T14:07:23",
				TransferMethodCountry:  "US",
				TransferMethodCurrency: "USD",
				UserToken:              "usr-c9d3126d-e26d-459d-9d66-9538876848be",
				Email:                  "94105@asde.com",
				Links: []Link{
					{
						Params: Params{Rel: "self"},
						Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/paypal-accounts/trm-0d66f04d-4340-4820-87a3-721a5e4a2754",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(paypalAccountList, expected) {
		t.Errorf("Unexpected result of GetPaypalAccountList func")
	}
}

func TestCreatePaypalAccount(t *testing.T) {
	testClient := NewTestPaypalAccountGateway()

	ctx := context.Background()

	createPaypalAccountData := CreatePaypalAccountData{
		TransferMethodCountry:  "US",
		TransferMethodCurrency: "USD",
		Type:                   "PAYPAL_ACCOUNT",
		Email:                  "94105@asde.com",
	}

	expected := &PaypalAccount{
		Token:                  "trm-0d66f04d-4340-4820-87a3-721a5e4a2754",
		Type:                   "PAYPAL_ACCOUNT",
		Status:                 "ACTIVATED",
		CreatedOn:              "2021-10-21T14:07:23",
		TransferMethodCountry:  "US",
		TransferMethodCurrency: "USD",
		UserToken:              "usr-c9d3126d-e26d-459d-9d66-9538876848be",
		Email:                  "94105@asde.com",
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/paypal-accounts/trm-0d66f04d-4340-4820-87a3-721a5e4a2754",
			},
		},
	}

	paypalAccount, err := testClient.CreatePaypalAccount(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", createPaypalAccountData)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(paypalAccount, expected) {
		t.Errorf("Unexpected result of CreatePaypalAccount func")
	}
}

func TestUpdatePaypalAccount(t *testing.T) {
	testClient := NewTestPaypalAccountGateway()

	ctx := context.Background()

	updatePaypalAccountData := UpdatePaypalAccountData{
		Email: "tst5@gmali.comm",
	}

	expected := &PaypalAccount{
		Token:                  "trm-0d66f04d-4340-4820-87a3-721a5e4a2754",
		Type:                   "PAYPAL_ACCOUNT",
		Status:                 "ACTIVATED",
		CreatedOn:              "2021-10-21T14:07:23",
		TransferMethodCountry:  "US",
		TransferMethodCurrency: "USD",
		UserToken:              "usr-c9d3126d-e26d-459d-9d66-9538876848be",
		Email:                  "tst5@gmali.comm",
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/paypal-accounts/trm-0d66f04d-4340-4820-87a3-721a5e4a2754",
			},
		},
	}

	paypalAccount, err := testClient.UpdatePaypalAccount(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", "trm-0d66f04d-4340-4820-87a3-721a5e4a2754", updatePaypalAccountData)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(paypalAccount, expected) {
		t.Errorf("Unexpected result of UpdatePaypalAccount func")
	}
}


func TestRetrievePaypalAccount(t *testing.T) {
	testClient := NewTestPaypalAccountGateway()

	ctx := context.Background()

	expected := &PaypalAccount{
		Token:                  "trm-0d66f04d-4340-4820-87a3-721a5e4a2754",
		Type:                   "PAYPAL_ACCOUNT",
		Status:                 "ACTIVATED",
		CreatedOn:              "2021-10-21T14:07:23",
		TransferMethodCountry:  "US",
		TransferMethodCurrency: "USD",
		UserToken:              "usr-c9d3126d-e26d-459d-9d66-9538876848be",
		Email:                  "94105@asde.com",
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/paypal-accounts/trm-0d66f04d-4340-4820-87a3-721a5e4a2754",
			},
		},
	}

	paypalAccount, err := testClient.RetrievePaypalAccount(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", "trm-0d66f04d-4340-4820-87a3-721a5e4a2754")
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(paypalAccount, expected) {
		t.Errorf("Unexpected result of RetrievePaypalAccount func")
	}
}
