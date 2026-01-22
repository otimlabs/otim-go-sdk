package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/otimlabs/otim-go-sdk/signer"

	"github.com/ethereum/go-ethereum/common"
)

type DelegateAddressResponse struct {
	Address string `json:"otimDelegateAddress"`
}

type Client struct {
	signer.Signer
	apiKey string
	apiUrl string
	// Otim delegate address, automatically fetched from API during initialization
	otimDelegateAddr common.Address
}

func NewClient(signer signer.Signer, apiUrl string, apiKey string) (*Client, error) {
	apiUrl = strings.TrimRight(apiUrl, "/")

	c := &Client{
		Signer: signer,
		apiKey: apiKey,
		apiUrl: apiUrl,
	}

	// Fetch delegate address for mainnet (chain ID 1)
	var resp DelegateAddressResponse
	if err := c.getJSON(context.Background(), "/config/delegate/address/1", &resp); err != nil {
		return nil, fmt.Errorf("fetch delegate address: %w", err)
	}

	if !common.IsHexAddress(resp.Address) {
		return nil, fmt.Errorf("invalid delegate address from API: %s", resp.Address)
	}

	c.otimDelegateAddr = common.HexToAddress(resp.Address)

	return c, nil
}

func (c *Client) doRequest(
	ctx context.Context,
	method string,
	path string,
	body any,
	out any,
) error {
	url := c.apiUrl + path

	var reqBody io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
		reqBody = bytes.NewReader(encoded)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(b))
	}

	if out == nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}

func (c *Client) getJSON(
	ctx context.Context,
	path string,
	out any,
) error {
	return c.doRequest(ctx, http.MethodGet, path, nil, out)
}

func (c *Client) postJSON(
	ctx context.Context,
	path string,
	body any,
	out any,
) error {
	return c.doRequest(ctx, http.MethodPost, path, body, out)
}
