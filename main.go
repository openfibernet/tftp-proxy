package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/pin/tftp"
)

var url string
var dir string

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	file_path := filepath.Clean(path.Join(dir, filename))

	if _, err := os.Stat(file_path); err == nil {
		file, err := os.Open(file_path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}

		fi, err := file.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}

		rf.(tftp.OutgoingTransfer).SetSize(fi.Size())
		n, err := rf.ReadFrom(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}

		fmt.Printf("%s %d bytes sent\n", filename, n)
	} else { // File not found locally. Proxying the request.
		fileUrl := url + "/" + filename
		resp, err := http.Get(fileUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			io.Copy(io.Discard, resp.Body)
			return fmt.Errorf("received status code: %d", resp.StatusCode)
		}

		rf.(tftp.OutgoingTransfer).SetSize(resp.ContentLength)
		n, err := rf.ReadFrom(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}

		fmt.Printf("%s %d bytes sent\n", filename, n)
	}

	return nil
}

func main() {
	flag.StringVar(&dir, "dir", "/var/lib/tftpboot", "The directory to serve files from. For example /var/lib/tftpboot")
	flag.StringVar(&url, "url", "http://example.com", "The URL to proxy requests to. For example http://example.com")
	flag.Parse()

	// use nil in place of handler to disable read or write operations
	s := tftp.NewServer(readHandler, nil)
	s.SetTimeout(5 * time.Second)  // optional
	err := s.ListenAndServe(":69") // blocks until s.Shutdown() is called
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}
}
