//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Build() error {
	mg.Deps(LoadDeps)
	return sh.Run("go", "build", "-o", "./bin/tasks", "./cmd/tasks.go")
}

func LoadDeps() error {
  return sh.Run("go", "mod", "tidy")
}
