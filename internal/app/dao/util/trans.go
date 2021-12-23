package util

import (
	"context"

	"bio-assay-ms/internal/app/contextx"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var TransSet = wire.NewSet(wire.Struct(new(Trans), "*"))

type Trans struct {
	DB *gorm.DB
}

func (a *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := contextx.FromTrans(ctx); ok {
		return fn(ctx)
	}

	return a.DB.Transaction(func(db *gorm.DB) error {
		return fn(contextx.NewTrans(ctx, db))
	})
}