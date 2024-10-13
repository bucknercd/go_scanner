# go_scanner

This is a ip/port scanner that will scan an ip range, or single ip, to see if its alive and if it is, see what ports , if any are open in the first 10k.
It uses go routines and mutex locks.

Example usage:  ./scan -t 192.168.1.1-192.168.1.254 or ./scan -t 192.168.0.110

If you wish to recompile scan.go you will need to set GOPATH to $(pwd)/lib

test

Example: export GOPATH=$(pwd)/lib
