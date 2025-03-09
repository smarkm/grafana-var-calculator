//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Build() error {
	mg.Deps(BuildBackend)
	return nil
}

func BuildBackend() error {
	return sh.RunV("go", "build", "-o", "./dist/calculator_linux_amd64", "./pkg")
}
