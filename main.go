package main

import (
  "context"
  "fmt"
  "github.com/gerritjvv/tcpshaper/bandwidth"
  "net"
)

func main() { 
  // Get a Listener, e.g tcp on any port
  tcpListner, err := net.Listen("tcp", ":0")
  if err != nil {
   panic(err)
  }

  ctx := context.Background()

  // Configure rate limit to:
  // total read+write traffic == 1mbs
  // Each connection read+write == 2kbs
  // The maximum that can be read or written at any one time is 2kb
  serverRate := bandwidth.NewRateConfig(1024*1024, 2048)
  connRate := bandwidth.NewRateConfig(2048, 2048)

  listener := bandwidth.NewListener(ctx, &bandwidth.ListenerConfig{
    ReadServerRate:  serverRate,
    WriteServerRate: serverRate,
    ReadConnRate:    connRate,
    WriteConnRate:   connRate,
  }, tcpListner)

  // Now use the listener
  // e.g listener.Accept()
  
  err = listener.Close()
  if err != nil {
    fmt.Printf("error while closing listener %s", err)
  }
}
