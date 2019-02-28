# tftp-proxy
A TFTP server that proxies request to an HTTP backend if a file is not found.

# How to build
    go build

# How to run
    ./tftp-proxy -url=http://example.com -dir=/var/lib/tftpboot &
