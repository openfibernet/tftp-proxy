# tftp-proxy
A TFTP server that proxies request to an HTTP backend if a file is not found.

# Building

    go build

or for linux:

    env GOOS=linux GOARCH=amd64 go build

# How to run
    ./tftp-proxy -url=http://example.com -dir=/var/lib/tftpboot &

or for linux:

	scp tftp-proxy.service tftp-proxy destination-host:
	ssh destination-host
	sudo mv tftp-proxy.service /etc/systemd/system/
	sudo mv tftp-proxy /usr/bin/tftp-proxy
	sudo vi /etc/systemd/system/tftp-proxy.service
	systemctl enable tftp-proxy