package types

// NewQueryInterchainAccountRequest creates and returns a new QueryInterchainAccountFromAddressRequest
func NewQueryInterchainAccountRequest(owner string) *QueryInterchainAccountFromAddressRequest {
	return &QueryInterchainAccountFromAddressRequest{
		Owner: owner,
	}
}

// NewQueryInterchainAccountResponse creates and returns a new QueryInterchainAccountFromAddressResponse
func NewQueryInterchainAccountResponse(interchainAccAddr string) *QueryInterchainAccountFromAddressResponse {
	return &QueryInterchainAccountFromAddressResponse{
		InterchainAccountAddress: interchainAccAddr,
	}
}
