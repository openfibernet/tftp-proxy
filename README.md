# tftp-proxy
A TFTP server that proxies request to an HTTP backend if a file is not found.

# How to build
    go build main.go

# How to run
    ./tftp -url=http://example.com -dir=/var/lib/tftpboot &
