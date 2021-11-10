package hyperwallet

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func NewTestBankAccountGateway() *BankAccountGateway {
	return &BankAccountGateway{
		NewClient(),
	}
}

func TestCreateBankAccountDataValidate(t *testing.T) {
	t.Parallel()

	type test struct {
		input    CreateBankAccountData
		expected string
	}

	tests := []test{
		{
			CreateBankAccountData{
				ProfileType:            PROFILE_TYPE_INDIVIDUAL,
				TransferMethodCountry:  "US",
				TransferMethodCurrency: "USD",
				Type:                   "WIRE_ACCOUNT",
				BankID:                 "AsdGas12345",
				BankAccountID:          "987654321",
				FirstName:              "Alex",
				MiddleName:             "Serg",
				LastName:               "Niki",
				Country:                "US",
				StateProvince:          "CA",
				AddressLine1:           "575 Market St",
				City:                   "San Francisco",
				PostalCode:             "94105",
			},
			"",
		},
		{
			CreateBankAccountData{
				ProfileType:            PROFILE_TYPE_INDIVIDUAL,
				TransferMethodCountry:  "US",
				TransferMethodCurrency: "USD",
				Type:                   "WIRE_ACCOUNT",
				BankID:                 "123456789AS",
				BankAccountID:          "987654321",
				FirstName:              "Alex!@!",
				MiddleName:             "Serg",
				LastName:               "Niki",
				Country:                "USlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
				StateProvince:          "CA",
				AddressLine1:           "575 Market St$%",
				City:                   "San Francisco",
				PostalCode:             "94105longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
			},
			"Bad value for BankID\n" +
				"Bad value for FirstName\n" +
				"Bad value for Country\n" +
				"Bad value for AddressLine1\n" +
				"Bad value for PostalCode",
		},
		{
			CreateBankAccountData{
				ProfileType:            PROFILE_TYPE_INDIVIDUAL,
				TransferMethodCurrency: "USD",
				Type:                   "WIRE_ACCOUNT",
				BankID:                 "123456789AS",
				BankAccountID:          "987654321",
				FirstName:              "Alex!@!",
				MiddleName:             "Serg",
				LastName:               "Niki",
				Country:                "USlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
				StateProvince:          "CA",
				AddressLine1:           "575 Market St$%",
				City:                   "San Francisco",
				PostalCode:             "94105longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
			},
			"Key: 'CreateBankAccountData.TransferMethodCountry' Error:Field validation for 'TransferMethodCountry' failed on the 'required' tag",
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

func TestUpdateBankAccountDataValidate(t *testing.T) {
	t.Parallel()

	type test struct {
		input    UpdateBankAccountData
		expected string
	}

	tests := []test{
		{
			UpdateBankAccountData{
				ProfileType:            PROFILE_TYPE_INDIVIDUAL,
				TransferMethodCountry:  "US",
				TransferMethodCurrency: "USD",
				Type:                   "WIRE_ACCOUNT",
				AddressLine1:           "575 Market St",
				City:                   "San Francisco",
				PostalCode:             "94105",
			},
			"",
		},
		{
			UpdateBankAccountData{
				ProfileType:            PROFILE_TYPE_INDIVIDUAL,
				TransferMethodCountry:  "US",
				TransferMethodCurrency: "USD",
				Type:                   "WIRE_ACCOUNT",
				BankID:                 "123456789AS",
				BankAccountID:          "987654321",
				FirstName:              "Alex!@!",
				MiddleName:             "Serg",
				LastName:               "Niki",
				Country:                "USlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
				StateProvince:          "CA",
				AddressLine1:           "575 Market St$%",
				City:                   "San Francisco",
				PostalCode:             "94105longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
			},
			"Bad value for BankID\n" +
				"Bad value for FirstName\n" +
				"Bad value for Country\n" +
				"Bad value for AddressLine1\n" +
				"Bad value for PostalCode",
		},
		{
			UpdateBankAccountData{
				ProfileType:   PROFILE_TYPE_INDIVIDUAL,
				Type:          "WIRE_ACCOUNT",
				BankAccountID: "987654321",
				FirstName:     "Alex!@!",
				MiddleName:    "Serg",
				LastName:      "Niki",
				Country:       "USlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
				StateProvince: "CA",
				AddressLine1:  "575 Market St$%",
				City:          "San Francisco",
				PostalCode:    "94105longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
			},
			"Bad value for FirstName\n" +
				"Bad value for Country\n" +
				"Bad value for AddressLine1\n" +
				"Bad value for PostalCode",
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

func TestGetBankAccountList(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

	expected := &BankAccountList{
		Count:  1,
		Offset: 0,
		Limit:  10,
		Data: []BankAccount{
			{
				Token:                  "trm-ea101b26-f009-4918-857b-19d226381fd9",
				Type:                   "BANK_ACCOUNT",
				Status:                 "ACTIVATED",
				VerificationStatus:     "NOT_REQUIRED",
				CreatedOn:              "2021-10-21T13:19:06",
				TransferMethodCountry:  "US",
				TransferMethodCurrency: "USD",
				BankName:               "WELLS FARGO BANK                    ",
				BranchID:               "101089292",
				BankAccountID:          "****1343",
				BankAccountPurpose:     "SAVINGS",
				UserToken:              "usr-c9d3126d-e26d-459d-9d66-9538876848be",
				ProfileType:            "INDIVIDUAL",
				FirstName:              "Alex",
				MiddleName:             "Serg",
				LastName:               "Niki",
				AddressLine1:           "575 Market St",
				City:                   "San Francisco",
				StateProvince:          "CA",
				Country:                "US",
				PostalCode:             "94105",
				Links: []Link{
					{
						Params: Params{Rel: "self"},
						Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/bank-accounts/trm-ea101b26-f009-4918-857b-19d226381fd9",
					},
				},
			},
		},
	}

	const userToken = "usr-c9d3126d-e26d-459d-9d66-9538876848be"

	httpmock.RegisterRegexpResponder(
		"GET",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/users/"+userToken+"/bank-accounts"),
		httpmock.NewBytesResponder(200,
			[]byte("{\"count\":1,\"offset\":0,\"limit\":10,\"data\":[{\"token\":\"trm-ea101b26-f009-4918-857b-19d226381fd9\",\"type\":\"BANK_ACCOUNT\",\"status\":\"ACTIVATED\",\"verificationStatus\":\"NOT_REQUIRED\",\"createdOn\":\"2021-10-21T13:19:06\",\"transferMethodCountry\":\"US\",\"transferMethodCurrency\":\"USD\",\"bankName\":\"WELLS FARGO BANK                    \",\"branchId\":\"101089292\",\"bankAccountId\":\"****1343\",\"bankAccountPurpose\":\"SAVINGS\",\"userToken\":\"usr-c9d3126d-e26d-459d-9d66-9538876848be\",\"profileType\":\"INDIVIDUAL\",\"firstName\":\"Alex\",\"middleName\":\"Serg\",\"lastName\":\"Niki\",\"addressLine1\":\"575 Market St\",\"city\":\"San Francisco\",\"stateProvince\":\"CA\",\"country\":\"US\",\"postalCode\":\"94105\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/bank-accounts/trm-ea101b26-f009-4918-857b-19d226381fd9\"}]}]}"),
		),
	)

	bg := BankAccountGateway{testClient}

	bankAccountList, err := bg.GetBankAccountList(ctx, userToken, GetBankAccountListQuery{})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, bankAccountList)
	}
}

func TestCreateBankAccount(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

	createBankAccountData := CreateBankAccountData{
		ProfileType:            PROFILE_TYPE_INDIVIDUAL,
		TransferMethodCountry:  "US",
		TransferMethodCurrency: "USD",
		Type:                   "WIRE_ACCOUNT",
		BankID:                 "AsdGas12345",
		BankAccountID:          "987654321",
		FirstName:              "Alex",
		MiddleName:             "Serg",
		LastName:               "Niki",
		Country:                "US",
		StateProvince:          "CA",
		AddressLine1:           "575 Market St",
		City:                   "San Francisco",
		PostalCode:             "94105",
	}

	expected := &BankAccount{
		Token:                  "trm-ea101b26-f009-4918-857b-19d226381fd9",
		Type:                   "BANK_ACCOUNT",
		Status:                 "ACTIVATED",
		VerificationStatus:     "NOT_REQUIRED",
		CreatedOn:              "2021-10-21T13:19:06",
		TransferMethodCountry:  "US",
		TransferMethodCurrency: "USD",
		BankName:               "WELLS FARGO BANK",
		BranchID:               "101089292",
		BankAccountID:          "****1343",
		BankAccountPurpose:     "SAVINGS",
		UserToken:              "usr-c9d3126d-e26d-459d-9d66-9538876848be",
		ProfileType:            "INDIVIDUAL",
		FirstName:              "Alex",
		MiddleName:             "Serg",
		LastName:               "Niki",
		AddressLine1:           "575 Market St",
		City:                   "San Francisco",
		StateProvince:          "CA",
		Country:                "US",
		PostalCode:             "94105",
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/bank-accounts/trm-ea101b26-f009-4918-857b-19d226381fd9",
			},
		},
	}

	const userToken = "usr-c9d3126d-e26d-459d-9d66-9538876848be"

	httpmock.RegisterRegexpResponder(
		"POST",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/users/"+userToken+"/bank-accounts"),
		httpmock.NewBytesResponder(200,
			[]byte("{\"token\":\"trm-ea101b26-f009-4918-857b-19d226381fd9\",\"type\":\"BANK_ACCOUNT\",\"status\":\"ACTIVATED\",\"verificationStatus\":\"NOT_REQUIRED\",\"createdOn\":\"2021-10-21T13:19:06\",\"transferMethodCountry\":\"US\",\"transferMethodCurrency\":\"USD\",\"bankName\":\"WELLS FARGO BANK\",\"branchId\":\"101089292\",\"bankAccountId\":\"****1343\",\"bankAccountPurpose\":\"SAVINGS\",\"userToken\":\"usr-c9d3126d-e26d-459d-9d66-9538876848be\",\"profileType\":\"INDIVIDUAL\",\"firstName\":\"Alex\",\"middleName\":\"Serg\",\"lastName\":\"Niki\",\"addressLine1\":\"575 Market St\",\"city\":\"San Francisco\",\"stateProvince\":\"CA\",\"country\":\"US\",\"postalCode\":\"94105\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/bank-accounts/trm-ea101b26-f009-4918-857b-19d226381fd9\"}]}"),
		),
	)

	bg := BankAccountGateway{testClient}

	bankAccount, err := bg.CreateBankAccount(ctx, userToken, createBankAccountData)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, bankAccount)
	}
}

func TestUpdateBankAccount(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

	updateBankAccountData := UpdateBankAccountData{
		FirstName:              "Testios",
		LastName:               "Tomoto",
		AddressLine1:           "5 Street 34/6",
	}

	expected := &BankAccount{
		Token:                  "trm-ea101b26-f009-4918-857b-19d226381fd9",
		Type:                   "BANK_ACCOUNT",
		Status:                 "ACTIVATED",
		VerificationStatus:     "NOT_REQUIRED",
		CreatedOn:              "2021-10-21T13:19:06",
		TransferMethodCountry:  "US",
		TransferMethodCurrency: "USD",
		BankName:               "WELLS FARGO BANK",
		BranchID:               "101089292",
		BankAccountID:          "****1343",
		BankAccountPurpose:     "SAVINGS",
		UserToken:              "usr-c9d3126d-e26d-459d-9d66-9538876848be",
		ProfileType:            "INDIVIDUAL",
		FirstName:              "Testios",
		MiddleName:             "Serg",
		LastName:               "Tomoto",
		AddressLine1:           "5 Street 34/6",
		City:                   "San Francisco",
		StateProvince:          "CA",
		Country:                "US",
		PostalCode:             "94105",
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/bank-accounts/trm-ea101b26-f009-4918-857b-19d226381fd9",
			},
		},
	}

	const userToken = "usr-c9d3126d-e26d-459d-9d66-9538876848be"
	const bankAccountToken = "trm-ea101b26-f009-4918-857b-19d226381fd9"

	httpmock.RegisterRegexpResponder(
		"PUT",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/users/"+userToken+"/bank-accounts/"+bankAccountToken),
		httpmock.NewBytesResponder(200,
			[]byte("{\"token\":\"trm-ea101b26-f009-4918-857b-19d226381fd9\",\"type\":\"BANK_ACCOUNT\",\"status\":\"ACTIVATED\",\"verificationStatus\":\"NOT_REQUIRED\",\"createdOn\":\"2021-10-21T13:19:06\",\"transferMethodCountry\":\"US\",\"transferMethodCurrency\":\"USD\",\"bankName\":\"WELLS FARGO BANK\",\"branchId\":\"101089292\",\"bankAccountId\":\"****1343\",\"bankAccountPurpose\":\"SAVINGS\",\"userToken\":\"usr-c9d3126d-e26d-459d-9d66-9538876848be\",\"profileType\":\"INDIVIDUAL\",\"firstName\":\"Testios\",\"middleName\":\"Serg\",\"lastName\":\"Tomoto\",\"addressLine1\":\"5 Street 34/6\",\"city\":\"San Francisco\",\"stateProvince\":\"CA\",\"country\":\"US\",\"postalCode\":\"94105\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/bank-accounts/trm-ea101b26-f009-4918-857b-19d226381fd9\"}]}"),
		),
	)

	bg := BankAccountGateway{testClient}

	bankAccount, err := bg.UpdateBankAccount(ctx, userToken, bankAccountToken, updateBankAccountData)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, bankAccount)
	}
}

func TestRetrieveBankAccount(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

	expected := &BankAccount{
		Token:                  "trm-ea101b26-f009-4918-857b-19d226381fd9",
		Type:                   "BANK_ACCOUNT",
		Status:                 "ACTIVATED",
		VerificationStatus:     "NOT_REQUIRED",
		CreatedOn:              "2021-10-21T13:19:06",
		TransferMethodCountry:  "US",
		TransferMethodCurrency: "USD",
		BankName:               "WELLS FARGO BANK",
		BranchID:               "101089292",
		BankAccountID:          "****1343",
		BankAccountPurpose:     "SAVINGS",
		UserToken:              "usr-c9d3126d-e26d-459d-9d66-9538876848be",
		ProfileType:            "INDIVIDUAL",
		FirstName:              "Alex",
		MiddleName:             "Serg",
		LastName:               "Niki",
		AddressLine1:           "575 Market St",
		City:                   "San Francisco",
		StateProvince:          "CA",
		Country:                "US",
		PostalCode:             "94105",
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/bank-accounts/trm-ea101b26-f009-4918-857b-19d226381fd9",
			},
		},
	}

	const userToken = "usr-c9d3126d-e26d-459d-9d66-9538876848be"
	const bankAccountToken = "trm-ea101b26-f009-4918-857b-19d226381fd9"

	httpmock.RegisterRegexpResponder(
		"GET",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/users/"+userToken+"/bank-accounts/"+bankAccountToken),
		httpmock.NewBytesResponder(200,
			[]byte("{\"token\":\"trm-ea101b26-f009-4918-857b-19d226381fd9\",\"type\":\"BANK_ACCOUNT\",\"status\":\"ACTIVATED\",\"verificationStatus\":\"NOT_REQUIRED\",\"createdOn\":\"2021-10-21T13:19:06\",\"transferMethodCountry\":\"US\",\"transferMethodCurrency\":\"USD\",\"bankName\":\"WELLS FARGO BANK\",\"branchId\":\"101089292\",\"bankAccountId\":\"****1343\",\"bankAccountPurpose\":\"SAVINGS\",\"userToken\":\"usr-c9d3126d-e26d-459d-9d66-9538876848be\",\"profileType\":\"INDIVIDUAL\",\"firstName\":\"Alex\",\"middleName\":\"Serg\",\"lastName\":\"Niki\",\"addressLine1\":\"575 Market St\",\"city\":\"San Francisco\",\"stateProvince\":\"CA\",\"country\":\"US\",\"postalCode\":\"94105\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/bank-accounts/trm-ea101b26-f009-4918-857b-19d226381fd9\"}]}"),
		),
	)

	bg := BankAccountGateway{testClient}

	bankAccount, err := bg.RetrieveBankAccount(ctx, userToken, bankAccountToken)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, bankAccount)
	}
}
