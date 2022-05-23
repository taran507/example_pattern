package facade

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

const getUsersNames = `SELECT name FROM user`

type Provider struct {
	pool *pgxpool.Pool
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) GetUsersName(ctx context.Context) ([]string, error) {
	rows, err := p.pool.Query(ctx, getUsersNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		names []string
		name  string
	)

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}
