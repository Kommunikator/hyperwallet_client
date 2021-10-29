package hyperwallet

import (
	"context"
	"reflect"
	"testing"
)

func NewTestUsersGateway() *UsersGateway {
	return &UsersGateway{
		NewTestClient(),
	}
}

func TestCreateUserDataValidate(t *testing.T) {
	t.Parallel()

	type test struct {
		input    CreateUserData
		expected string
	}

	tests := []test{
		{
			CreateUserData{
				ProgramToken:     "tst-prg-123",
				ClientUserId:     "qwerty",
				ProfileType:      PROFILE_TYPE_INDIVIDUAL,
				FirstName:        "Alex",
				LastName:         "Grete",
				DateOfBirth:      "1988-01-05",
				MobileNumber:     "+75001234567",
				Email:            "tst@test.com",
				DriversLicenseId: "1q2w3e4r",
				AddressLine1:     "Pushkina str, 12/54-1",
				City:             "Moscow",
				StateProvince:    "NY",
				Country:          "US",
				PostalCode:       "117968",
			},
			"",
		},
		{
			CreateUserData{
				ProgramToken:     "tst-prg-123",
				ClientUserId:     "qwerty",
				ProfileType:      PROFILE_TYPE_UNKNOWN,
				FirstName:        "",
				LastName:         "",
				DateOfBirth:      "1988-01-05",
				MobileNumber:     "+75001234567",
				Email:            "@test.com",
				DriversLicenseId: "1q2w3e4r",
				AddressLine1:     "Pushkina str, 12/54-1",
				StateProvince:    "NY",
				Country:          "US",
				PostalCode:       "117968",
			},
			"Key: 'CreateUserData.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag\n" +
				"Key: 'CreateUserData.LastName' Error:Field validation for 'LastName' failed on the 'required' tag\n" +
				"Key: 'CreateUserData.Email' Error:Field validation for 'Email' failed on the 'email' tag\n" +
				"Key: 'CreateUserData.City' Error:Field validation for 'City' failed on the 'required' tag",
		},
		{
			CreateUserData{
				ProgramToken:  "tst-prg-123",
				ClientUserId:  "qwerty longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
				ProfileType:   PROFILE_TYPE_UNKNOWN,
				FirstName:     "Alex$%",
				LastName:      "Grete$&",
				DateOfBirth:   "2020-01-05",
				MobileNumber:  "+75001234567",
				Email:         "tst@test.com",
				AddressLine1:  "Pushkina str, 12/54-1 &#",
				City:          "Moscow",
				StateProvince: "NY",
				Country:       "longlonglonglonglonglonglonglonglonglonglonglonglong",
				PostalCode:    "117968",
			},
			"Bad value for ClientUserId\n" +
				"Bad value for ProfileType\n" +
				"Bad value for FirstName\n" +
				"Bad value for LastName\n" +
				"Bad value for DateOfBirth\n" +
				"Bad value for AddressLine1\n" +
				"Bad value for Country",
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

func TestUpdateUserDataValidate(t *testing.T) {
	t.Parallel()

	type test struct {
		input    UpdateUserData
		expected string
	}

	tests := []test{
		{
			UpdateUserData{
				ProgramToken:     "tst-prg-123",
				ClientUserId:     "qwerty",
				ProfileType:      PROFILE_TYPE_INDIVIDUAL,
				FirstName:        "Alex",
				LastName:         "Grete",
				DateOfBirth:      "1988-01-05",
				MobileNumber:     "+75001234567",
				Email:            "tst@test.com",
				DriversLicenseId: "1q2w3e4r",
				AddressLine1:     "Pushkina str, 12/54-1",
				City:             "Moscow",
				StateProvince:    "NY",
				Country:          "US",
				PostalCode:       "117968",
			},
			"",
		},
		{
			UpdateUserData{
				ProgramToken:     "tst-prg-123",
				FirstName:        "",
				LastName:         "",
				DateOfBirth:      "1988-01-05",
				MobileNumber:     "+75001234567",
				Email:            "tst@test.com",
				DriversLicenseId: "1q2w3e4r",
				AddressLine1:     "Pushkina str, 12/54-1",
				StateProvince:    "NY",
				Country:          "US",
				PostalCode:       "117968",
			},
			"",
		},
		{
			UpdateUserData{
				ProgramToken:  "tst-prg-123",
				ClientUserId:  "qwerty longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
				ProfileType:   PROFILE_TYPE_UNKNOWN,
				FirstName:     "Alex$%",
				LastName:      "Grete$&",
				DateOfBirth:   "2020-01-05",
				MobileNumber:  "+75001234567",
				Email:         "tst@test.com",
				AddressLine1:  "Pushkina str, 12/54-1 &#",
				City:          "Moscow",
				StateProvince: "NY",
				Country:       "longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
				PostalCode:    "117968",
			},
			"Bad value for ClientUserId\n" +
				"Bad value for ProfileType\n" +
				"Bad value for FirstName\n" +
				"Bad value for LastName\n" +
				"Bad value for DateOfBirth\n" +
				"Bad value for AddressLine1\n" +
				"Bad value for Country",
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

func TestGetUsers(t *testing.T) {
	testClient := NewTestUsersGateway()
	ctx := context.Background()

	users, err := testClient.GetUserList(ctx, GetUserListQuery{})
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	expected := &UserList{
		Count:  2,
		Offset: 0,
		Limit:  10,
		Data: []User{
			User{
				Token:              "usr-b6979792-22db-4777-8088-f24128833a28",
				Status:             "PRE_ACTIVATED",
				VerificationStatus: "NOT_REQUIRED",
				CreatedOn:          "2021-10-20T14:22:04",
				ClientUserID:       "1634731156873",
				ProfileType:        "INDIVIDUAL",
				FirstName:          "Alex",
				LastName:           "Niki",
				DateOfBirth:        "1989-01-02",
				Email:              "exa@gmail.comm",
				AddressLine1:       "575 Market St",
				City:               "San Francisco",
				StateProvince:      "CA",
				Country:            "US",
				PostalCode:         "94105",
				Language:           "en",
				TimeZone:           "GMT",
				ProgramToken:       "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
				Links: []Link{
					{
						Params: Params{Rel: "self"},
						Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-b6979792-22db-4777-8088-f24128833a28",
					},
				},
			},
			{
				Token:              "usr-bc8310f4-58ad-437b-a2f9-4865a0b61d3d",
				Status:             "PRE_ACTIVATED",
				VerificationStatus: "NOT_REQUIRED",
				CreatedOn:          "2021-10-21T10:24:41",
				ClientUserID:       "16347311568skldfj73",
				ProfileType:        "INDIVIDUAL",
				FirstName:          "Alexius",
				LastName:           "Nikifd",
				DateOfBirth:        "1989-01-03",
				Email:              "edsxa@gmail.comm",
				AddressLine1:       "575 Market St",
				City:               "San Francisco",
				StateProvince:      "CA",
				Country:            "US",
				PostalCode:         "94105",
				Language:           "en",
				TimeZone:           "GMT",
				ProgramToken:       "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
				Links: []Link{
					{
						Params: Params{Rel: "self"},
						Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-bc8310f4-58ad-437b-a2f9-4865a0b61d3d",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(users, expected) {
		t.Errorf("Unexpected result of GetUserList func")
	}
}

func TestCreateUser(t *testing.T) {
	testClient := NewTestUsersGateway()
	ctx := context.Background()

	createUserData := CreateUserData{
		ProgramToken:  "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
		ClientUserId:  "qwer123tsd",
		ProfileType:   PROFILE_TYPE_INDIVIDUAL,
		FirstName:     "Testius",
		LastName:      "Rewersus",
		DateOfBirth:   "1985-01-03",
		Email:         "qwe@gmail.comm",
		AddressLine1:  "575 Market St",
		City:          "San Francisco",
		StateProvince: "CA",
		Country:       "US",
		PostalCode:    "94105",
	}

	expected := &User{
		Token:              "usr-2bb8b9d8-f3c3-43fc-a3db-d473ac57a58e",
		Status:             "PRE_ACTIVATED",
		VerificationStatus: "NOT_REQUIRED",
		CreatedOn:          "2021-10-29T09:45:26",
		ClientUserID:       "qwer123tsd",
		ProfileType:        "INDIVIDUAL",
		FirstName:          "Testius",
		LastName:           "Rewersus",
		DateOfBirth:        "1985-01-03",
		Email:              "qwe123@gmail.comm",
		AddressLine1:       "575 Market St",
		City:               "San Francisco",
		StateProvince:      "CA",
		Country:            "US",
		PostalCode:         "94105",
		Language:           "en",
		TimeZone:           "GMT",
		ProgramToken:       "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-2bb8b9d8-f3c3-43fc-a3db-d473ac57a58e",
			},
		},
	}

	user, err := testClient.CreateUser(ctx, createUserData)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("Unexpected result of CreateUser func")
	}
}

func TestRetrieveUser(t *testing.T) {
	testClient := NewTestUsersGateway()
	ctx := context.Background()

	expected := &User{
		Token:              "usr-c9d3126d-e26d-459d-9d66-9538876848be",
		Status:             "PRE_ACTIVATED",
		VerificationStatus: "NOT_REQUIRED",
		CreatedOn:          "2021-10-21T10:44:27",
		ClientUserID:       "163qwe4731sd1568skldfj73",
		ProfileType:        "INDIVIDUAL",
		FirstName:          "Alexius greate",
		LastName:           "Nikifd ddd",
		DateOfBirth:        "1990-01-03",
		Email:              "ewdsxva@gmail.comm",
		AddressLine1:       "575 Market St",
		City:               "San Francisco",
		StateProvince:      "CA",
		Country:            "US",
		PostalCode:         "94105",
		Language:           "en",
		TimeZone:           "GMT",
		ProgramToken:       "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be",
			},
		},
	}

	user, err := testClient.RetrieveUser(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be")
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("Unexpected result of RetrieveUser func")
	}
}

func TestCreateAuthenticationToken(t *testing.T) {
	testClient := NewTestUsersGateway()
	ctx := context.Background()

	expected := &AuthenticationToken{Value: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ1c3ItYzlkMzEyNmQtZTI2ZC00NTlkLTlkNjYtOTUzODg3Njg0OGJlIiwiaWF0IjoxNjM1NTEyODA3LCJleHAiOjE2MzU1MTM0MDcsImF1ZCI6InBndS0zOWVmZDIwNi1mNjk1LTQwMDItYTEwZS04YzNhZTI2ZGUzZmEiLCJpc3MiOiJwcmctNWNkOGE1MjUtMDU1My00ZTMwLThlNDctYzU0NDBiNzQzODU1IiwicmVzdC11cmkiOiJodHRwczovL2FwaS5zYW5kYm94Lmh5cGVyd2FsbGV0LmNvbS9yZXN0L3YzLyIsImdyYXBocWwtdXJpIjoiaHR0cHM6Ly9hcGkuc2FuZGJveC5oeXBlcndhbGxldC5jb20vZ3JhcGhxbCIsImluc2lnaHRzLXVyaSI6Imh0dHBzOi8vYXBpLnBheXBhbC5jb20vdjEvdHJhY2tpbmcvYmF0Y2gvZXZlbnRzIiwiZW52aXJvbm1lbnQiOiJVQVQiLCJwcm9ncmFtLW1vZGVsIjoiRElSRUNUX0RFUE9TSVRfTU9ERUwifQ.axhkW3uZlssdJtaWjGX5ivFHxvue28xngvb1fLpL9J3shQ_AvdHG1PWlmRvEkGY4_A4eVaFePVazGIt_Xqs9Kg"}

	token, err := testClient.CreateAuthenticationToken(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be")
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(token, expected) {
		t.Errorf("Unexpected result of CreateAuthenticationToken func")
	}
}

func TestGetUserBalanceList(t *testing.T) {
	testClient := NewTestUsersGateway()
	ctx := context.Background()

	expected := &UserBalanceList{
		Count:  1,
		Offset: 0,
		Limit:  10,
		Data: []UserBalance{
			{
				Currency: "USD",
				Amount:   "0.00",
			},
		},
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/balances?offset=0&limit=10",
			},
		},
	}

	balanceList, err := testClient.GetUserBalanceList(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", GetUserBalanceListQuery{})
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(balanceList, expected) {
		t.Errorf("Unexpected result of GetUserBalanceList func")
	}
}

func TestGetUserReceiptList(t *testing.T) {
	testClient := NewTestUsersGateway()
	ctx := context.Background()

	expected := &UserReceiptList{
		Count:  4,
		Offset: 0,
		Limit:  10,
		Data: []UserReceipt{
			{
				Token:            "rcp-4c2bbee8-efc5-476f-9de7-afbeb9dd0610",
				JournalID:        "9192071",
				Type:             "PAYMENT",
				CreatedOn:        "2021-10-25T06:55:03",
				Entry:            "CREDIT",
				SourceToken:      "act-54f85d42-b564-41c5-8965-bc6a3de32877",
				DestinationToken: "usr-c9d3126d-e26d-459d-9d66-9538876848be",
				Amount:           "100.00",
				Fee:              "0.00",
				Currency:         "USD",
				Details: struct {
					ClientPaymentID string `json:"clientPaymentId"`
    				PayeeName       string `json:"payeeName"`
				}{
					ClientPaymentID: "163qwe4731sd1568skldfj73asd",
					PayeeName: "Alexius greate Nikifd ddd",
				},
			},
			{
				Token:            "rcp-d83ba620-1504-4499-b32f-46d21ec426fd",
				JournalID:        "9192072",
				Type:             "TRANSFER_TO_BANK_ACCOUNT",
				CreatedOn:        "2021-10-25T06:55:03",
				Entry:            "DEBIT",
				SourceToken:      "usr-c9d3126d-e26d-459d-9d66-9538876848be",
				DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
				Amount:           "100.00",
				Fee:              "0.00",
				Currency:         "USD",
				Details: struct {
					ClientPaymentID string `json:"clientPaymentId"`
					PayeeName       string `json:"payeeName"`
				}{
					ClientPaymentID: "",
					PayeeName: "Alex Serg Niki",
				},
			},
			{
				Token:            "rcp-c882f630-b97e-422e-a82c-c6bb2d35fbda",
				JournalID:        "9192401",
				Type:             "PAYMENT",
				CreatedOn:        "2021-10-25T09:01:40",
				Entry:            "CREDIT",
				SourceToken:      "act-54f85d42-b564-41c5-8965-bc6a3de32877",
				DestinationToken: "usr-c9d3126d-e26d-459d-9d66-9538876848be",
				Amount:           "100.00",
				Fee:              "0.00",
				Currency:         "USD",
				Details: struct {
					ClientPaymentID string `json:"clientPaymentId"`
					PayeeName       string `json:"payeeName"`
				}{
					ClientPaymentID: "163qwe4731cz3d1568skldfj73asd",
					PayeeName: "Alexius greate Nikifd ddd",
				},
			},
			{
				Token:            "rcp-0ca4345d-8711-42f3-a27b-d27822b0f15a",
				JournalID:        "9192402",
				Type:             "TRANSFER_TO_BANK_ACCOUNT",
				CreatedOn:        "2021-10-25T09:01:40",
				Entry:            "DEBIT",
				SourceToken:      "usr-c9d3126d-e26d-459d-9d66-9538876848be",
				DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
				Amount:           "100.00",
				Fee:              "0.00",
				Currency:         "USD",
				Details: struct {
					ClientPaymentID string `json:"clientPaymentId"`
					PayeeName       string `json:"payeeName"`
				}{
					ClientPaymentID: "",
					PayeeName: "Alex Serg Niki",
				},
			},
		},
		Links: []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/users/usr-c9d3126d-e26d-459d-9d66-9538876848be/receipts?offset=0&limit=10",
			},
		},
	}

	balanceList, err := testClient.GetUserReceiptList(ctx, "usr-c9d3126d-e26d-459d-9d66-9538876848be", GetUserReceiptListQuery{})
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(balanceList, expected) {
		t.Errorf("Unexpected result of GetUserBalanceList func")
	}
}
