package hyperwallet

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"time"
)

type BankAccountList struct {
	Count  int           `json:"count"`
	Offset int           `json:"offset"`
	Limit  int           `json:"limit"`
	Data   []BankAccount `json:"data"`
}

type BankAccount struct {
	Token                  string `json:"token"`
	Type                   string `json:"type"`
	Status                 string `json:"status"`
	VerificationStatus     string `json:"verificationStatus"`
	CreatedOn              string `json:"createdOn"`
	TransferMethodCountry  string `json:"transferMethodCountry"`
	TransferMethodCurrency string `json:"transferMethodCurrency"`
	BankName               string `json:"bankName"`
	BranchID               string `json:"branchId"`
	BankAccountID          string `json:"bankAccountId"`
	BankAccountPurpose     string `json:"bankAccountPurpose"`
	UserToken              string `json:"userToken"`
	ProfileType            string `json:"profileType"`
	FirstName              string `json:"firstName"`
	MiddleName             string `json:"middleName"`
	LastName               string `json:"lastName"`
	AddressLine1           string `json:"addressLine1"`
	City                   string `json:"city"`
	StateProvince          string `json:"stateProvince"`
	Country                string `json:"country"`
	PostalCode             string `json:"postalCode"`
	Links                  []Link `json:"links"`
}

type CreateBankAccountData struct {
	ProfileType            string `json:"profileType" validate:"required"`
	TransferMethodCountry  string `json:"transferMethodCountry" validate:"required"`
	TransferMethodCurrency string `json:"transferMethodCurrency" validate:"required"`
	Type                   string `json:"type" validate:"required"`
	BankID                 string `json:"bankId" validate:"required"`
	BankAccountID          string `json:"bankAccountId" validate:"required"`
	FirstName              string `json:"firstName" validate:"required"`
	MiddleName             string `json:"middleName" validate:"omitempty"`
	LastName               string `json:"lastName" validate:"required"`
	Country                string `json:"country" validate:"required"`
	StateProvince          string `json:"stateProvince" validate:"required"`
	AddressLine1           string `json:"addressLine1" validate:"required"`
	AddressLine2           string `json:"addressLine2,omitempty" validate:"omitempty"`
	City                   string `json:"city" validate:"required"`
	PostalCode             string `json:"postalCode" validate:"required"`
}

func (c *CreateBankAccountData) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		return err
	}

	var m bool
	var errorTexts []string

	m, err = regexp.MatchString(`^INDIVIDUAL|BUSINESS$`, c.ProfileType)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for ProfileType")
	}

	m, err = regexp.MatchString(`^[\w]{2}$`, c.TransferMethodCountry)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for TransferMethodCountry")
	}

	m, err = regexp.MatchString(`^[\w]{3}$`, c.TransferMethodCurrency)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for TransferMethodCurrency")
	}

	m, err = regexp.MatchString(`^WIRE_ACCOUNT$`, c.Type)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for Type")
	}

	m, err = regexp.MatchString(`^[a-zA-Z]{6}[0-9a-zA-Z]{5}$`, c.BankID)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for BankID")
	}

	m, err = regexp.MatchString(`^[a-zA-Z0-9-]{1,34}$`, c.BankAccountID)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for BankAccountID")
	}

	m, err = regexp.MatchString(`^[\w',\-. ]{1,50}$`, c.FirstName)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for FirstName")
	}

	if c.MiddleName != "" {
		m, err = regexp.MatchString(`^[\w',\-. ]{1,50}$`, c.MiddleName)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for FirstName")
		}

	}

	m, err = regexp.MatchString(`^[\w',\-. ]{1,50}$`, c.LastName)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for FirstName")
	}

	m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.Country)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for Country")
	}

	if c.StateProvince != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.StateProvince)
		if err != nil {
			return err
		}
		if m == false || (c.Country == "US" && len(c.StateProvince) != 2) {
			errorTexts = append(errorTexts, "Bad value for StateProvince")
		}
	}

	m, err = regexp.MatchString(`^[\w #'(),\-./:;°]{1,100}$`, c.AddressLine1)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for AddressLine1")
	}

	if c.AddressLine2 != "" {
		m, err = regexp.MatchString(`^[\w #'(),\-./:;°]{1,100}$`, c.AddressLine2)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for AddressLine2")
		}
	}

	m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.City)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for City")
	}

	if c.PostalCode != "" {
		m, err = regexp.MatchString(`^[\w \-]{1,16}$`, c.PostalCode)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for PostalCode")
		}
	}

	if len(errorTexts) > 0 {
		return errors.New(strings.Join(errorTexts, "\n"))
	}

	return nil
}

type UpdateBankAccountData struct {
	ProfileType            string `json:"profileType" validate:"omitempty"`
	TransferMethodCountry  string `json:"transferMethodCountry" validate:"omitempty"`
	TransferMethodCurrency string `json:"transferMethodCurrency" validate:"omitempty"`
	Type                   string `json:"type" validate:"omitempty"`
	BankID                 string `json:"bankId" validate:"omitempty"`
	BankAccountID          string `json:"bankAccountId" validate:"omitempty"`
	FirstName              string `json:"firstName" validate:"omitempty"`
	MiddleName             string `json:"middleName" validate:"omitempty"`
	LastName               string `json:"lastName" validate:"omitempty"`
	Country                string `json:"country" validate:"omitempty"`
	StateProvince          string `json:"stateProvince" validate:"omitempty"`
	AddressLine1           string `json:"addressLine1" validate:"omitempty"`
	AddressLine2           string `json:"addressLine2,omitempty" validate:"omitempty"`
	City                   string `json:"city" validate:"omitempty"`
	PostalCode             string `json:"postalCode" validate:"omitempty"`
}

func (c *UpdateBankAccountData) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		return err
	}

	var m bool
	var errorTexts []string

	if c.ProfileType != "" {
		m, err = regexp.MatchString(`^INDIVIDUAL|BUSINESS$`, c.ProfileType)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for ProfileType")
		}
	}

	if c.TransferMethodCountry != "" {
		m, err = regexp.MatchString(`^[\w]{2}$`, c.TransferMethodCountry)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for TransferMethodCountry")
		}
	}

	if c.TransferMethodCurrency != "" {
		m, err = regexp.MatchString(`^[\w]{3}$`, c.TransferMethodCurrency)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for TransferMethodCurrency")
		}
	}

	if c.Type != "" {
		m, err = regexp.MatchString(`^WIRE_ACCOUNT$`, c.Type)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for Type")
		}
	}

	if c.BankID != "" {
		m, err = regexp.MatchString(`^[a-zA-Z]{4}[a-zA-Z]{2}[0-9a-zA-Z]{2}[0-9a-zA-Z]{3}$`, c.BankID)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BankID")
		}
	}

	if c.BankAccountID != "" {
		m, err = regexp.MatchString(`^[a-zA-Z0-9-]{1,34}$`, c.BankAccountID)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BankAccountID")
		}
	}

	if c.FirstName != "" {
		m, err = regexp.MatchString(`^[\w',\-. ]{1,50}$`, c.FirstName)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for FirstName")
		}
	}

	if c.MiddleName != "" {
		m, err = regexp.MatchString(`^[\w',\-. ]{1,50}$`, c.MiddleName)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for FirstName")
		}

	}

	if c.LastName != "" {
		m, err = regexp.MatchString(`^[\w',\-. ]{1,50}$`, c.LastName)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for FirstName")
		}
	}

	if c.Country != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.Country)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for Country")
		}
	}

	if c.StateProvince != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.StateProvince)
		if err != nil {
			return err
		}
		if m == false || (c.Country == "US" && len(c.StateProvince) != 2) {
			errorTexts = append(errorTexts, "Bad value for StateProvince")
		}
	}

	if c.AddressLine1 != "" {
		m, err = regexp.MatchString(`^[\w #'(),\-./:;°]{1,100}$`, c.AddressLine1)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for AddressLine1")
		}
	}

	if c.AddressLine2 != "" {
		m, err = regexp.MatchString(`^[\w #'(),\-./:;°]{1,100}$`, c.AddressLine2)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for AddressLine2")
		}
	}

	if c.City != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.City)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for City")
		}
	}

	if c.PostalCode != "" {
		m, err = regexp.MatchString(`^[\w \-]{1,16}$`, c.PostalCode)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for PostalCode")
		}
	}

	if len(errorTexts) > 0 {
		return errors.New(strings.Join(errorTexts, "\n"))
	}

	return nil
}

type GetBankAccountListQuery struct {
	CreatedBefore time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T15:04:05"`
	CreatedAfter  time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T15:04:05"`
	Type          string    `url:"type,omitempty"`
	Status        string    `url:"status,omitempty"`
	SortBy        string    `url:"sortBy,omitempty"`
	Offset        string    `url:"offset,omitempty"`
	Limit         string    `url:"limit,omitempty"`
}
