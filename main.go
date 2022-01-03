package main

import (
	"github.com/valyala/fasthttp"
	proxy "github.com/yeqown/fasthttp-reverse-proxy/v2"
	"log"
)

func ProxyHandler(ctx *fasthttp.RequestCtx) {
	myProxy.ServeHTTP(ctx)
	ctx.Response.Header.Add("myname", "jerry")
}

var myProxy = proxy.NewReverseProxy("www.jtthink.com")

func main() {
	err := fasthttp.ListenAndServe(":80", ProxyHandler)
	if err != nil {
		log.Fatal(err)
	}
}
