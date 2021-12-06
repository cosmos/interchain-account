 #!/bin/bash
HERMES_BINARY=hermes
HERMES_DIRECTORY=./network/hermes/
CONFIG_DIR=./network/hermes/config.toml

echo "Using hermes relayer version: "
$HERMES_BINARY --version | sed 's/^/    /'
