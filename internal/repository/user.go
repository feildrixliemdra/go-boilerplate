package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"go-boilerplate/internal/model"

	sq "github.com/Masterminds/squirrel"
)

type IUserRepository interface {
	Create(ctx context.Context, user model.User) error
	GetBy(ctx context.Context, usr model.User) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user model.User) error
	DeleteByID(ctx context.Context, id uint64) error
}

type user struct {
	DB *sqlx.DB
}

func NewUserRepository(opt Option) IUserRepository {
	return &user{
		DB: opt.DB,
	}
}

func (r *user) DeleteByID(ctx context.Context, id uint64) (err error) {
	query, args, err := sq.Update(model.User{}.TableName()).
		SetMap(
			sq.Eq{
				"deleted_at": "now()", //soft delete
			},
		).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return
}

func (r *user) Create(ctx context.Context, user model.User) (err error) {
	query, args, err := sq.Insert(model.User{}.TableName()).
		SetMap(
			sq.Eq{
				"email":      user.Email,
				"name":       user.Name,
				"password":   user.Password,
				"created_at": "now()",
				"updated_at": "now()",
			},
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return
}

func (r *user) GetBy(ctx context.Context, usr model.User) (*model.User, error) {
	var result model.User

	q := sq.Select(
		"id",
		"name",
		"email",
		"created_at",
		"updated_at",
	).
		From(result.TableName()).
		Where(sq.Eq{"deleted_at": nil})

	if usr.Name != "" {
		q = q.Where(sq.ILike{"name": "%" + usr.Name + "%"})
	}

	if usr.Email != "" {
		q = q.Where(sq.Eq{"email": usr.Email})
	}

	if usr.ID != 0 {
		q = q.Where(sq.Eq{"id": usr.ID})
	}

	query, args, err := q.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	err = r.DB.GetContext(ctx, &result, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (r *user) GetAll(ctx context.Context) (result []model.User, err error) {
	result = []model.User{}

	query, args, err := sq.Select(
		"id",
		"name",
		"email",
		"created_at",
		"updated_at",
	).
		From(model.User{}.TableName()).
		Where(sq.Eq{"deleted_at": nil}).
		ToSql()

	if err != nil {
		return
	}

	err = r.DB.SelectContext(ctx, &result, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, nil
		}

		return
	}

	return
}

func (r *user) Update(ctx context.Context, user model.User) (err error) {

	query, args, err := sq.Update(model.User{}.TableName()).
		SetMap(
			sq.Eq{
				"email":      user.Email,
				"name":       user.Name,
				"password":   user.Password,
				"updated_at": "now()",
			},
		).
		Where(sq.Eq{"id": user.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return
}
