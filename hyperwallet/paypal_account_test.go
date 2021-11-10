package hyperwallet

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

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
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

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

	const userToken = "usr-c9d3126d-e26d-459d-9d66-9538876848be"

	httpmock.RegisterRegexpResponder(
		"GET",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/users/"+userToken+"/paypal-accounts"),
		httpmock.NewBytesResponder(200,
			[]byte("{\"count\":1,\"offset\":0,\"limit\":10,\"data\":[{\"token\":\"trm-0d66f04d-4340-4820-87a3-721a5e4a2754\",\"type\":\"PAYPAL_ACCOUNT\",\"status\":\"ACTIVATED\",\"createdOn\":\"2021-10-21T14:07:23\",\"transferMethodCountry\":\"US\",\"transferMethodCurrency\":\"USD\",\"userToken\":\"usr-c9d3126d-e26d-459d-9d66-9538876848be\",\"email\":\"94105@asde.com\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/paypal-accounts/trm-0d66f04d-4340-4820-87a3-721a5e4a2754\"}]}]}"),
		),
	)

	pg := PaypalAccountGateway{testClient}

	paypalAccountList, err := pg.GetPaypalAccountList(ctx, userToken, GetPaypalAccountListQuery{})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, paypalAccountList)
	}
}

func TestCreatePaypalAccount(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

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

	const userToken = "usr-c9d3126d-e26d-459d-9d66-9538876848be"

	httpmock.RegisterRegexpResponder(
		"POST",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/users/"+userToken+"/paypal-accounts"),
		httpmock.NewBytesResponder(200,
			[]byte("{\"token\":\"trm-0d66f04d-4340-4820-87a3-721a5e4a2754\",\"type\":\"PAYPAL_ACCOUNT\",\"status\":\"ACTIVATED\",\"createdOn\":\"2021-10-21T14:07:23\",\"transferMethodCountry\":\"US\",\"transferMethodCurrency\":\"USD\",\"userToken\":\"usr-c9d3126d-e26d-459d-9d66-9538876848be\",\"email\":\"94105@asde.com\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/paypal-accounts/trm-0d66f04d-4340-4820-87a3-721a5e4a2754\"}]}"),
		),
	)

	pg := PaypalAccountGateway{testClient}

	paypalAccount, err := pg.CreatePaypalAccount(ctx, userToken, createPaypalAccountData)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, paypalAccount)
	}
}

func TestUpdatePaypalAccount(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

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

	const userToken = "usr-c9d3126d-e26d-459d-9d66-9538876848be"
	const paypalAccountToken = "trm-0d66f04d-4340-4820-87a3-721a5e4a2754"

	httpmock.RegisterRegexpResponder(
		"PUT",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/users/"+userToken+"/paypal-accounts/"+paypalAccountToken),
		httpmock.NewBytesResponder(200,
			[]byte("{\"token\":\"trm-0d66f04d-4340-4820-87a3-721a5e4a2754\",\"type\":\"PAYPAL_ACCOUNT\",\"status\":\"ACTIVATED\",\"createdOn\":\"2021-10-21T14:07:23\",\"transferMethodCountry\":\"US\",\"transferMethodCurrency\":\"USD\",\"userToken\":\"usr-c9d3126d-e26d-459d-9d66-9538876848be\",\"email\":\"tst5@gmali.comm\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/paypal-accounts/trm-0d66f04d-4340-4820-87a3-721a5e4a2754\"}]}"),
		),
	)

	pg := PaypalAccountGateway{testClient}

	paypalAccount, err := pg.UpdatePaypalAccount(ctx, userToken, paypalAccountToken, updatePaypalAccountData)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, paypalAccount)
	}
}

func TestRetrievePaypalAccount(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

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

	const userToken = "usr-c9d3126d-e26d-459d-9d66-9538876848be"
	const paypalAccountToken = "trm-0d66f04d-4340-4820-87a3-721a5e4a2754"

	httpmock.RegisterRegexpResponder(
		"GET",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/users/"+userToken+"/paypal-accounts/"+paypalAccountToken),
		httpmock.NewBytesResponder(200,
			[]byte("{\"token\":\"trm-0d66f04d-4340-4820-87a3-721a5e4a2754\",\"type\":\"PAYPAL_ACCOUNT\",\"status\":\"ACTIVATED\",\"createdOn\":\"2021-10-21T14:07:23\",\"transferMethodCountry\":\"US\",\"transferMethodCurrency\":\"USD\",\"userToken\":\"usr-c9d3126d-e26d-459d-9d66-9538876848be\",\"email\":\"94105@asde.com\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/paypal-accounts/trm-0d66f04d-4340-4820-87a3-721a5e4a2754\"}]}"),
		),
	)

	pg := PaypalAccountGateway{testClient}

	paypalAccount, err := pg.RetrievePaypalAccount(ctx, userToken, paypalAccountToken)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, paypalAccount)
	}
}
