/*
 * ImageCashLetter API
 *
 * Moov ImageCashLetter () implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CashLetter struct {
	CashLetterHeader CashLetterHeader `json:"cashLetterHeader,omitempty"`
	CreditItems []CreditItem `json:"creditItems,omitempty"`
	Bundles []Bundle `json:"bundles,omitempty"`
	RoutingNumberSummary []RoutingNumberSummary `json:"routingNumberSummary,omitempty"`
	CashLetterControl CashLetterControl `json:"cashLetterControl,omitempty"`
}