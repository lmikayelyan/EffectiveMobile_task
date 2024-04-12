package repository

import (
	"context"
	"effective_mobile_task/internal/model"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Car interface {
	Create(ctx context.Context, items []model.Car) error
	Update(ctx context.Context, itemID uint, item *model.CarUpdate) error
	Get(ctx context.Context, limit, offset int) ([]model.Car, error)
	Delete(ctx context.Context, itemID uint) error
}

type car struct {
	pool *pgxpool.Pool
}

func CarRepo(pool *pgxpool.Pool) Car {
	return &car{pool: pool}
}

func (r *car) Create(ctx context.Context, items []model.Car) error {
	batch := &pgx.Batch{}
	batchQuery := "insert into cars (reg_number, mark, model, year, owner_id) values($1, $2, $3, $4, $5)"
	for _, item := range items {
		r.prepareBatchInsert(batch, &item, batchQuery)
	}

	br := r.pool.SendBatch(ctx, batch)
	_, err := br.Exec()
	if err != nil {
		return fmt.Errorf("cars.Create.batchExec: %v", err)
	}

	err = br.Close()
	if err != nil {
		return fmt.Errorf("cars.Create.batchClose: %v", err)
	}

	return nil
}

func (r *car) Update(ctx context.Context, itemID uint, item *model.CarUpdate) error {
	queryStr := "update cars set mark=$1, model=$2, owner_id=$3, year=$4 where id=$5"

	_, err := r.pool.Exec(ctx, queryStr, item.Mark, item.Model, item.OwnerID, itemID)

	return err
}

func (r *car) Delete(ctx context.Context, itemID uint) error {
	queryStr := "delete from cars where id=$1"
	_, err := r.pool.Exec(ctx, queryStr, itemID)

	return err
}

func (r *car) Get(ctx context.Context, limit, offset int) ([]model.Car, error) {
	queryStr := "select * from cars limit $1 offset $2"
	rows, err := r.pool.Query(ctx, queryStr, limit, offset)

	var cars []model.Car
	for rows.Next() {
		var item model.Car
		err = rows.Scan(&item.ID,
			&item.RegNum,
			&item.Mark,
			&item.Model,
			&item.Year,
			&item.OwnerID)

		if err != nil {
			return nil, fmt.Errorf("cars.Get.rows.Scan: %v", err)
		}

		cars = append(cars, item)
	}

	return cars, nil
}

func (r *car) prepareBatchInsert(batch *pgx.Batch, item *model.Car, query string) {
	batch.Queue(query,
		item.RegNum,
		item.Mark,
		item.Model,
		item.Year,
		item.OwnerID)
}
