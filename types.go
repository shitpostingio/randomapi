package main

import "time"

type netConfigs struct {
	endpoint        string
	storageEndpoint string
}

type requestedPost struct {
	path        string
	mediatype   string
	requestdate time.Time
}
