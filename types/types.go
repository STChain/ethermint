package types

import (
    "reflect"

    "github.com/ethereum/go-ethereum/common"
    ethTypes "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/log"
    "github.com/tendermint/abci/types"
)

// MinerRewardStrategy is a mining strategy
type MinerRewardStrategy interface {
    Receiver() common.Address
    SetReceiver(address string)
}

// ValidatorsStrategy is a validator strategy
type ValidatorsStrategy interface {
    SetValidators(validators []*types.Validator)
    CollectTx(tx *ethTypes.Transaction)
    GetUpdatedValidators() []*types.Validator
}

// Strategy encompasses all available strategies
type Strategy struct {
    coinbase string
    currentValidators []*types.Validator
    MinerRewardStrategy
    ValidatorsStrategy
}



func (strategy *Strategy) Receiver() common.Address {
    return common.HexToAddress(strategy.coinbase)
}

func (strategy *Strategy) SetReceiver(address string) {
    strategy.coinbase = address;
}

// SetValidators updates the current validators
func (strategy *Strategy) SetValidators(validators []*types.Validator) {
    strategy.currentValidators = validators
}

// CollectTx collects the rewards for a transaction
func (strategy *Strategy) CollectTx(tx *ethTypes.Transaction) {
    if reflect.DeepEqual(tx.To(), common.HexToAddress("0000000000000000000000000000000000000001")) {
        log.Info("Adding validator", "data", tx.Data())
        strategy.currentValidators = append(
            strategy.currentValidators,
            &types.Validator{
                PubKey: tx.Data(),
                Power:  tx.Value().Uint64(),
            },
        )
    }
}

// GetUpdatedValidators returns the current validators
func (strategy *Strategy) GetUpdatedValidators() []*types.Validator {
    return strategy.currentValidators
}