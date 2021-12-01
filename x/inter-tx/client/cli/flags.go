package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagConnectionId             = "connection-id"
	FlagCounterpartyConnectionId = "counterparty-connection-id"
)

// common flagsets to add to various functions
var (
	fsConnectionId = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsConnectionId.String(FlagConnectionId, "", "Connection ID")
	fsConnectionId.String(FlagCounterpartyConnectionId, "", "Counterparty Connection ID")
}
