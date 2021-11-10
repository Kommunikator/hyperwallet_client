package hyperwallet

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"time"
)

type CreatePaypalAccountData struct {
	TransferMethodCountry  string `json:"transferMethodCountry" validate:"required"`
	TransferMethodCurrency string `json:"transferMethodCurrency" validate:"required"`
	Type                   string `json:"type" validate:"required"`
	Email                  string `json:"email" validate:"required,email"`
}

func (c *CreatePaypalAccountData) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		return err
	}

	var errorTexts []string

	m, err := regexp.MatchString(`^[\w]{3}$`, c.TransferMethodCurrency)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for TransferMethodCurrency")
	}

	if c.Type != "PAYPAL_ACCOUNT" {
		errorTexts = append(errorTexts, "Bad value for Type")
	}

	if len(errorTexts) > 0 {
		return errors.New(strings.Join(errorTexts, "\n"))
	}

	return nil
}

type UpdatePaypalAccountData struct {
	Email                  string `json:"email" validate:"required,email"`
}

func (c *UpdatePaypalAccountData) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type PaypalAccount struct {
	Token                  string `json:"token"`
	Type                   string `json:"type"`
	Status                 string `json:"status"`
	CreatedOn              string `json:"createdOn"`
	TransferMethodCountry  string `json:"transferMethodCountry"`
	TransferMethodCurrency string `json:"transferMethodCurrency"`
	UserToken              string `json:"userToken"`
	Email                  string `json:"email"`
	Links                  []Link `json:"links"`
}

type PaypalAccountList struct {
	Count  int             `json:"count"`
	Offset int             `json:"offset"`
	Limit  int             `json:"limit"`
	Data   []PaypalAccount `json:"data"`
}

type GetPaypalAccountListQuery struct {
	CreatedBefore time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T03:04:05Z"`
	CreatedOn     time.Time `url:"createdOn,omitempty" layout:"2006-01-02T03:04:05Z"`
	CreatedAfter  time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T03:04:05Z"`
	Type          string    `url:"type,omitempty"`
	Status        string    `url:"status,omitempty"`
	SortBy        string    `url:"sortBy,omitempty"`
	Offset        string    `url:"offset,omitempty"`
	Limit         string    `url:"limit,omitempty"`
}
