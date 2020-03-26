/*
Go PlayGround:
A web service running on go lang setvers at google.package tour

Three parts:
Iamge: https://blog.golang.org/playground
1. Back end:
	runs on Google’s servers,
	receives RPC requests, compiles the code, executes
	sends result/ compilation errors back.
2. Front end:
	Google App engine
	receives HTTP requests from the client
	sends to back end,
	caches results also
3. Client:
	JS app, HTTP calls and UI.

Features:
1. Time is static: 2009-11-10 23:00:00 UTC
2. Limited CPU and memory usage
3. No other hosts connections allowed
*/

package main

import (
	"fmt"
	"time"
)

func manyImports(){
	fmt.Println("Inside many imoprts")
	fmt.Println("Time is", time.Now())
}
