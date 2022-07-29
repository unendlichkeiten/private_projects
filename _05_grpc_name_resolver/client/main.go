package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	pb "github.com/unendlichkeiten/private_projects/pb"
)

const (
	myScheme      = "custom"
	myServiceName = "resolver.custom.hamming.com"

	backendAddr = "localhost:50051"
)

func callUnaryEcho(c pb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.Message)
}

func makeRPCs(cc *grpc.ClientConn, n int) {
	hwc := pb.NewEchoClient(cc)
	for i := 0; i < n; i++ {
		callUnaryEcho(hwc, "this is examples/name_resolving")
	}
}

func main() {
	passthroughConn, err := grpc.Dial(
		// passthrough 是 gRPC 内置的一个 scheme
		// Dial to "passthrough:///localhost:50051"
		fmt.Sprintf("passthrough:///%s", backendAddr),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer passthroughConn.Close()

	fmt.Printf("calling SayHello to \"passthrough:///%s\"\n", backendAddr)
	makeRPCs(passthroughConn, 10)

	fmt.Println()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	exampleConn, err := grpc.DialContext(
		ctx,
		// 使用自定义的名字解析器
		// Dial to "custom:///resolver.custom.hamming.com"
		fmt.Sprintf("%s:///%s", myScheme, myServiceName),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer exampleConn.Close()

	fmt.Printf("calling SayHello to \"%s:///%s\"\n", myScheme, myServiceName)
	makeRPCs(exampleConn, 10)
}

// resolver 的实现
// Following is an example name resolver. It includes a
// ResolverBuilder(https://godoc.org/google.golang.org/grpc/resolver#Builder)
// and a Resolver(https://godoc.org/google.golang.org/grpc/resolver#Resolver).
//
// A ResolverBuilder is registered for a scheme (in this example, "example" is
// the scheme). When a ClientConn is created for this scheme, the
// ResolverBuilder will be picked to build a Resolver. Note that a new Resolver
// is built for each ClientConn. The Resolver will watch the updates for the
// target, and send updates to the ClientConn.

// customResolverBuilder is a
// ResolverBuilder(https://godoc.org/google.golang.org/grpc/resolver#Builder).
type customResolverBuilder struct{}

// Build 构建解析器
func (*customResolverBuilder) Build(
	target resolver.Target,
	cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {

	r := &customResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			myServiceName: {backendAddr},
		},
	}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}

// Scheme 返回 customResolverBuilder 对应的 scheme
func (*customResolverBuilder) Scheme() string { return myScheme }

// customResolver is a
// Resolver(https://godoc.org/google.golang.org/grpc/resolver#Resolver).
type customResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *customResolver) ResolveNow(o resolver.ResolveNowOptions) {
	// 直接从map中取出对于的addrList
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (*customResolver) Close() {}

func init() {
	// Register the example ResolverBuilder. This is usually done in a package's
	// init() function.
	resolver.Register(&customResolverBuilder{})
}
