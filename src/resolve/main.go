package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	godns "github.com/miekg/dns"
)

func main() {
	x, y := resolve("119.29.29.29", "dns-xxxx.qbox.me", "1.1.1.1")
}
