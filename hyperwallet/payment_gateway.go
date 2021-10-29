package hyperwallet

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	jsoniter "github.com/json-iterator/go"
)

type PaymentGateway struct {
	Client
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{
		NewClient(),
	}
}

func (c *PaymentGateway) CreatePayment(ctx context.Context, paymentData CreatePaymentData) (*Payment, error) {
	body, _ := jsoniter.Marshal(&paymentData)
	responseBody, err := c.Execute(ctx, "POST", "payments", nil, string(body))
	if err != nil {
		return nil, err
	}

	payment := &Payment{}
	err = jsoniter.Unmarshal(responseBody, &payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (c *PaymentGateway) GetPaymentList(ctx context.Context, filters GetPaymentListQuery) (*PaymentList, error) {
	q, err := query.Values(filters)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.Execute(ctx, "GET", "payments", q, "")
	if err != nil {
		return nil, err
	}

	paymentList := &PaymentList{}
	err = jsoniter.Unmarshal(responseBody, &paymentList)
	if err != nil {
		return nil, err
	}

	return paymentList, nil
}

func (c *PaymentGateway) RetrievePayment(ctx context.Context, paymentToken string) (*Payment, error) {
	responseBody, err := c.Execute(ctx, "GET", fmt.Sprintf("payments/%s", paymentToken), nil, "")
	if err != nil {
		return nil, err
	}

	payment := &Payment{}
	err = jsoniter.Unmarshal(responseBody, &payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
