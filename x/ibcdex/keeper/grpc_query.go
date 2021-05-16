package keeper

import (
	"github.com/darthlukan/starport-tutorial-app/x/ibcdex/types"
)

var _ types.QueryServer = Keeper{}
