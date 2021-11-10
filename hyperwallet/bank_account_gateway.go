package hyperwallet

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	jsoniter "github.com/json-iterator/go"
)

type BankAccountGateway struct {
	*Hyperwallet
}

func NewBankAccountGateway() *BankAccountGateway {
	return &BankAccountGateway{NewClient()}
}

func (c *BankAccountGateway) CreateBankAccount(ctx context.Context, userToken string, bankAccountData CreateBankAccountData) (*BankAccount, error) {
	err := bankAccountData.Validate()
	if err != nil {
		return nil, err
	}

	body, err := jsoniter.Marshal(&bankAccountData)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.Execute(ctx, "POST", fmt.Sprintf("users/%s/bank-accounts", userToken), nil, string(body))
	if err != nil {
		return nil, err
	}

	bankAccount := &BankAccount{}
	err = jsoniter.Unmarshal(responseBody, &bankAccount)
	if err != nil {
		return nil, err
	}

	return bankAccount, nil
}

func (c *BankAccountGateway) GetBankAccountList(ctx context.Context, userToken string, filters GetBankAccountListQuery) (*BankAccountList, error) {
	q, err := query.Values(filters)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.Execute(ctx, "GET", fmt.Sprintf("users/%s/bank-accounts", userToken), q, "")
	if err != nil {
		return nil, err
	}

	bankAccountList := &BankAccountList{}
	err = jsoniter.Unmarshal(responseBody, &bankAccountList)
	if err != nil {
		return nil, err
	}

	return bankAccountList, nil
}

func (c *BankAccountGateway) RetrieveBankAccount(ctx context.Context, userToken string, bankAccountToken string) (*BankAccount, error) {
	responseBody, err := c.Execute(ctx, "GET", fmt.Sprintf("users/%s/bank-accounts/%s", userToken, bankAccountToken), nil,"")
	if err != nil {
		return nil, err
	}

	bankAccount := &BankAccount{}
	err = jsoniter.Unmarshal(responseBody, &bankAccount)
	if err != nil {
		return nil, err
	}

	return bankAccount, nil
}

func (c *BankAccountGateway) UpdateBankAccount(ctx context.Context, userToken string, bankAccountToken string, bankAccountData UpdateBankAccountData) (*BankAccount, error) {
	err := bankAccountData.Validate()
	if err != nil {
		return nil, err
	}

	body, err := jsoniter.Marshal(&bankAccountData)
	responseBody, err := c.Execute(ctx, "PUT", fmt.Sprintf("users/%s/bank-accounts/%s", userToken, bankAccountToken), nil, string(body))
	if err != nil {
		return nil, err
	}

	bankAccount := &BankAccount{}
	err = jsoniter.Unmarshal(responseBody, &bankAccount)
	if err != nil {
		return nil, err
	}

	return bankAccount, nil
}
