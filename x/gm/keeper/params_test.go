package keeper_test

import (
	"testing"

	testkeeper "github.com/notbdu/gm/testutil/keeper"
	"github.com/notbdu/gm/x/gm/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GmKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
