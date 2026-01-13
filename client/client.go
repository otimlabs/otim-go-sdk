package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"otim-go-sdk/signer"

	"github.com/ethereum/go-ethereum/common"
)

type Client struct {
	signer.Signer
	apiKey string
	apiUrl string
	// Otim delegate address, should instead be a method parameter if we want it flexible
	otimDelegateAddr common.Address
}

func NewClient(signer signer.Signer, apiUrl string, apiKey string, otimDelegateAddr common.Address) *Client {
	apiUrl = strings.TrimRight(apiUrl, "/")
	return &Client{
		Signer:           signer,
		apiKey:           apiKey,
		apiUrl:           apiUrl,
		otimDelegateAddr: otimDelegateAddr,
	}
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
