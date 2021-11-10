package hyperwallet

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestCreatePaymentDataValidate(t *testing.T) {
	t.Parallel()

	type test struct {
		input    CreatePaymentData
		expected string
	}

	tests := []test{
		{
			CreatePaymentData{
				Amount:           "100",
				ClientPaymentID:  "163qwe4731sd1568skldfj73asd",
				Currency:         "USD",
				DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
				ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
				Purpose:          "OTHER",
				ExpiresOn:        time.Time{},
				Memo:             "tst payment",
				Notes:            "foo bar",
				ReleaseOn:        time.Time{},
			},
			"",
		},
		{
			CreatePaymentData{
				Amount:           "100",
				ClientPaymentID:  "163qwe4731sd1568skldfj73asd",
				Currency:         "USD",
				DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
				ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
				Purpose:          "OTHER",
			},
			"",
		},
		{
			CreatePaymentData{
				Amount:           "100",
				ClientPaymentID:  "163qwe4731sd1568skldfj73asd",
				Currency:         "USD",
				DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
				ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
				Purpose:          "OTHER",
			},
			"",
		},
		{
			CreatePaymentData{
				ClientPaymentID:  "163qwe4731sd1568skldfj73asd163qwe4731sd1568skldfj73asd163qwe4731sd1568skldfj73asd",
				Currency:         "USDUSD",
				DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
				ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
				Purpose:          "OTHER",
			},
			"Key: 'CreatePaymentData.Amount' Error:Field validation for 'Amount' failed on the 'required' tag",
		},
		{
			CreatePaymentData{
				Amount:           "100",
				ClientPaymentID:  "163qwe4731sd1568skldfj73asd163qwe4731sd1568skldfj73asd163qwe4731sd1568skldfj73asd",
				Currency:         "USDUSD",
				DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
				ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
				Purpose:          "OTHER",
			},
			"Bad value for ClientPaymentID\n" +
				"Bad value for Currency",
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

func TestGetPaymentList(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

	expected := &PaymentList{
		Count:  2,
		Offset: 0,
		Limit:  10,
		Data: []Payment{
			{
				Token:            "pmt-df5f8246-9af8-41aa-873d-34db7d8421c1",
				Status:           "IN_PROGRESS",
				CreatedOn:        "2021-10-25T06:55:01",
				Amount:           "100.00",
				Currency:         "USD",
				ClientPaymentID:  "163qwe4731sd1568skldfj73asd",
				Purpose:          "OTHER",
				ExpiresOn:        "2022-04-23T06:55:01",
				DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
				ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
				Memo:             "",
				Notes:            "",
				ReleaseOn:        "",
				Links: []Link{
					{
						Params: Params{Rel: "self"},
						Href:   "https://api.sandbox.hyperwallet.com/rest/v3/payments/pmt-df5f8246-9af8-41aa-873d-34db7d8421c1",
					},
				},
			},
			{
				Token:            "pmt-476a97ac-882d-4d02-82f9-b7982656295b",
				Status:           "IN_PROGRESS",
				CreatedOn:        "2021-10-25T09:01:38",
				Amount:           "100.00",
				Currency:         "USD",
				ClientPaymentID:  "163qwe4731sd1568skldfj73asd",
				Purpose:          "OTHER",
				ExpiresOn:        "2022-04-23T09:01:38",
				DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
				ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
				Memo:             "",
				Notes:            "",
				ReleaseOn:        "",
				Links: []Link{
					{
						Params: Params{Rel: "self"},
						Href:   "https://api.sandbox.hyperwallet.com/rest/v3/payments/pmt-476a97ac-882d-4d02-82f9-b7982656295b",
					},
				},
			},
		},
	}

	httpmock.RegisterRegexpResponder(
		"GET",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/payments"),
		httpmock.NewBytesResponder(200,
			[]byte("{\"count\":2,\"offset\":0,\"limit\":10,\"data\":[{\"token\":\"pmt-df5f8246-9af8-41aa-873d-34db7d8421c1\",\"status\":\"IN_PROGRESS\",\"createdOn\":\"2021-10-25T06:55:01\",\"amount\":\"100.00\",\"currency\":\"USD\",\"clientPaymentId\":\"163qwe4731sd1568skldfj73asd\",\"purpose\":\"OTHER\",\"expiresOn\":\"2022-04-23T06:55:01\",\"destinationToken\":\"trm-ea101b26-f009-4918-857b-19d226381fd9\",\"programToken\":\"prg-5cd8a525-0553-4e30-8e47-c5440b743855\",\"memo\":\"\",\"notes\":\"\",\"releaseOn\":\"\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/payments/pmt-df5f8246-9af8-41aa-873d-34db7d8421c1\"}]},{\"token\":\"pmt-476a97ac-882d-4d02-82f9-b7982656295b\",\"status\":\"IN_PROGRESS\",\"createdOn\":\"2021-10-25T09:01:38\",\"amount\":\"100.00\",\"currency\":\"USD\",\"clientPaymentId\":\"163qwe4731sd1568skldfj73asd\",\"purpose\":\"OTHER\",\"expiresOn\":\"2022-04-23T09:01:38\",\"destinationToken\":\"trm-ea101b26-f009-4918-857b-19d226381fd9\",\"programToken\":\"prg-5cd8a525-0553-4e30-8e47-c5440b743855\",\"memo\":\"\",\"notes\":\"\",\"releaseOn\":\"\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/payments/pmt-476a97ac-882d-4d02-82f9-b7982656295b\"}]}]}"),
		),
	)

	pg := PaymentGateway{testClient}

	paymentList, err := pg.GetPaymentList(ctx, GetPaymentListQuery{})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, paymentList)
	}
}

func TestCreatePayment(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

		createPaymentData := CreatePaymentData{
			Amount:           "100",
			ClientPaymentID:  "163qwe4731sd1568skldfj73asd",
			Currency:         "USD",
			DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
			ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
			Purpose:          "OTHER",
		}

	expected := &Payment{
		Token:            "pmt-df5f8246-9af8-41aa-873d-34db7d8421c1",
		Status:           "IN_PROGRESS",
		CreatedOn:        "2021-10-25T06:55:01",
		Amount:           "100.00",
		Currency:         "USD",
		ClientPaymentID:  "163qwe4731sd1568skldfj73asd",
		Purpose:          "OTHER",
		ExpiresOn:        "2022-04-23T06:55:01",
		DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
		ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
		Memo:             "",
		Notes:            "",
		ReleaseOn:        "",
		Links:            []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/payments/pmt-df5f8246-9af8-41aa-873d-34db7d8421c1",
			},
		},
	}

	httpmock.RegisterRegexpResponder(
		"POST",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/payments"),
		httpmock.NewBytesResponder(200,
			[]byte("{\"token\":\"pmt-df5f8246-9af8-41aa-873d-34db7d8421c1\",\"status\":\"IN_PROGRESS\",\"createdOn\":\"2021-10-25T06:55:01\",\"amount\":\"100.00\",\"currency\":\"USD\",\"clientPaymentId\":\"163qwe4731sd1568skldfj73asd\",\"purpose\":\"OTHER\",\"expiresOn\":\"2022-04-23T06:55:01\",\"destinationToken\":\"trm-ea101b26-f009-4918-857b-19d226381fd9\",\"programToken\":\"prg-5cd8a525-0553-4e30-8e47-c5440b743855\",\"memo\":\"\",\"notes\":\"\",\"releaseOn\":\"\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/payments/pmt-df5f8246-9af8-41aa-873d-34db7d8421c1\"}]}"),
		),
	)

	pg := PaymentGateway{testClient}

	payment, err := pg.CreatePayment(ctx, createPaymentData)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, payment)
	}
}

func TestRetrievePayment(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	testClient := NewClient()

	httpmock.ActivateNonDefault(testClient.HttpClient)

	expected := &Payment{
		Token:            "pmt-df5f8246-9af8-41aa-873d-34db7d8421c1",
		Status:           "IN_PROGRESS",
		CreatedOn:        "2021-10-25T06:55:01",
		Amount:           "100.00",
		Currency:         "USD",
		ClientPaymentID:  "163qwe4731sd1568skldfj73asd",
		Purpose:          "OTHER",
		ExpiresOn:        "2022-04-23T06:55:01",
		DestinationToken: "trm-ea101b26-f009-4918-857b-19d226381fd9",
		ProgramToken:     "prg-5cd8a525-0553-4e30-8e47-c5440b743855",
		Memo:             "",
		Notes:            "",
		ReleaseOn:        "",
		Links:            []Link{
			{
				Params: Params{Rel: "self"},
				Href:   "https://api.sandbox.hyperwallet.com/rest/v3/payments/pmt-df5f8246-9af8-41aa-873d-34db7d8421c1",
			},
		},
	}

	const paymentToken = "pmt-df5f8246-9af8-41aa-873d-34db7d8421c1"

	httpmock.RegisterRegexpResponder(
		"GET",
		regexp.MustCompile("https://api\\.sandbox\\.hyperwallet\\.com/rest/v3/payments/"+paymentToken),
		httpmock.NewBytesResponder(200,
			[]byte("{\"token\":\"pmt-df5f8246-9af8-41aa-873d-34db7d8421c1\",\"status\":\"IN_PROGRESS\",\"createdOn\":\"2021-10-25T06:55:01\",\"amount\":\"100.00\",\"currency\":\"USD\",\"clientPaymentId\":\"163qwe4731sd1568skldfj73asd\",\"purpose\":\"OTHER\",\"expiresOn\":\"2022-04-23T06:55:01\",\"destinationToken\":\"trm-ea101b26-f009-4918-857b-19d226381fd9\",\"programToken\":\"prg-5cd8a525-0553-4e30-8e47-c5440b743855\",\"memo\":\"\",\"notes\":\"\",\"releaseOn\":\"\",\"links\":[{\"params\":{\"rel\":\"self\"},\"href\":\"https://api.sandbox.hyperwallet.com/rest/v3/payments/pmt-df5f8246-9af8-41aa-873d-34db7d8421c1\"}]}"),
		),
	)

	pg := PaymentGateway{testClient}

	payment, err := pg.RetrievePayment(ctx, paymentToken)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, payment)
	}
}