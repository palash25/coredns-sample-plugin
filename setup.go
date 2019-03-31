package demo

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"

	"log"

	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin("demo", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	fw, _ := parse(c)

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		fw.Next = next
		return fw
	})

	return nil
}

func parse(c *caddy.Controller) (*Demo, error) {
	log.Printf("Controller struct %v", c)
	f, _ := New()
	for c.Next() { // Skip the directive name.
		log.Printf("enter for %v", c.Next())
		if !c.NextArg() { // Expect at least one value.
			log.Printf("enter if %v", c.NextArg())
			return f, c.ArgErr() // Otherwise it's an error.
		}
		value := c.Val()
		f.ip = value
		log.Printf("ip is %v", f.ip)
	}
	return f, nil
}

/*
.:1053 {
    demo source {
        allow "127.0.0.1"
        block "192.168.1.11"
    }
}
*/
