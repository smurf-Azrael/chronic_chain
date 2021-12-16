package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilityTypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	connectionTypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	channelTypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibcExported "github.com/cosmos/ibc-go/modules/core/exported"
)

// BankViewKeeper defines a subset of methods implemented by the cosmos-sdk bank keeper
type BankViewKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

// Burner is a subset of the sdk bank keeper methods
type Burner interface {
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// BankKeeper defines a subset of methods implemented by the cosmos-sdk bank keeper
type BankKeeper interface {
	BankViewKeeper
	Burner
	IsSendEnabledCoins(ctx sdk.Context, coins ...sdk.Coin) error
	BlockedAddr(addr sdk.AccAddress) bool
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

// AccountKeeper defines a subset of methods implemented by the cosmos-sdk account keeper
type AccountKeeper interface {
	// NewAccountWithAddress Return a new account with the next account number and the specified address. Does not save the new account to the store.
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authTypes.AccountI
	// GetAccount Retrieve an account from the store.
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authTypes.AccountI
	// SetAccount Set an account in the store.
	SetAccount(ctx sdk.Context, acc authTypes.AccountI)
}

// DistributionKeeper defines a subset of methods implemented by the cosmos-sdk distribution keeper
type DistributionKeeper interface {
	DelegationRewards(c context.Context, req *types.QueryDelegationRewardsRequest) (*types.QueryDelegationRewardsResponse, error)
}

// StakingKeeper defines a subset of methods implemented by the cosmos-sdk staking keeper
type StakingKeeper interface {
	// BondDenom - Bondable coin denomination
	BondDenom(ctx sdk.Context) (res string)
	// GetValidator get a single validator
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingTypes.Validator, found bool)
	// GetBondedValidatorsByPower get the current group of bonded validators sorted by power-rank
	GetBondedValidatorsByPower(ctx sdk.Context) []stakingTypes.Validator
	// GetAllDelegatorDelegations return all delegations for a delegator
	GetAllDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress) []stakingTypes.Delegation
	// GetDelegation return a specific delegation
	GetDelegation(ctx sdk.Context,
		delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation stakingTypes.Delegation, found bool)
	// HasReceivingRedelegation check if validator is receiving a redelegation
	HasReceivingRedelegation(ctx sdk.Context,
		delAddr sdk.AccAddress, valDstAddr sdk.ValAddress) bool
}

// ChannelKeeper defines the expected IBC channel keeper
type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channelTypes.Channel, found bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
	SendPacket(ctx sdk.Context, channelCap *capabilityTypes.Capability, packet ibcExported.PacketI) error
	ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilityTypes.Capability) error
	GetAllChannels(ctx sdk.Context) (channels []channelTypes.IdentifiedChannel)
	IterateChannels(ctx sdk.Context, cb func(channelTypes.IdentifiedChannel) bool)
}

// ClientKeeper defines the expected IBC client keeper
type ClientKeeper interface {
	GetClientConsensusState(ctx sdk.Context, clientID string) (connection ibcExported.ConsensusState, found bool)
}

// ConnectionKeeper defines the expected IBC connection keeper
type ConnectionKeeper interface {
	GetConnection(ctx sdk.Context, connectionID string) (connection connectionTypes.ConnectionEnd, found bool)
}

// PortKeeper defines the expected IBC port keeper
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilityTypes.Capability
}

type CapabilityKeeper interface {
	GetCapability(ctx sdk.Context, name string) (*capabilityTypes.Capability, bool)
	ClaimCapability(ctx sdk.Context, cap *capabilityTypes.Capability, name string) error
	AuthenticateCapability(ctx sdk.Context, capability *capabilityTypes.Capability, name string) bool
}

// ICS20TransferPortSource is a subset of the ibc transfer keeper.
type ICS20TransferPortSource interface {
	GetPort(ctx sdk.Context) string
}