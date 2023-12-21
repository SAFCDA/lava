package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/testutil/common"
	"github.com/lavanet/lava/utils/sigs"
	"github.com/lavanet/lava/x/rewards/types"
	subscription "github.com/lavanet/lava/x/subscription/keeper"
	"github.com/stretchr/testify/require"
)

// for this test there are no relays, this means no rewards will be given to the providers, and this means no bonus rewards should be sent
func TestZeroProvidersRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment from the subscription
	// example: chain starts at block 1 and next distribution is in 100blocks, block 101 (1 month)
	// a user bought a subscription in block 2, it will expire in block 102
	// in block 101 we will distribute rewards but since there hasnt been any subscription monthly expiration no rewards were sent and thus the first month is always 0 rewards
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 0)

	ts.AdvanceMonths(1)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 0)
}

// the rewards here is maxboost*totalbaserewards, in this test the rewards for the providers are low (first third of the graph)
func TestBasicBoostProvidersRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	baserewards := uint64(100)
	// the rewards by the subscription will be limited by LIMIT_TOKEN_PER_CU
	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, baserewards)
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, res.Rewards[0].Amount.Amount.Uint64(), baserewards*subscription.LIMIT_TOKEN_PER_CU)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, res.Rewards[0].Amount.Amount.Uint64(), baserewards*subscription.LIMIT_TOKEN_PER_CU+baserewards*subscription.LIMIT_TOKEN_PER_CU)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// the rewards here is spec payout allocation (full rewards from the pool), in this test the rewards for the providers are medium (second third of the graph)
func TestSpecAllocationProvidersRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, res.Rewards[0].Amount, ts.plan.Price)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, distBalance.QuoRaw(int64(ts.Keepers.Rewards.MaxRewardBoost(ts.Ctx))), res.Rewards[0].Amount.Amount)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// the rewards here is the diminishing part of the reward, in this test the rewards for the providers are high (third third of the graph)
func TestProvidersDiminishingRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	for i := 0; i < 7; i++ {
		consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
		_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
		require.Nil(t, err)

		msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
		_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
		require.Nil(t, err)
	}

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, res.Rewards[0].Amount.Amount, ts.plan.Price.Amount.MulRaw(7))
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)

	require.Equal(t, sdk.NewDecWithPrec(15, 1).MulInt(distBalance).Sub(sdk.NewDecWithPrec(5, 1).MulInt(ts.plan.Price.Amount.MulRaw(7))).TruncateInt().QuoRaw(int64(ts.Keepers.Rewards.MaxRewardBoost(ts.Ctx))), res.Rewards[0].Amount.Amount)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// the rewards here is the zero since base rewards are very big, in this test the rewards for the providers are at the end of the graph
func TestProvidersEndRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	for i := 0; i < 50; i++ {
		consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
		_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
		require.Nil(t, err)

		msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
		_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
		require.Nil(t, err)
	}

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, res.Rewards[0].Amount.Amount, ts.plan.Price.Amount.MulRaw(50))
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 0)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// in this test we create 2 specs with 1 provider each, one of the specs shares is zero
// this means that no matter how much rewards the providers in this spec will get, they will get 0 bonus rewards
func Test2SpecsZeroShares(t *testing.T) {
	ts := newTester(t)
	spec2 := ts.spec
	spec2.Index = "mock2"
	spec2.Name = spec2.Index
	spec2.Shares = 0
	ts.AddSpec(spec2.Index, spec2)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, 2*testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	err = ts.StakeProvider(providerAcc.Addr.String(), spec2, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	consumerAcc2, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc2.Addr.String(), consumerAcc2.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg = ts.SendRelay(providerAcc.Addr.String(), consumerAcc2, []string{spec2.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 2)
	require.Equal(t, res.Rewards[0].Amount, ts.plan.Price)
	require.Equal(t, res.Rewards[1].Amount, ts.plan.Price)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), "")
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, distBalance.QuoRaw(int64(ts.Keepers.Rewards.MaxRewardBoost(ts.Ctx))), res.Rewards[0].Amount.Amount)
	require.Equal(t, res.Rewards[0].ChainId, ts.spec.Index)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// 2 specs with one of them double the shares than the other
// the providers will have the same amount of CU used, thus the same rewards
// the bonus for the provider with double the shares should be double than the other provider
func Test2SpecsDoubleShares(t *testing.T) {
	ts := newTester(t)
	spec2 := ts.spec
	spec2.Index = "mock2"
	spec2.Name = spec2.Index
	spec2.Shares *= 2
	ts.AddSpec(spec2.Index, spec2)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, 2*testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	err = ts.StakeProvider(providerAcc.Addr.String(), spec2, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	consumerAcc2, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc2.Addr.String(), consumerAcc2.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg = ts.SendRelay(providerAcc.Addr.String(), consumerAcc2, []string{spec2.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 2)
	require.Equal(t, res.Rewards[0].Amount, ts.plan.Price)
	require.Equal(t, res.Rewards[1].Amount, ts.plan.Price)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), "")
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 2)
	require.Equal(t, res.Rewards[0].Amount.Amount.QuoRaw(2), res.Rewards[1].Amount.Amount)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// in this test we setup 3 providers, each with different cu used (-> 3 different rewards from the plan) (1,2,4)
// the providers should get bonus rewards according to their plan rewards
func TestBonusRewards3Providers(t *testing.T) {
	ts := newTester(t)

	providerAcc1, _ := ts.AddAccount(common.PROVIDER, 1, 2*testBalance)
	err := ts.StakeProvider(providerAcc1.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	providerAcc2, _ := ts.AddAccount(common.PROVIDER, 2, 2*testBalance)
	err = ts.StakeProvider(providerAcc2.Addr.String(), ts.spec, 2*testBalance)
	require.Nil(t, err)

	providerAcc3, _ := ts.AddAccount(common.PROVIDER, 3, 3*testBalance)
	err = ts.StakeProvider(providerAcc3.Addr.String(), ts.spec, 3*testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg := ts.SendRelay(providerAcc1.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64()/2)
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	msg = ts.SendRelay(providerAcc2.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	msg = ts.SendRelay(providerAcc3.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64()*2)
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc1.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	_, err = ts.TxDualstakingClaimRewards(providerAcc1.Addr.String(), "")
	require.Nil(t, err)

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc2.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	_, err = ts.TxDualstakingClaimRewards(providerAcc2.Addr.String(), "")
	require.Nil(t, err)

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc3.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	_, err = ts.TxDualstakingClaimRewards(providerAcc3.Addr.String(), "")
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	res1, err := ts.QueryDualstakingDelegatorRewards(providerAcc1.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	// we sub 3 because of truncating
	require.Equal(t, res1.Rewards[0].Amount.Amount, distBalance.QuoRaw(7*int64(ts.Keepers.Rewards.MaxRewardBoost(ts.Ctx))).SubRaw(1))
	_, err = ts.TxDualstakingClaimRewards(providerAcc1.Addr.String(), providerAcc1.Addr.String())
	require.Nil(t, err)

	res2, err := ts.QueryDualstakingDelegatorRewards(providerAcc2.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	// we sub 1 because of truncating
	require.Equal(t, res2.Rewards[0].Amount.Amount, distBalance.QuoRaw(7*int64(ts.Keepers.Rewards.MaxRewardBoost(ts.Ctx))).MulRaw(2))
	_, err = ts.TxDualstakingClaimRewards(providerAcc2.Addr.String(), providerAcc2.Addr.String())
	require.Nil(t, err)

	res3, err := ts.QueryDualstakingDelegatorRewards(providerAcc3.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	// we add 6 because of truncating
	require.Equal(t, res3.Rewards[0].Amount.Amount, distBalance.QuoRaw(7*int64(ts.Keepers.Rewards.MaxRewardBoost(ts.Ctx))).MulRaw(4).AddRaw(1))
	_, err = ts.TxDualstakingClaimRewards(providerAcc3.Addr.String(), providerAcc3.Addr.String())
	require.Nil(t, err)
}

func TestBonusRewardsEquall5Providers(t *testing.T) {
	ts := newTester(t)

	count := 5
	providerAccs := []sigs.Account{}
	consAccs := []sigs.Account{}

	for i := 0; i < count; i++ {
		providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
		err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
		providerAccs = append(providerAccs, providerAcc)
		require.Nil(t, err)

		consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
		_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
		consAccs = append(consAccs, consumerAcc)
		require.Nil(t, err)
	}

	for i := 1; i < 10; i++ {
		ts.AdvanceEpoch()

		for _, providerAcc := range providerAccs {
			for _, consAcc := range consAccs {
				msg := ts.SendRelay(providerAcc.Addr.String(), consAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64()/uint64(count)/1000)
				_, err := ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
				require.Nil(t, err)
			}
		}
	}

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	for _, providerAcc := range providerAccs {
		res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), "", "")
		require.Nil(t, err)
		require.Len(t, res.Rewards, 1)
		_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), "")
		require.Nil(t, err)
	}

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	for _, providerAcc := range providerAccs {
		res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), "", "")
		require.Nil(t, err)
		require.Len(t, res.Rewards, 1)
		require.Equal(t, distBalance.QuoRaw(int64(count)), res.Rewards[0].Amount.Amount)
		_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), "")
		require.Nil(t, err)
	}
}

// in this test we have 5 providers and 5 consumers
// all the providers serve the same amount of cu in total
// cons1 relays only to prov1 -> expected adjustment 1/5 (1 out of maxrewardboost)
// cons2-5 relays to all prov2-5 -> expected adjustment 4/5 (1 out of maxrewardboost)
func TestBonusRewards5Providers(t *testing.T) {
	ts := newTester(t)

	count := 5
	providerAccs := []sigs.Account{}
	consAccs := []sigs.Account{}

	for i := 0; i < count; i++ {
		providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
		err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
		providerAccs = append(providerAccs, providerAcc)
		require.Nil(t, err)

		consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
		_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
		consAccs = append(consAccs, consumerAcc)
		require.Nil(t, err)
	}

	for i := 1; i < 10; i++ {
		ts.AdvanceEpoch()

		msg := ts.SendRelay(providerAccs[0].Addr.String(), consAccs[0], []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64()/100)
		_, err := ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
		require.Nil(t, err)

		for _, providerAcc := range providerAccs[1:] {
			for _, consAcc := range consAccs[1:] {
				msg := ts.SendRelay(providerAcc.Addr.String(), consAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64()/uint64(count)/100)
				_, err := ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
				require.Nil(t, err)
			}
		}
	}

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	for _, providerAcc := range providerAccs {
		res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), "", "")
		require.Nil(t, err)
		require.Len(t, res.Rewards, 1)
		_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), "")
		require.Nil(t, err)
	}

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	// distribution pool divided between all providers (5) equally (they served the same amount of CU in total)
	fullProvReward := distBalance.QuoRaw(5)
	for i, providerAcc := range providerAccs {
		var expected math.Int
		if i == 0 {
			// gets only 1/5 of the full reward (sub 1 for trancating)
			expected = fullProvReward.MulRaw(1).QuoRaw(5).SubRaw(1)
		} else {
			// gets only 4/5 of the full reward (sub 2 for trancating)
			expected = fullProvReward.MulRaw(4).QuoRaw(5).SubRaw(2)
		}
		res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), "", "")
		require.Nil(t, err)
		require.Len(t, res.Rewards, 1)
		require.Equal(t, expected, res.Rewards[0].Amount.Amount)
		_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), "")
		require.Nil(t, err)
	}
}
