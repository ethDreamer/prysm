package precompute

import (
	"context"
	"testing"

	"github.com/prysmaticlabs/go-bitfield"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/epoch"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	ethpb "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/shared/params"
)

func TestProcessRewardsAndPenaltiesPrecompute(t *testing.T) {
	helpers.ClearAllCaches()
	e := params.BeaconConfig().SlotsPerEpoch
	validatorCount := uint64(2048)
	state := buildState(e+3, validatorCount)
	startShard := uint64(960)
	atts := make([]*pb.PendingAttestation, 3)
	for i := 0; i < len(atts); i++ {
		atts[i] = &pb.PendingAttestation{
			Data: &ethpb.AttestationData{
				Crosslink: &ethpb.Crosslink{
					Shard:    startShard + uint64(i),
					DataRoot: []byte{'A'},
				},
				Target: &ethpb.Checkpoint{},
				Source: &ethpb.Checkpoint{},
			},
			AggregationBits: bitfield.Bitlist{0xC0, 0xC0, 0xC0, 0xC0, 0x01},
			InclusionDelay:  1,
		}
	}
	state.PreviousEpochAttestations = atts
	state.CurrentCrosslinks[startShard] = &ethpb.Crosslink{
		DataRoot: []byte{'A'},
	}
	state.CurrentCrosslinks[startShard+1] = &ethpb.Crosslink{
		DataRoot: []byte{'A'},
	}
	state.CurrentCrosslinks[startShard+2] = &ethpb.Crosslink{
		DataRoot: []byte{'A'},
	}

	vp, bp := New(context.Background(), state)
	vp, bp, err := ProcessAttestations(context.Background(), state, vp, bp)
	if err != nil {
		t.Fatal(err)
	}

	state, err = ProcessRewardsAndPenaltiesPrecompute(state, bp, vp)
	if err != nil {
		t.Fatal(err)
	}

	// Indices that voted everything except for head, lost a bit money
	wanted := uint64(31999995452)
	if state.Balances[4] != wanted {
		t.Errorf("wanted balance: %d, got: %d",
			wanted, state.Balances[4])
	}

	// Indices that did not vote, lost more money
	wanted = uint64(31999949392)
	if state.Balances[0] != wanted {
		t.Errorf("wanted balance: %d, got: %d",
			wanted, state.Balances[0])
	}
}

func TestAttestationDeltaPrecompute(t *testing.T) {
	helpers.ClearAllCaches()
	e := params.BeaconConfig().SlotsPerEpoch
	validatorCount := uint64(2048)
	state := buildState(e+2, validatorCount)
	startShard := uint64(960)
	atts := make([]*pb.PendingAttestation, 3)
	for i := 0; i < len(atts); i++ {
		atts[i] = &pb.PendingAttestation{
			Data: &ethpb.AttestationData{
				Crosslink: &ethpb.Crosslink{
					Shard:    startShard + uint64(i),
					DataRoot: []byte{'A'},
				},
				Target: &ethpb.Checkpoint{},
				Source: &ethpb.Checkpoint{},
			},
			AggregationBits: bitfield.Bitlist{0xC0, 0xC0, 0xC0, 0xC0, 0x01},
			InclusionDelay:  1,
		}
	}
	state.PreviousEpochAttestations = atts
	state.CurrentCrosslinks[startShard] = &ethpb.Crosslink{
		DataRoot: []byte{'A'},
	}
	state.CurrentCrosslinks[startShard+1] = &ethpb.Crosslink{
		DataRoot: []byte{'A'},
	}

	vp, bp := New(context.Background(), state)
	vp, bp, err := ProcessAttestations(context.Background(), state, vp, bp)
	if err != nil {
		t.Fatal(err)
	}

	rewards, penalties, err := attestationDeltas(state, bp, vp)
	if err != nil {
		t.Fatal(err)
	}
	attestedBalance, err := epoch.AttestingBalance(state, atts)
	if err != nil {
		t.Error(err)
	}
	totalBalance, err := helpers.TotalActiveBalance(state)
	if err != nil {
		t.Fatal(err)
	}

	attestedIndices := []uint64{5, 754, 797, 1637, 1770, 1862, 1192}
	for _, i := range attestedIndices {
		base, err := epoch.BaseReward(state, i)
		if err != nil {
			t.Errorf("Could not get base reward: %v", err)
		}
		// Base rewards for getting source right
		wanted := 3 * (base * attestedBalance / totalBalance)
		// Base rewards for proposer and attesters working together getting attestation
		// on chain in the fatest manner
		proposerReward := base / params.BeaconConfig().ProposerRewardQuotient
		wanted += (base - proposerReward) * params.BeaconConfig().MinAttestationInclusionDelay
		if rewards[i] != wanted {
			t.Errorf("Wanted reward balance %d, got %d", wanted, rewards[i])
		}
		// Since all these validators attested, they shouldn't get penalized.
		if penalties[i] != 0 {
			t.Errorf("Wanted penalty balance 0, got %d", penalties[i])
		}
	}

	nonAttestedIndices := []uint64{12, 23, 45, 79}
	for _, i := range nonAttestedIndices {
		base, err := epoch.BaseReward(state, i)
		if err != nil {
			t.Errorf("Could not get base reward: %v", err)
		}
		wanted := 3 * base
		// Since all these validators did not attest, they shouldn't get rewarded.
		if rewards[i] != 0 {
			t.Errorf("Wanted reward balance 0, got %d", rewards[i])
		}
		// Base penalties for not attesting.
		if penalties[i] != wanted {
			t.Errorf("Wanted penalty balance %d, got %d", wanted, penalties[i])
		}
	}
}

func TestCrosslinkDeltaPrecompute(t *testing.T) {
	helpers.ClearAllCaches()
	e := params.BeaconConfig().SlotsPerEpoch
	helpers.ClearShuffledValidatorCache()
	validatorCount := uint64(2048)
	state := buildState(e+2, validatorCount)
	startShard := uint64(960)
	atts := make([]*pb.PendingAttestation, 2)
	for i := 0; i < len(atts); i++ {
		atts[i] = &pb.PendingAttestation{
			Data: &ethpb.AttestationData{
				Crosslink: &ethpb.Crosslink{
					Shard:    startShard + uint64(i),
					DataRoot: []byte{'A'},
				},
				Target: &ethpb.Checkpoint{},
				Source: &ethpb.Checkpoint{},
			},
			InclusionDelay:  uint64(i + 100),
			AggregationBits: bitfield.Bitlist{0xC0, 0xC0, 0xC0, 0xC0, 0x01},
		}
	}
	state.PreviousEpochAttestations = atts
	state.CurrentCrosslinks[startShard] = &ethpb.Crosslink{
		DataRoot: []byte{'A'}, Shard: startShard,
	}
	state.CurrentCrosslinks[startShard+1] = &ethpb.Crosslink{
		DataRoot: []byte{'A'}, Shard: startShard + 1,
	}

	vp, bp := New(context.Background(), state)
	vp, bp, err := ProcessAttestations(context.Background(), state, vp, bp)
	if err != nil {
		t.Fatal(err)
	}

	rewards, penalties, err := crosslinkDeltaPreCompute(state, bp, vp)
	if err != nil {
		t.Fatal(err)
	}

	attestedIndices := []uint64{5, 16, 336, 797, 1082, 1450, 1770, 1958}
	for _, i := range attestedIndices {
		// Since all these validators attested, they should get the same rewards.
		want := uint64(12649)
		if rewards[i] != want {
			t.Errorf("Wanted reward balance %d, got %d", want, rewards[i])
		}
		// Since all these validators attested, they shouldn't get penalized.
		if penalties[i] != 0 {
			t.Errorf("Wanted penalty balance 0, got %d", penalties[i])
		}
	}

	nonAttestedIndices := []uint64{12, 23, 45, 79}
	for _, i := range nonAttestedIndices {
		base, err := epoch.BaseReward(state, i)
		if err != nil {
			t.Errorf("Could not get base reward: %v", err)
		}
		wanted := base
		// Since all these validators did not attest, they shouldn't get rewarded.
		if rewards[i] != 0 {
			t.Errorf("Wanted reward balance 0, got %d", rewards[i])
		}
		// Base penalties for not attesting.
		if penalties[i] != wanted {
			t.Errorf("Wanted penalty balance %d, got %d", wanted, penalties[i])
		}
	}
}

func buildState(slot uint64, validatorCount uint64) *pb.BeaconState {
	validators := make([]*ethpb.Validator, validatorCount)
	for i := 0; i < len(validators); i++ {
		validators[i] = &ethpb.Validator{
			ExitEpoch:        params.BeaconConfig().FarFutureEpoch,
			EffectiveBalance: params.BeaconConfig().MaxEffectiveBalance,
		}
	}
	validatorBalances := make([]uint64, len(validators))
	for i := 0; i < len(validatorBalances); i++ {
		validatorBalances[i] = params.BeaconConfig().MaxEffectiveBalance
	}
	latestActiveIndexRoots := make(
		[][]byte,
		params.BeaconConfig().EpochsPerHistoricalVector,
	)
	for i := 0; i < len(latestActiveIndexRoots); i++ {
		latestActiveIndexRoots[i] = params.BeaconConfig().ZeroHash[:]
	}
	latestRandaoMixes := make(
		[][]byte,
		params.BeaconConfig().EpochsPerHistoricalVector,
	)
	for i := 0; i < len(latestRandaoMixes); i++ {
		latestRandaoMixes[i] = params.BeaconConfig().ZeroHash[:]
	}
	return &pb.BeaconState{
		Slot:                        slot,
		Balances:                    validatorBalances,
		Validators:                  validators,
		CurrentCrosslinks:           make([]*ethpb.Crosslink, params.BeaconConfig().ShardCount),
		RandaoMixes:                 make([][]byte, params.BeaconConfig().EpochsPerHistoricalVector),
		ActiveIndexRoots:            make([][]byte, params.BeaconConfig().EpochsPerHistoricalVector),
		CompactCommitteesRoots:      make([][]byte, params.BeaconConfig().EpochsPerSlashingsVector),
		Slashings:                   make([]uint64, params.BeaconConfig().EpochsPerSlashingsVector),
		BlockRoots:                  make([][]byte, params.BeaconConfig().SlotsPerEpoch*10),
		FinalizedCheckpoint:         &ethpb.Checkpoint{},
		PreviousJustifiedCheckpoint: &ethpb.Checkpoint{},
		CurrentJustifiedCheckpoint:  &ethpb.Checkpoint{},
	}
}
