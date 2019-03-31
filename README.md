# coredns-sample-plugin

This plugin uses source based IP discovery to return various responses. his is a hack on the demo plugin.

- IPs prefixed with `"172."` return `"1.1.1.1"` as a message and `dns.RcodeSuccess` as a status
- IPs prefixed with `"127."` return `"8.8.8.8"` as a message and `dns.RcodeSuccess` as a status
- Rest of the IPs return `"1.1.1.1"` as a message and `dns.RcodeNameError` as a status

## Corefile Syntax

.:1053 {
    demo allow "127.0.0.1"
}