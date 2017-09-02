package log

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"git.vie.cybertrap.com/ppacher/dnslog/request"
	"github.com/miekg/dns"
)

type LogMiddleware struct{}

func (*LogMiddleware) Name() string {
	return "log"
}

func (*LogMiddleware) Serve(ctx context.Context, req *request.Request) (context.Context, *dns.Msg, error) {
	return ctx, nil, nil
}

func (*LogMiddleware) Mangle(ctx context.Context, req *request.Request, response *dns.Msg) error {
	answer := "Not Resolved"
	if response != nil && len(response.Answer) > 0 {
		answer = fmt.Sprintf("%s", response.Answer[0].String())
	}

	log.Printf("[log] %s requested %q class=%s type=%s, resolved to: %s", req.RemoteAddr().String(), req.Name(), req.Class(), req.Type(), answer)

	return nil
}
