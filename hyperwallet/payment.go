package hyperwallet

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"time"
)

type CreatePaymentData struct {
	Amount           string    `json:"amount" validate:"required,numeric"`
	ClientPaymentID  string    `json:"clientPaymentId" validate:"required"`
	Currency         string    `json:"currency" validate:"required"`
	DestinationToken string    `json:"destinationToken" validate:"required"`
	ProgramToken     string    `json:"programToken" validate:"required"`
	Purpose          string    `json:"purpose" validate:"required"`
	ExpiresOn        time.Time `json:"expiresOn" layout:"2006-01-02T15:04:05" validate:"omitempty"`
	Memo             string    `json:"memo" validate:"omitempty"`
	Notes            string    `json:"notes" validate:"omitempty"`
	ReleaseOn        time.Time `json:"releaseOn" layout:"2006-01-02T15:04:05" validate:"omitempty"`
}

func (c *CreatePaymentData) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		return err
	}

	var errorTexts []string

	m, err := regexp.MatchString(`^[\d]{1,20}$`, c.Amount)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for ClientUserID")
	}

	if len(c.ClientPaymentID) > 50 {
		errorTexts = append(errorTexts, "Bad value for ClientPaymentID")
	}

	m, err = regexp.MatchString(`^[\w]{3}$`, c.Currency)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for Currency")
	}

	if len(errorTexts) > 0 {
		return errors.New(strings.Join(errorTexts, "\n"))
	}

	return nil
}

type Payment struct {
	Token            string `json:"token"`
	Status           string `json:"status"`
	CreatedOn        string `json:"createdOn"`
	Amount           string `json:"amount"`
	Currency         string `json:"currency"`
	ClientPaymentID  string `json:"clientPaymentId"`
	Purpose          string `json:"purpose"`
	ExpiresOn        string `json:"expiresOn"`
	DestinationToken string `json:"destinationToken"`
	ProgramToken     string `json:"programToken"`
	Memo             string `json:"memo"`
	Notes            string `json:"notes"`
	ReleaseOn        string `json:"releaseOn"`
	Links            []Link `json:"links"`
}

type PaymentList struct {
	Count  int       `json:"count"`
	Offset int       `json:"offset"`
	Limit  int       `json:"limit"`
	Data   []Payment `json:"data"`
}

type GetPaymentListQuery struct {
	ClientUserID  string    `url:"clientUserId,omitempty"`
	CreatedBefore time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T15:04:05"`
	CreatedAfter  time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T15:04:05"`
	Currency      string    `url:"currency,omitempty"`
	Memo          string    `url:"memo,omitempty"`
	ReleaseDate   time.Time `url:"releaseDate,omitempty" layout:"2006-01-02T15:04:05"`
	SortBy        string    `url:"sortBy,omitempty"`
	Offset        string    `url:"offset,omitempty"`
	Limit         string    `url:"limit,omitempty"`
}
