package hyperwallet

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	defaultTimeout = time.Second * 60
	BaseApiUrlV3   = "https://api.sandbox.hyperwallet.com/rest/v3"
)

type Client interface {
	Execute(ctx context.Context, method string, path string, urlQuery url.Values, body string) ([]byte, error)
}

type Hyperwallet struct {
	BaseApiUrl   string
	ProgramToken string
	UserName     string
	Password     string
	HttpClient   *http.Client
}

func NewClient() *Hyperwallet {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}

	return &Hyperwallet{
		BaseApiUrl:   BaseApiUrlV3,
		ProgramToken: os.Getenv("PROGRAM_TOKEN"),
		UserName:     os.Getenv("API_USER_NAME"),
		Password:     os.Getenv("API_PASSWORD"),
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
	request.SetBasicAuth(h.UserName, h.Password)

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
