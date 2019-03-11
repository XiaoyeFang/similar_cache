// Copyright © 2018 joy  <lzy@spf13.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	"flag"
	"similar_cache/cachegrpc"
	"similar_cache/config"
	"similar_cache/protos"
	"sync"
)

func main() {
	flag.Parse()
	glog.Flush()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		grpcStart()
	}()
	wg.Wait()

}

func grpcStart() {
	glog.V(0).Infoln(config.CacheConfig.GrpcListen)
	lis, err := net.Listen("tcp", config.CacheConfig.GrpcListen)
	if err != nil {
		panic(err)
	}
	glog.Errorf("am-grpc-port %v", config.CacheConfig.GrpcListen)
	server := grpc.NewServer()
	srv := cachegrpc.FragmentServer{}
	protos.RegisterFragMentServiceServer(server, &srv)
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
