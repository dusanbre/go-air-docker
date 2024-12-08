package migrations

import (
	"context"
	"dusanbre/test-1/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().
			Model((*models.User)(nil)).
			Exec(ctx)

		if err != nil {
			panic(err)
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().Model((*models.User)(nil)).IfExists().Exec(ctx)
		if err != nil {
			panic(err)
		}
		return nil
	})
}
