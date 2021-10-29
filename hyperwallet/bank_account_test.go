package hyperwallet

import (
	"context"
	"reflect"
	"testing"
)

func NewTestBankAccountGateway() *BankAccountGateway {
	return &BankAccountGateway{
		NewTestClient(),
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
				BankId:                 "AsdGas12345",
				BankAccountId:          "987654321",
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
				BankId:                 "123456789AS",
				BankAccountId:          "987654321",
				FirstName:              "Alex!@!",
				MiddleName:             "Serg",
				LastName:               "Niki",
				Country:                "USlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
				StateProvince:          "CA",
				AddressLine1:           "575 Market St$%",
				City:                   "San Francisco",
				PostalCode:             "94105longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
			},
			"Bad value for BankId\n" +
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
				BankId:                 "123456789AS",
				BankAccountId:          "987654321",
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
				BankId:                 "123456789AS",
				BankAccountId:          "987654321",
				FirstName:              "Alex!@!",
				MiddleName:             "Serg",
				LastName:               "Niki",
				Country:                "USlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
				StateProvince:          "CA",
				AddressLine1:           "575 Market St$%",
				City:                   "San Francisco",
				PostalCode:             "94105longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
			},
			"Bad value for BankId\n" +
				"Bad value for FirstName\n" +
				"Bad value for Country\n" +
				"Bad value for AddressLine1\n" +
				"Bad value for PostalCode",
		},
		{
			UpdateBankAccountData{
				ProfileType:   PROFILE_TYPE_INDIVIDUAL,
				Type:          "WIRE_ACCOUNT",
				BankAccountId: "987654321",
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
	testClient := NewTestBankAccountGateway()

	ctx := context.Background()

	bankAccounts, err := testClient.GetBankAccountList(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", GetBankAccountListQuery{})
	if err != nil {
		t.Errorf("%s", err.Error())
	}

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

	if !reflect.DeepEqual(bankAccounts, expected) {
		t.Errorf("Unexpected result of GetBankAccountList func")
	}
}

func TestCreateBankAccount(t *testing.T) {
	testClient := NewTestBankAccountGateway()

	ctx := context.Background()

	createBankAccountData := CreateBankAccountData{
		ProfileType:            PROFILE_TYPE_INDIVIDUAL,
		TransferMethodCountry:  "US",
		TransferMethodCurrency: "USD",
		Type:                   "WIRE_ACCOUNT",
		BankId:                 "AsdGas12345",
		BankAccountId:          "987654321",
		FirstName:              "Alex",
		MiddleName:             "Serg",
		LastName:               "Niki",
		Country:                "US",
		StateProvince:          "CA",
		AddressLine1:           "575 Market St",
		City:                   "San Francisco",
		PostalCode:             "94105",
	}

	bankAccount, err := testClient.CreateBankAccount(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", createBankAccountData)
	if err != nil {
		t.Errorf("%s", err.Error())
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

	if !reflect.DeepEqual(bankAccount, expected) {
		t.Errorf("Unexpected result of CreateBankAccount func")
	}
}

func TestUpdateBankAccount(t *testing.T) {
	testClient := NewTestBankAccountGateway()

	ctx := context.Background()

	updateBankAccountData := UpdateBankAccountData{
		FirstName:              "Testios",
		LastName:               "Tomoto",
		AddressLine1:           "5 Street 34/6",
	}

	bankAccount, err := testClient.UpdateBankAccount(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", "trm-ea101b26-f009-4918-857b-19d226381fd9", updateBankAccountData)
	if err != nil {
		t.Errorf("%s", err.Error())
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

	if !reflect.DeepEqual(bankAccount, expected) {
		t.Errorf("Unexpected result of UpdateBankAccount func")
	}
}

func TestRetrieveBankAccount(t *testing.T) {
	testClient := NewTestBankAccountGateway()

	ctx := context.Background()

	bankAccount, err := testClient.RetrieveBankAccount(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", "trm-ea101b26-f009-4918-857b-19d226381fd9")
	if err != nil {
		t.Errorf("%s", err.Error())
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

	if !reflect.DeepEqual(bankAccount, expected) {
		t.Errorf("Unexpected result of RetrieveBankAccount func")
	}
}
