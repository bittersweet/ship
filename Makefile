all:
	@@go build ship.go
	@@mv ship /usr/local/bin/
	@@echo "built and moved"
