package gm_test

import (
	"testing"

	keepertest "github.com/notbdu/gm/testutil/keeper"
	"github.com/notbdu/gm/testutil/nullify"
	"github.com/notbdu/gm/x/gm"
	"github.com/notbdu/gm/x/gm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GmKeeper(t)
	gm.InitGenesis(ctx, *k, genesisState)
	got := gm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
