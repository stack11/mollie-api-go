package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// PaymentLink is a resource that can be shared with your customers
// and will redirect them to them the payment page where they can
// complete the payment.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link
type PaymentLink struct {
	ID          string           `json:"id,omitempty"`
	Resource    string           `json:"resource,omitempty"`
	Description string           `json:"description,omitempty"`
	ProfileID   string           `json:"profileId,omitempty"`
	RedirectURL string           `json:"redirectUrl,omitempty"`
	WebhookURL  string           `json:"webhookUrl,omitempty"`
	Mode        Mode             `json:"mode,omitempty"`
	Amount      Amount           `json:"amount,omitempty"`
	CreatedAt   *time.Time       `json:"createdAt,omitempty"`
	PaidAt      *time.Time       `json:"paidAt,omitempty"`
	UpdatedAt   *time.Time       `json:"updatedAt,omitempty"`
	ExpiresAt   *time.Time       `json:"expiresAt,omitempty"`
	Links       PaymentLinkLinks `json:"_links,omitempty"`
}

// PaymentLinkLinks describes all the possible links returned with
// a payment link struct.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link
type PaymentLinkLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
	PaymentLink   *URL `json:"paymentLink,omitempty"`
	Next          *URL `json:"next,omitempty"`
	Previous      *URL `json:"previous,omitempty"`
}

// PaymentLinkOptions represents query string parameters to modify
// the payment links requests.
type PaymentLinkOptions struct {
	ProfileID string `url:"profileId,omitempty"`
	From      string `url:"from,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

// PaymentLinksList retrieves a list of payment links for the active
// profile or account token owner.
type PaymentLinksList struct {
	Count    int              `json:"count,omitempty"`
	Links    PaymentLinkLinks `json:"_links,omitempty"`
	Embedded struct {
		PaymentLinks []*PaymentLink `json:"payment_links,omitempty"`
	} `json:"_embedded,omitempty"`
}

// PaymentLinksService operates over the payment link resource.
type PaymentLinksService service

// Get retrieves a single payment link object by its id/token.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link
func (pls *PaymentLinksService) Get(ctx context.Context, id string) (res *Response, pl *PaymentLink, err error) {
	res, err = pls.client.get(ctx, fmt.Sprintf("v2/payment-links/%s", id), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}

	return
}

// Create generates payment links that by default, unlike regular payments, do not expire.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/create-payment-link
func (pls *PaymentLinksService) Create(ctx context.Context, p PaymentLink, opts *PaymentLinkOptions) (res *Response, np *PaymentLink, err error) {
	res, err = pls.client.post(ctx, "v2/payment-links", p, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &np); err != nil {
		return
	}

	return
}

// List retrieves all payments links created with the current website profile,
// ordered from newest to oldest.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/list-payment-links
func (pls *PaymentLinksService) List(ctx context.Context, opts *PaymentLinkOptions) (res *Response, pl *PaymentLinksList, err error) {
	res, err = pls.client.get(ctx, "v2/payment-links", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}

	return
}
