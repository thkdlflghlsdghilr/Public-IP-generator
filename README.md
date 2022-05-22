# Public-IP-generator
Generates all possible public IPs and writes them into a file or prints them in the stdout written in go
## Usage
```
go build
./genIPs [flags]
```
### Flags
`-w={filename}`  Outputs to filename provided, If no name is given it won't write to a file. `Default: ""`
`-o` Outputs genarated IPs to stdout separated with newline. `Default: False`
`-t` Prints time it took to generate IPs once finished. `Default: True`

#### Credits
- https://en.wikipedia.org/wiki/Reserved_IP_addresses
