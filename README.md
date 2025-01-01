# randomNetIP

`randomNetIP` is a Go package that provides a function to generate random, publicly accessible IPv4 addresses. It ensures the generated IP addresses do not belong to private or reserved IP ranges, offering a reliable method for generating valid public IPs.

## Installation

To install the package, use the following `go get` command:

```bash
go get github.com/CNMengHan/randomNetIP
```

## Usage

The `randomNetIP` package exposes a single function: `RandomNetIPv4()`. This function generates a random IPv4 address, ensuring it is a public IP (not within the private IP ranges).

### Example Usage

```go
package main

import (
	"fmt"
	"github.com/CNMengHan/randomNetIP"
)

func main() {
	// Generate a random public IPv4 address
	ip := randomNetIP.RandomNetIPv4()
	fmt.Println("Generated random public IPv4 address:", ip)
}
```

### Function Details

#### `RandomNetIPv4() string`

This function generates a random public IPv4 address by following these steps:

1. It randomly selects four octets for the IP address.
2. It checks whether the generated IP address falls within any known private IP ranges.
3. If the generated IP falls within a private range, the function re-generates the IP until it finds a valid public IP address.

### Example Output

```
Generated random public IPv4 address: 198.45.132.56
```

## Private IP Range Exclusions

The generated IP addresses are ensured not to fall within any of the following private or reserved ranges:

- `10.0.0.0/8` – Private network (Class A)
- `172.16.0.0/12` – Private network (Class B)
- `192.168.0.0/16` – Private network (Class C)
- `169.254.0.0/16` – Link-local address (used for automatic configuration)
- `100.64.0.0/10` – Shared address space (used for carrier-grade NAT)
- `127.0.0.0/8` – Loopback address (localhost)

The first octet of the generated IP is randomly selected between `1` and `223`, ensuring the address is not within any reserved blocks like `0.x.x.x` or `224.x.x.x` and higher, which are used for multicast.

## Dependencies

This package has no external dependencies and only relies on Go's standard libraries:

- `math/rand` – For generating random numbers.
- `net` – For working with IP addresses and CIDR ranges.

## Contributing

Contributions are welcome! Feel free to fork the repository and submit a pull request with improvements, bug fixes, or new features.

If you encounter any issues or have suggestions, please create an issue on GitHub.

