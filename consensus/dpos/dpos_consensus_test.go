package dpos

import (
	"math"
	"math/big"
	"testing"
)

func TestBlockReward(t *testing.T) {
	// begin with 100 reward
	var (
		// begin with 100 reward
		basicReward float64 = 100
		annulProfit         = 1.15
		// annulBlockAmount = big.NewInt(10)
		annulBlockAmount   = big.NewInt(3153600)
		blockReward        = big.NewInt(1e+18)
		currentBlockHeight = big.NewInt(0)
	)

	for i := 0; i < 100_000_000; i++ {
		currentBlockHeight.Add(currentBlockHeight, big.NewInt(1))
		yearNumber := currentBlockHeight.Int64() / annulBlockAmount.Int64()
		currentReward := (int64)(basicReward * math.Pow(annulProfit, float64(yearNumber)))
		if currentBlockHeight.Int64()%annulBlockAmount.Int64() == 0 {
			t.Logf("block number=%d currentReward=%d", currentBlockHeight, currentReward)
			precisionReward := new(big.Int).Mul(big.NewInt(currentReward), blockReward)
			t.Log(precisionReward)
		}
		//t.Logf("block number=%d currentReward=%d",currentBlockHeight,currentReward)
		//precisionReward := new(big.Int).Mul(big.NewInt(currentReward), blockReward)
		//t.Log(precisionReward)
	}

}
