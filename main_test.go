package main

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/stretchr/testify/assert"
	"nothing.com/scratch/proto/spex/gen/go/voucher_mp_loader.pb"
)

func TestLoad(t *testing.T) {
	reqMeta := &voucher_mp_loader.RequestMeta{}
	_ = faker.FakeData(reqMeta,
		options.WithRandomMapAndSliceMaxSize(20),
		options.WithRandomStringLength(5),
	)
	t.Log("reqMeta:", reqMeta)
	assert.NotNil(t, reqMeta)
}

func BenchmarkLoad(b *testing.B) {
	reqMeta := &voucher_mp_loader.RequestMeta{}
	_ = faker.FakeData(reqMeta,
		options.WithRandomMapAndSliceMaxSize(20),
		options.WithRandomStringLength(5),
	)

	b.Run("copy reqMeta", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rst := CopyMessageV2(&voucher_mp_loader.RequestMetaV2{}, reqMeta)
			_ = rst
		}
	})

	b.Run("copy reqMeta V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rst := CopyMessage(&voucher_mp_loader.RequestMetaV2{}, reqMeta, nil)
			_ = rst
		}
	})
}
