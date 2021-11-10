package hyperwallet

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	jsoniter "github.com/json-iterator/go"
)

type UsersGateway struct {
	*Hyperwallet
}

func NewUsersGateway() *UsersGateway {
	return &UsersGateway{NewClient()}
}

func (c *UsersGateway) CreateUser(ctx context.Context, userData CreateUserData) (*User, error) {
	err := userData.Validate()
	if err != nil {
		return nil, err
	}

	body, err := jsoniter.Marshal(&userData)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.Execute(ctx, "POST", "users", nil, string(body))
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = jsoniter.Unmarshal(responseBody, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UsersGateway) GetUserList(ctx context.Context, filters GetUserListQuery) (*UserList, error) {
	q, err := query.Values(filters)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.Execute(ctx, "GET", "users", q, "")
	if err != nil {
		return nil, err
	}

	userList := &UserList{}
	err = jsoniter.Unmarshal(responseBody, &userList)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (c *UsersGateway) RetrieveUser(ctx context.Context, userToken string) (*User, error) {
	responseBody, err := c.Execute(ctx, "GET", fmt.Sprintf("users/%s", userToken), nil, "")
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = jsoniter.Unmarshal(responseBody, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UsersGateway) UpdateUser(ctx context.Context, userToken string, userData UpdateUserData) (*User, error) {
	err := userData.Validate()
	if err != nil {
		return nil, err
	}

	body, err := jsoniter.Marshal(&userData)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.Execute(ctx, "PUT", fmt.Sprintf("users/%s", userToken), nil, string(body))
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = jsoniter.Unmarshal(responseBody, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UsersGateway) CreateAuthenticationToken(ctx context.Context, userToken string) (*AuthenticationToken, error) {
	responseBody, err := c.Execute(ctx, "POST", fmt.Sprintf("users/%s/authentication-token", userToken), nil, "")
	if err != nil {
		return nil, err
	}

	token := &AuthenticationToken{}
	err = jsoniter.Unmarshal(responseBody, &token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (c *UsersGateway) GetUserBalanceList(ctx context.Context, userToken string, filters GetUserBalanceListQuery) (*UserBalanceList, error) {
	q, err := query.Values(filters)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.Execute(ctx, "GET", fmt.Sprintf("users/%s/balances", userToken), q, "")
	if err != nil {
		return nil, err
	}

	userBalanceList := &UserBalanceList{}
	err = jsoniter.Unmarshal(responseBody, &userBalanceList)
	if err != nil {
		return nil, err
	}

	return userBalanceList, nil
}

func (c *UsersGateway) GetUserReceiptList(ctx context.Context, userToken string, filters GetUserReceiptListQuery) (*UserReceiptList, error) {
	q, err := query.Values(filters)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.Execute(ctx, "GET", fmt.Sprintf("users/%s/receipts", userToken), q, "")
	if err != nil {
		return nil, err
	}

	userReceiptList := &UserReceiptList{}
	err = jsoniter.Unmarshal(responseBody, &userReceiptList)
	if err != nil {
		return nil, err
	}

	return userReceiptList, nil
}
