package hyperwallet

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"time"
)

const (
	PROFILE_TYPE_INDIVIDUAL             = "INDIVIDUAL"
	PROFILE_TYPE_BUSINESS               = "BUSINESS"
	PROFILE_TYPE_UNKNOWN                = "UNKNOWN"
	BUSINESS_CONTACT_ROLE_DIRECTOR      = "DIRECTOR"
	BUSINESS_CONTACT_ROLE_OWNER         = "OWNER"
	BUSINESS_CONTACT_ROLE_OTHER         = "OTHER"
	GOVERNMENT_ID_TYPE_PASSPORT         = "PASSPORT"
	GOVERNMENT_ID_TYPE_NATIONAL_ID_CARD = "NATIONAL_ID_CARD"
	BUSINESS_TYPE_CORPORATION           = "CORPORATION"
	BUSINESS_TYPE_PARTNERSHIP           = "PARTNERSHIP"
	BUSINESS_TYPE_PRIVATE_COMPANY       = "PRIVATE_COMPANY"
)

type CreateUserData struct {
	ProgramToken                      string `json:"programToken,omitempty"  validate:"required"`
	ClientUserID                      string `json:"clientUserId,omitempty"  validate:"required"`
	ProfileType                       string `json:"profileType,omitempty" validate:"required,alpha,uppercase"`
	FirstName                         string `json:"firstName,omitempty" validate:"required"`
	MiddleName                        string `json:"middleName,omitempty" validate:"omitempty"`
	LastName                          string `json:"lastName,omitempty" validate:"required"`
	DateOfBirth                       string `json:"dateOfBirth,omitempty" validate:"required"`
	CountryOfBirth                    string `json:"countryOfBirth,omitempty" validate:"omitempty"`
	CountryOfNationality              string `json:"countryOfNationality,omitempty" validate:"omitempty"`
	PhoneNumber                       string `json:"phoneNumber,omitempty" validate:"omitempty"`
	MobileNumber                      string `json:"mobileNumber,omitempty" validate:"omitempty"`
	Email                             string `json:"email,omitempty" validate:"required,email"`
	GovernmentID                      string `json:"governmentId,omitempty" validate:"omitempty"`
	GovernmentIDType                  string `json:"governmentIdType,omitempty" validate:"omitempty"`
	PassportID                        string `json:"passportId,omitempty" validate:"omitempty"`
	DriversLicenseID                  string `json:"driversLicenseId,omitempty" validate:"omitempty"`
	EmployerID                        string `json:"employerId,omitempty" validate:"omitempty"`
	AddressLine1                      string `json:"addressLine1,omitempty" validate:"required"`
	AddressLine2                      string `json:"addressLine2,omitempty" validate:"omitempty"`
	City                              string `json:"city,omitempty" validate:"required"`
	StateProvince                     string `json:"stateProvince,omitempty" validate:"required"`
	Country                           string `json:"country,omitempty" validate:"required"`
	PostalCode                        string `json:"postalCode,omitempty" validate:"required"`
	Language                          string `json:"language,omitempty" validate:"omitempty"`
	BusinessType                      string `json:"businessType,omitempty" validate:"omitempty"`
	BusinessName                      string `json:"businessName,omitempty" validate:"omitempty"`
	BusinessOperatingName             string `json:"businessOperatingName,omitempty" validate:"omitempty"`
	BusinessRegistrationID            string `json:"businessRegistrationId,omitempty" validate:"omitempty"`
	BusinessRegistrationStateProvince string `json:"businessRegistrationStateProvince,omitempty" validate:"omitempty"`
	BusinessRegistrationCountry       string `json:"businessRegistrationCountry,omitempty" validate:"omitempty"`
	BusinessContactRole               string `json:"businessContactRole,omitempty" validate:"omitempty"`
	BusinessContactAddressLine1       string `json:"businessContactAddressLine1,omitempty" validate:"omitempty"`
	BusinessContactAddressLine2       string `json:"businessContactAddressLine2,omitempty" validate:"omitempty"`
	BusinessContactCity               string `json:"businessContactCity,omitempty" validate:"omitempty"`
	BusinessContactStateProvince      string `json:"businessContactStateProvince,omitempty" validate:"omitempty"`
	BusinessContactCountry            string `json:"businessContactCountry,omitempty" validate:"omitempty"`
	BusinessContactPostalCode         string `json:"businessContactPostalCode,omitempty" validate:"omitempty"`
}

func (c *CreateUserData) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		return err
	}

	var errorTexts []string

	m, err := regexp.MatchString(`^[\w+,-./~|]{1,75}$`, c.ClientUserID)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for ClientUserID")
	}

	m, err = regexp.MatchString(`^INDIVIDUAL|BUSINESS$`, c.ProfileType)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for ProfileType")
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
			errorTexts = append(errorTexts, "Bad value for MiddleName")
		}
	}

	m, err = regexp.MatchString(`^[\w',\-. ]{1,50}$`, c.LastName)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for LastName")
	}

	m, err = regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, c.DateOfBirth)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for DateOfBirth")
	}
	dob, err := time.Parse("2006-01-02", c.DateOfBirth)
	if err != nil {
		return err
	}
	if age(dob, time.Now()) < 18 {
		errorTexts = append(errorTexts, "Bad value for DateOfBirth")
	}

	if c.GovernmentIDType != "" {
		m, err = regexp.MatchString(`^PASSPORT|NATIONAL_ID_CARD$`, c.GovernmentIDType)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for GovernmentIDType")
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

	m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.StateProvince)
	if err != nil {
		return err
	}
	if m == false || (c.Country == "US" && len(c.StateProvince) != 2) {
		errorTexts = append(errorTexts, "Bad value for StateProvince")
	}

	m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.Country)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for Country")
	}

	m, err = regexp.MatchString(`^[\w \-]{1,16}$`, c.PostalCode)
	if err != nil {
		return err
	}
	if m == false {
		errorTexts = append(errorTexts, "Bad value for PostalCode")
	}

	if c.BusinessName != "" {
		m, err = regexp.MatchString(`^[\w !&'()+,\-./:;]{1,100}$`, c.BusinessName)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessName")
		}
	}

	if c.BusinessOperatingName != "" {
		m, err = regexp.MatchString(`^[\w !&'()+,\-./:;]{1,100}$`, c.BusinessOperatingName)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessOperatingName")
		}
	}

	if c.BusinessRegistrationID != "" {
		m, err = regexp.MatchString(`^[\w ()+\-./]{1,50}$`, c.BusinessRegistrationID)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessRegistrationID")
		}
	}

	if c.BusinessRegistrationStateProvince != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessRegistrationStateProvince)
		if err != nil {
			return err
		}
		if m == false || (c.Country == "US" && len(c.BusinessRegistrationStateProvince) != 2) {
			errorTexts = append(errorTexts, "Bad value for BusinessRegistrationStateProvince")
		}
	}

	if c.BusinessRegistrationCountry != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessRegistrationCountry)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessRegistrationCountry")
		}
	}

	if c.BusinessContactRole != "" {
		m, err = regexp.MatchString(`^DIRECTOR|OWNER|OTHER$`, c.BusinessContactRole)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactRole")
		}
	}

	if c.BusinessContactAddressLine1 != "" {
		m, err = regexp.MatchString(`^[\w #'(),\-./:;°]{1,100}$`, c.BusinessContactAddressLine1)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactAddressLine1")
		}
	}

	if c.BusinessContactAddressLine2 != "" {
		m, err = regexp.MatchString(`^[\w #'(),\-./:;°]{1,100}$`, c.BusinessContactAddressLine2)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactAddressLine2")
		}
	}

	if c.BusinessContactCity != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessContactCity)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactCity")
		}
	}

	if c.BusinessContactStateProvince != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessContactStateProvince)
		if err != nil {
			return err
		}
		if m == false || (c.Country == "US" && len(c.BusinessContactStateProvince) != 2) {
			errorTexts = append(errorTexts, "Bad value for BusinessContactStateProvince")
		}
	}

	if c.BusinessContactCountry != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessContactCountry)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactCountry")
		}
	}

	if c.BusinessContactPostalCode != "" {
		m, err = regexp.MatchString(`^[\w \-]{1,16}$`, c.BusinessContactPostalCode)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactPostalCode")
		}
	}

	if len(errorTexts) > 0 {
		return errors.New(strings.Join(errorTexts, "\n"))
	}

	return nil
}

type UpdateUserData struct {
	ProgramToken                      string `json:"programToken,omitempty"  validate:"omitempty"`
	ClientUserID                      string `json:"clientUserId,omitempty"  validate:"omitempty"`
	ProfileType                       string `json:"profileType,omitempty" validate:"omitempty,alpha,uppercase"`
	FirstName                         string `json:"firstName,omitempty" validate:"omitempty"`
	MiddleName                        string `json:"middleName,omitempty" validate:"omitempty"`
	LastName                          string `json:"lastName,omitempty" validate:"omitempty"`
	DateOfBirth                       string `json:"dateOfBirth,omitempty" validate:"omitempty"`
	CountryOfBirth                    string `json:"countryOfBirth,omitempty" validate:"omitempty"`
	CountryOfNationality              string `json:"countryOfNationality,omitempty" validate:"omitempty"`
	PhoneNumber                       string `json:"phoneNumber,omitempty" validate:"omitempty"`
	MobileNumber                      string `json:"mobileNumber,omitempty" validate:"omitempty"`
	Email                             string `json:"email,omitempty" validate:"omitempty,email"`
	GovernmentID                      string `json:"governmentId,omitempty" validate:"omitempty"`
	GovernmentIDType                  string `json:"governmentIdType,omitempty" validate:"omitempty"`
	PassportID                        string `json:"passportId,omitempty" validate:"omitempty"`
	DriversLicenseID                  string `json:"driversLicenseId,omitempty" validate:"omitempty"`
	EmployerID                        string `json:"employerId,omitempty" validate:"omitempty"`
	AddressLine1                      string `json:"addressLine1,omitempty" validate:"omitempty"`
	AddressLine2                      string `json:"addressLine2,omitempty" validate:"omitempty"`
	City                              string `json:"city,omitempty" validate:"omitempty"`
	StateProvince                     string `json:"stateProvince,omitempty" validate:"omitempty"`
	Country                           string `json:"country,omitempty" validate:"omitempty"`
	PostalCode                        string `json:"postalCode,omitempty" validate:"omitempty"`
	Language                          string `json:"language,omitempty" validate:"omitempty"`
	BusinessType                      string `json:"businessType,omitempty" validate:"omitempty"`
	BusinessName                      string `json:"businessName,omitempty" validate:"omitempty"`
	BusinessOperatingName             string `json:"businessOperatingName,omitempty" validate:"omitempty"`
	BusinessRegistrationID            string `json:"businessRegistrationId,omitempty" validate:"omitempty"`
	BusinessRegistrationStateProvince string `json:"businessRegistrationStateProvince,omitempty" validate:"omitempty"`
	BusinessRegistrationCountry       string `json:"businessRegistrationCountry,omitempty" validate:"omitempty"`
	BusinessContactRole               string `json:"businessContactRole,omitempty" validate:"omitempty"`
	BusinessContactAddressLine1       string `json:"businessContactAddressLine1,omitempty" validate:"omitempty"`
	BusinessContactAddressLine2       string `json:"businessContactAddressLine2,omitempty" validate:"omitempty"`
	BusinessContactCity               string `json:"businessContactCity,omitempty" validate:"omitempty"`
	BusinessContactStateProvince      string `json:"businessContactStateProvince,omitempty" validate:"omitempty"`
	BusinessContactCountry            string `json:"businessContactCountry,omitempty" validate:"omitempty"`
	BusinessContactPostalCode         string `json:"businessContactPostalCode,omitempty" validate:"omitempty"`
}

func (c *UpdateUserData) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		return err
	}

	var m bool
	var errorTexts []string

	if c.ClientUserID != "" {
		m, err := regexp.MatchString(`^[\w+,-./~|]{1,75}$`, c.ClientUserID)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for ClientUserID")
		}
	}

	if c.ProfileType != "" {
		m, err = regexp.MatchString(`^INDIVIDUAL|BUSINESS$`, c.ProfileType)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for ProfileType")
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
			errorTexts = append(errorTexts, "Bad value for MiddleName")
		}

	}

	if c.LastName != "" {
		m, err = regexp.MatchString(`^[\w',\-. ]{1,50}$`, c.LastName)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for LastName")
		}
	}

	if c.DateOfBirth != "" {
		m, err = regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, c.DateOfBirth)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for DateOfBirth")
		}
		dob, err := time.Parse("2006-01-02", c.DateOfBirth)
		if err != nil {
			return err
		}
		if age(dob, time.Now()) < 18 {
			errorTexts = append(errorTexts, "Bad value for DateOfBirth")
		}
	}

	if c.GovernmentIDType != "" {
		m, err = regexp.MatchString(`^PASSPORT|NATIONAL_ID_CARD$`, c.GovernmentIDType)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for GovernmentIDType")
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

	if c.StateProvince != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.StateProvince)
		if err != nil {
			return err
		}
		if m == false || (c.Country == "US" && len(c.StateProvince) != 2) {
			errorTexts = append(errorTexts, "Bad value for StateProvince")
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

	if c.PostalCode != "" {
		m, err = regexp.MatchString(`^[\w \-]{1,16}$`, c.PostalCode)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for PostalCode")
		}
	}

	if c.BusinessName != "" {
		m, err = regexp.MatchString(`^[\w !&'()+,\-./:;]{1,100}$`, c.BusinessName)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessName")
		}
	}

	if c.BusinessOperatingName != "" {
		m, err = regexp.MatchString(`^[\w !&'()+,\-./:;]{1,100}$`, c.BusinessOperatingName)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessOperatingName")
		}
	}

	if c.BusinessRegistrationID != "" {
		m, err = regexp.MatchString(`^[\w ()+\-./]{1,50}$`, c.BusinessRegistrationID)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessRegistrationID")
		}
	}

	if c.BusinessRegistrationStateProvince != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessRegistrationStateProvince)
		if err != nil {
			return err
		}
		if m == false || (c.Country == "US" && len(c.BusinessRegistrationStateProvince) != 2) {
			errorTexts = append(errorTexts, "Bad value for BusinessRegistrationStateProvince")
		}
	}

	if c.BusinessRegistrationCountry != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessRegistrationCountry)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessRegistrationCountry")
		}
	}

	if c.BusinessContactRole != "" {
		m, err = regexp.MatchString(`^DIRECTOR|OWNER|OTHER$`, c.BusinessContactRole)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactRole")
		}
	}

	if c.BusinessContactAddressLine1 != "" {
		m, err = regexp.MatchString(`^[\w #'(),\-./:;°]{1,100}$`, c.BusinessContactAddressLine1)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactAddressLine1")
		}
	}

	if c.BusinessContactAddressLine2 != "" {
		m, err = regexp.MatchString(`^[\w #'(),\-./:;°]{1,100}$`, c.BusinessContactAddressLine2)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactAddressLine2")
		}
	}

	if c.BusinessContactCity != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessContactCity)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactCity")
		}
	}

	if c.BusinessContactStateProvince != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessContactStateProvince)
		if err != nil {
			return err
		}
		if m == false || (c.Country == "US" && len(c.BusinessContactStateProvince) != 2) {
			errorTexts = append(errorTexts, "Bad value for BusinessContactStateProvince")
		}
	}

	if c.BusinessContactCountry != "" {
		m, err = regexp.MatchString(`^[\w &'()\-.]{1,50}$`, c.BusinessContactCountry)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactCountry")
		}
	}

	if c.BusinessContactPostalCode != "" {
		m, err = regexp.MatchString(`^[\w \-]{1,16}$`, c.BusinessContactPostalCode)
		if err != nil {
			return err
		}
		if m == false {
			errorTexts = append(errorTexts, "Bad value for BusinessContactPostalCode")
		}
	}

	if len(errorTexts) > 0 {
		return errors.New(strings.Join(errorTexts, "\n"))
	}

	return nil
}

func age(birthdate, today time.Time) int {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}

type User struct {
	Token              string `json:"token"`
	Status             string `json:"status"`
	VerificationStatus string `json:"verificationStatus"`
	CreatedOn          string `json:"createdOn"`
	ClientUserID       string `json:"clientUserId"`
	ProfileType        string `json:"profileType"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	DateOfBirth        string `json:"dateOfBirth"`
	Email              string `json:"email"`
	AddressLine1       string `json:"addressLine1"`
	City               string `json:"city"`
	StateProvince      string `json:"stateProvince"`
	Country            string `json:"country"`
	PostalCode         string `json:"postalCode"`
	Language           string `json:"language"`
	TimeZone           string `json:"timeZone"`
	ProgramToken       string `json:"programToken"`
	Links              []Link `json:"links"`
}

type Link struct {
	Params Params `json:"params"`
	Href   string `json:"href"`
}

type Params struct {
	Rel string `json:"rel"`
}

type UserList struct {
	Count  int    `json:"count"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Data   []User `json:"data"`
}

type AuthenticationToken struct {
	Value string `json:"value"`
}

type UserBalance struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type UserBalanceList struct {
	Count  int           `json:"count"`
	Offset int           `json:"offset"`
	Limit  int           `json:"limit"`
	Data   []UserBalance `json:"data"`
	Links  []Link        `json:"links"`
}

type UserReceipt struct {
	Token            string `json:"token"`
	JournalID        string `json:"journalId"`
	Type             string `json:"type"`
	CreatedOn        string `json:"createdOn"`
	Entry            string `json:"entry"`
	SourceToken      string `json:"sourceToken"`
	DestinationToken string `json:"destinationToken,omitempty"`
	Amount           string `json:"amount"`
	Fee              string `json:"fee,omitempty"`
	Currency         string `json:"currency"`
	Details          struct {
		ClientPaymentID string `json:"clientPaymentId"`
		PayeeName       string `json:"payeeName"`
	} `json:"details,omitempty"`
}

type UserReceiptList struct {
	Count  int           `json:"count"`
	Offset int           `json:"offset"`
	Limit  int           `json:"limit"`
	Data   []UserReceipt `json:"data"`
	Links  []Link        `json:"links"`
}

type GetUserListQuery struct {
	ClientUserID       string    `url:"clientUserId,omitempty"`
	CreatedBefore      time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T15:04:05"`
	CreatedAfter       time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T15:04:05"`
	Email              string    `url:"email,omitempty" validate:"email"`
	ProgramToken       string    `url:"programToken,omitempty"`
	Status             string    `url:"status,omitempty"`
	VerificationStatus string    `url:"verificationStatus,omitempty"`
	SortBy             string    `url:"sortBy,omitempty"`
	Offset             string    `url:"offset,omitempty"`
	Limit              string    `url:"limit,omitempty"`
}

type GetUserBalanceListQuery struct {
	Currency string `url:"currency,omitempty"`
	SortBy   string `url:"sortBy,omitempty"`
	Offset   string `url:"offset,omitempty"`
	Limit    string `url:"limit,omitempty"`
}

type GetUserReceiptListQuery struct {
	CreatedBefore time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T15:04:05"`
	CreatedAfter  time.Time `url:"createdBefore,omitempty" layout:"2006-01-02T15:04:05"`
	Currency      string    `url:"currency,omitempty"`
	SortBy        string    `url:"sortBy,omitempty"`
	Offset        string    `url:"offset,omitempty"`
	Limit         string    `url:"limit,omitempty"`
}
