// Copyright 2016 SMFS Inc DBA GRIMM. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"gocsp-responder/pkg/responder"
)

func main() {
	resp := responder.Responder()
	flag.StringVar(&resp.IndexFile, "index", resp.IndexFile, "CA index filename")
	flag.StringVar(&resp.CaCertFile, "cacert", resp.CaCertFile, "CA certificate filename")
	flag.StringVar(&resp.RespCertFile, "rcert", resp.RespCertFile, "responder certificate filename")
	flag.StringVar(&resp.RespKeyFile, "rkey", resp.RespKeyFile, "responder key filename")
	flag.StringVar(&resp.LogFile, "logfile", resp.LogFile, "file to log to")
	flag.StringVar(&resp.Address, "bind", resp.Address, "bind address")
	flag.IntVar(&resp.Port, "port", resp.Port, "listening port")
	flag.BoolVar(&resp.Ssl, "ssl", resp.Ssl, "use SSL, this is not widely supported and not recommended")
	flag.BoolVar(&resp.Strict, "strict", resp.Strict, "require content type HTTP header")
	flag.BoolVar(&resp.LogToStdout, "stdout", resp.LogToStdout, "log to stdout, not the log file")
	flag.Parse()
	resp.Serve()
}
