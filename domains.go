package linodego

import (
	"fmt"
	"time"

	"github.com/go-resty/resty"
)

// Domain represents a Domain object
type Domain struct {
	ID int
	// UpdatedStr string `json:"updated"`

	Updated *time.Time `json:"-"`
}

// DomainsPagedResponse represents a paginated Domain API response
type DomainsPagedResponse struct {
	*PageOptions
	Data []*Domain
}

// endpoint gets the endpoint URL for Domain
func (DomainsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.Domains.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends Domains when processing paginated Domain responses
func (resp *DomainsPagedResponse) appendData(r *DomainsPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of Domain
func (DomainsPagedResponse) setResult(r *resty.Request) {
	r.SetResult(DomainsPagedResponse{})
}

// ListDomains lists Domains
func (c *Client) ListDomains(opts *ListOptions) ([]*Domain, error) {
	response := DomainsPagedResponse{}
	err := c.listHelper(&response, opts)
	for _, el := range response.Data {
		el.fixDates()
	}
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *Domain) fixDates() *Domain {
	// v.Created, _ = parseDates(v.CreatedStr)
	// v.Updated, _ = parseDates(v.UpdatedStr)
	return v
}

// GetDomain gets the template with the provided ID
func (c *Client) GetDomain(id string) (*Domain, error) {
	e, err := c.Domains.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, id)
	r, err := c.R().SetResult(&Domain{}).Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*Domain).fixDates(), nil
}