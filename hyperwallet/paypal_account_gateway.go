package hyperwallet

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	jsoniter "github.com/json-iterator/go"
)

type PaypalAccountGateway struct {
	Client
}

func NewPaypalAccountGateway() *PaypalAccountGateway {
	return &PaypalAccountGateway{
		NewClient(),
	}
}

func (c *PaypalAccountGateway) CreatePaypalAccount(ctx context.Context, userToken string, paypalAccountData CreatePaypalAccountData) (*PaypalAccount, error) {
	body, _ := jsoniter.Marshal(&paypalAccountData)
	responseBody, err := c.Execute(ctx, "POST", fmt.Sprintf("users/%s/paypal-accounts", userToken), nil, string(body))
	if err != nil {
		return nil, err
	}

	paypalAccount := &PaypalAccount{}
	err = jsoniter.Unmarshal(responseBody, &paypalAccount)
	if err != nil {
		return nil, err
	}

	return paypalAccount, nil
}

func (c *PaypalAccountGateway) GetPaypalAccountList(ctx context.Context, userToken string, filters GetPaypalAccountListQuery) (*PaypalAccountList, error) {
	q, err := query.Values(filters)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.Execute(ctx, "GET", fmt.Sprintf("users/%s/paypal-accounts", userToken), q, "")
	if err != nil {
		return nil, err
	}

	paypalAccountList := &PaypalAccountList{}
	err = jsoniter.Unmarshal(responseBody, &paypalAccountList)
	if err != nil {
		return nil, err
	}

	return paypalAccountList, nil
}

func (c *PaypalAccountGateway) RetrievePaypalAccount(ctx context.Context, userToken string, paypalAccountToken string) (*PaypalAccount, error) {
	responseBody, err := c.Execute(ctx, "GET", fmt.Sprintf("users/%s/paypal-accounts/%s", userToken, paypalAccountToken), nil, "")
	if err != nil {
		return nil, err
	}

	paypalAccount := &PaypalAccount{}
	err = jsoniter.Unmarshal(responseBody, &paypalAccount)
	if err != nil {
		return nil, err
	}

	return paypalAccount, nil
}

func (c *PaypalAccountGateway) UpdatePaypalAccount(ctx context.Context, userToken string, paypalAccountToken string, paypalAccountData UpdatePaypalAccountData) (*PaypalAccount, error) {
	body, _ := jsoniter.Marshal(&paypalAccountData)
	responseBody, err := c.Execute(ctx, "PUT", fmt.Sprintf("users/%s/paypal-accounts/%s", userToken, paypalAccountToken), nil, string(body))
	if err != nil {
		return nil, err
	}

	paypalAccount := &PaypalAccount{}
	err = jsoniter.Unmarshal(responseBody, &paypalAccount)
	if err != nil {
		return nil, err
	}

	return paypalAccount, nil
}
