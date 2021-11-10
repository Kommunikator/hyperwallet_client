package hyperwallet

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultTimeout = time.Second * 60
	// todo Данные для доступа потом надо брать из переменных окружения
	BaseApiUrlV3   = "https://api.sandbox.hyperwallet.com/rest/v3"
	ProgramToken = ""
	ApiUserName = ""
	ApiPassword = ""
)

type Hyperwallet struct {
	baseApiUrl   string
	programToken string
	userName     string
	password     string
	HttpClient   *http.Client
}

func NewClient() *Hyperwallet {
	return &Hyperwallet{
		baseApiUrl:   BaseApiUrlV3,
		programToken: ProgramToken,
		userName:     ApiUserName,
		password:     ApiPassword,
		HttpClient:   &http.Client{Timeout: defaultTimeout},
	}
}

func (h *Hyperwallet) Execute(ctx context.Context, method string, path string, urlQuery url.Values, body string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", BaseApiUrlV3, path), strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = urlQuery.Encode()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.SetBasicAuth(h.userName, h.password)

	response, err := h.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
