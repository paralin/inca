package inca

import (
	"crypto/sha256"
	"testing"
)

func TestSelectProposer(t *testing.T) {
	blockDigest := sha256.Sum256([]byte("example"))
	vs := &ValidatorSet{
		Validators: []*Validator{
			&Validator{
				PubKey:      []byte("key1"),
				VotingPower: 10,
			},
			&Validator{
				PubKey:      []byte("key2"),
				VotingPower: 20,
			},
			&Validator{
				PubKey:      []byte("key3"),
				VotingPower: 2,
			},
		},
	}

	type heightRound struct {
		height uint64
		round  uint64
	}

	checkHeightRound := func(hr heightRound, expectedKey string) {
		val, _ := vs.SelectProposer(blockDigest[:], hr.height, hr.round)
		if val == nil {
			t.Fatal("returned nil validator")
		}

		pk := string(val.PubKey)
		if pk != expectedKey {
			t.Errorf("(%d, %d): selected %s != expected %s", hr.height, hr.round, pk, expectedKey)
		}
	}

	vs.SortValidators()

	// This test ensures that selections are deterministic between builds and Go versions
	checkHeightRound(heightRound{height: 0, round: 0}, "key2")
	checkHeightRound(heightRound{height: 0, round: 1}, "key2")
	checkHeightRound(heightRound{height: 0, round: 2}, "key1")
	checkHeightRound(heightRound{height: 0, round: 3}, "key2")
	checkHeightRound(heightRound{height: 0, round: 4}, "key1")
	checkHeightRound(heightRound{height: 0, round: 5}, "key2")
	checkHeightRound(heightRound{height: 1, round: 0}, "key2")
	checkHeightRound(heightRound{height: 2, round: 0}, "key2")
}
