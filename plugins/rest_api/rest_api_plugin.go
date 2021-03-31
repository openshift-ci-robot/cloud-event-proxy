package main

import (
	"sync"

	restapi "github.com/redhat-cne/rest-api"
	"github.com/redhat-cne/sdk-go/pkg/channel"
)

//var server *restapi.Server

//Start start rest service to creat delete , update and handle subscriptions,publishers, events and status
func Start(wg *sync.WaitGroup, port int, apiPath, storePath string, eventOutCh chan<- *channel.DataChan, close <-chan bool) *restapi.Server { //nolint:deadcode,unused
	server := restapi.InitServer(port, apiPath, storePath, eventOutCh, close)
	defer wg.Done()
	wg.Add(1)
	go server.Start(wg)
	return server
}