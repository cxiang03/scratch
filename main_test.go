package main

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func BenchmarkLoop5NotExist(b *testing.B) {
	data := make([]string, 5)
	for i := 0; i < 5; i++ {
		data = append(data, RandStringRunes(10))
	}
	for i := 0; i < b.N; i++ {
		for _, v := range data {
			if v == "not_exist" {
				break
			}
		}
	}
}

func BenchmarkLoop10NotExist(b *testing.B) {
	data := make([]string, 10)
	for i := 0; i < 10; i++ {
		data = append(data, RandStringRunes(10))
	}
	for i := 0; i < b.N; i++ {
		for _, v := range data {
			if v == "not_exist" {
				break
			}
		}
	}
}

func BenchmarkLoop100NotExist(b *testing.B) {
	data := make([]string, 100)
	for i := 0; i < 100; i++ {
		data = append(data, RandStringRunes(10))
	}
	for i := 0; i < b.N; i++ {
		for _, v := range data {
			if v == "not_exist" {
				break
			}
		}
	}
}

func BenchmarkSet5NotExist(b *testing.B) {
	data := make(map[string]struct{}, 10)
	for i := 0; i < 5; i++ {
		data[RandStringRunes(10)] = struct{}{}
	}
	for i := 0; i < b.N; i++ {
		_, ok := data["not_exist"]
		_ = ok
	}
}

func BenchmarkSet10NotExist(b *testing.B) {
	data := make(map[string]struct{}, 10)
	for i := 0; i < 10; i++ {
		data[RandStringRunes(10)] = struct{}{}
	}
	for i := 0; i < b.N; i++ {
		_, ok := data["not_exist"]
		_ = ok
	}
}

func BenchmarkSet100NotExist(b *testing.B) {
	data := make(map[string]struct{}, 100)
	for i := 0; i < 100; i++ {
		data[RandStringRunes(10)] = struct{}{}
	}
	for i := 0; i < b.N; i++ {
		_, ok := data["not_exist"]
		_ = ok
	}
}
