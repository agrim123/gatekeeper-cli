package main

import (
	"context"
	"fmt"

	"github.com/agrim123/gatekeeper"
)

func main() {
	fmt.Println(gatekeeper.NewGatekeeper(context.Background()))
}
