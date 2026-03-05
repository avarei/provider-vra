/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/v2/pkg/pipeline"

	"github.com/avarei/provider-vra/v2/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}

	pc, err := config.GetProvider(context.Background(), true)
	if err != nil {
		panic(fmt.Sprintf("cannot get cluster provider configuration: %v", err))
	}
	pns, err := config.GetProviderNamespaced(context.Background(), true)
	if err != nil {
		panic(fmt.Sprintf("cannot get namespaced provider configuration: %v", err))
	}

	pipeline.Run(pc, pns, absRootDir)
}
