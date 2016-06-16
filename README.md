# `docker build` utility for [docker/engine-api](https://github.com/docker/engine-api)

[![GoDoc](https://godoc.org/github.com/AkihiroSuda/go-docker-builder?status.svg)](https://godoc.org/github.com/AkihiroSuda/go-docker-builder)
[![Build Status](https://travis-ci.org/AkihiroSuda/go-docker-builder.svg?branch=master)](https://travis-ci.org/AkihiroSuda/go-docker-builder)
[![Go Report Card](https://goreportcard.com/badge/github.com/AkihiroSuda/go-docker-builder)](https://goreportcard.com/report/github.com/AkihiroSuda/go-docker-builder)

go-docker-builder provides a tar stream for ['docker/engine-api.(*Client).ImageBuild'](https://godoc.org/github.com/docker/engine-api/client#Client.ImageBuild) so as to build an image from a local directory.

For usage, please refer to [`builder_test.go`](builder_test.go).
