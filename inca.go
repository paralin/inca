package inca

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"hash/crc32"
	"math"
	"sort"

	"github.com/aperturerobotics/pbobject"
	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-crypto"
	lpeer "github.com/libp2p/go-libp2p-peer"
)

// GetObjectTypeID returns the object type string, used to identify types.
func (g *Genesis) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/genesis")
}

// GetObjectTypeID returns the object type string, used to identify types.
func (g *NodeMessage) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/node-message")
}

// GetObjectTypeID returns the object type string, used to identify types.
func (g *ValidatorSet) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/validator-set")
}

// GetObjectTypeID returns the object type string, used to identify types.
func (g *ChainConfig) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/chain-config")
}

// GetObjectTypeID returns the object type string, used to identify types.
func (b *BlockHeader) GetObjectTypeID() *pbobject.ObjectTypeID {
	return &pbobject.ObjectTypeID{TypeUuid: "/inca/block-header"}
}

// GetObjectTypeID returns the object type string, used to identify types.
func (b *Block) GetObjectTypeID() *pbobject.ObjectTypeID {
	return &pbobject.ObjectTypeID{TypeUuid: "/inca/block"}
}

// GetObjectTypeID returns the object type string, used to identify types.
func (b *Vote) GetObjectTypeID() *pbobject.ObjectTypeID {
	return &pbobject.ObjectTypeID{TypeUuid: "/inca/vote"}
}

// Computes the sort heuristic.
// Validators are sorted by crc32 of address and voting power concatinated.
func (v *Validator) ComputeSortHeuristic() uint32 {
	votingPowerBin := proto.EncodeVarint(v.GetVotingPower())
	appn := make([]byte, len(v.GetPubKey())+len(votingPowerBin))
	copy(appn, v.GetPubKey())
	copy(appn[len(v.GetPubKey()):], votingPowerBin)
	return crc32.ChecksumIEEE(appn)
}

// SortValidators sorts the validator set correctly.
// Validators are sorted by crc32 of address and voting power concatinated.
func (g *ValidatorSet) SortValidators() {
	if len(g.GetValidators()) < 2 {
		return
	}

	// less function
	validators := g.GetValidators()
	sort.Slice(validators, func(i int, j int) bool {
		return validators[i].ComputeSortHeuristic() < validators[j].ComputeSortHeuristic()
	})
}

func hashBytesUint64(buf []byte, num uint64) [sha1.Size]byte {
	var datBuf bytes.Buffer
	// TODO: determine if this is consistent
	datBuf.Write(buf)
	binary.Write(&datBuf, binary.LittleEndian, num)
	return sha1.Sum(datBuf.Bytes())
}

// SelectProposer selects a validator from the set to propose.
func (g *ValidatorSet) SelectProposer(lastBlockDigest []byte, height, round uint64) (*Validator, int32) {
	var buf bytes.Buffer
	h1 := hashBytesUint64(lastBlockDigest[:8], height)
	buf.Write(h1[:])
	h2 := hashBytesUint64(lastBlockDigest[8:16], round)
	buf.Write(h2[:])

	var validators []*Validator
	var weightSum int32
	for _, v := range g.GetValidators() {
		if v.GetOperationMode() != Validator_OperationMode_OPERATING {
			continue
		}

		weightSum += int32(v.GetVotingPower())
		validators = append(validators, v)
	}

	// bufCrcK will be between 0 and maxWeight
	bufCrcK := int32((float32(crc32.ChecksumIEEE(buf.Bytes())) / float32(math.MaxUint32)) * float32(weightSum))
	for _, validator := range validators {
		bufCrcK -= int32(validator.GetVotingPower())
		if bufCrcK <= 0 {
			return validator, weightSum
		}
	}

	return nil, weightSum
}

// ParsePublicKey parses the validator's public key.
func (v *Validator) ParsePublicKey() (crypto.PubKey, error) {
	return crypto.UnmarshalPublicKey(v.GetPubKey())
}

// ParsePeerID parses the peer id and public key.
func (v *Validator) ParsePeerID() (lpeer.ID, crypto.PubKey, error) {
	pubKey, err := v.ParsePublicKey()
	if err != nil {
		return lpeer.ID(""), nil, err
	}

	id, err := lpeer.IDFromPublicKey(pubKey)
	if err != nil {
		return lpeer.ID(""), nil, err
	}

	return id, pubKey, nil
}
