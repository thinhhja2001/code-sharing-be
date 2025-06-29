package repo

import "context"

type Repository interface {
	Insert(ctx context.Context, entity interface{}) (interface{}, error)
	Get(ctx context.Context, entity interface{}) (interface{}, error)
	GetList(ctx context.Context, entities []interface{}) ([]interface{}, error)
}
