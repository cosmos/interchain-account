package types

// NewQueryIBCAccountFromAddressResponse creates a new QueryIBCAccountFromAddressResponse instance
func NewQueryIBCAccountFromAddressResponse(addr string) *QueryIBCAccountFromAddressResponse {
	return &QueryIBCAccountFromAddressResponse{
		Address: addr,
	}
}
