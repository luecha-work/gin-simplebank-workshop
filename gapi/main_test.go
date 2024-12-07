package gapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	db "github.com/techschool/simple_bank/db/sqlc"
	"github.com/techschool/simple_bank/utils"
	"github.com/techschool/simple_bank/worker"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
