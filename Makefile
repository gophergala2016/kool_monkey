PROJECT = $(shell pwd)

kool-server:
	GOPATH=${PROJECT} go install kool-server
