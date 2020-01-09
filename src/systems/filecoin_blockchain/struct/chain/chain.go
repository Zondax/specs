package chain

import (
	"github.com/filecoin-project/specs/actors"
	"github.com/filecoin-project/specs/util"
)

// Returns the tipset at or immediately prior to `epoch`.
func (chain *Chain_I) TipsetAtEpoch(epoch actors.ChainEpoch) Tipset {
	current := chain.HeadTipset()
	for current.Epoch() > epoch {
		current = current.Parents()
	}

	return current
}

// Draws randomness from the tipset at or immediately prior to `epoch`.
func (chain *Chain_I) RandomnessAtEpoch(epoch actors.ChainEpoch) util.Bytes {
	ts := chain.TipsetAtEpoch(epoch)
	return ts.MinTicket().DrawRandomness(epoch)
}
