package main

import (
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestThree(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

// go clean -testcache
