package fuzz_utils

import (
	"math/rand"

	"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "fuzz_utils")

// /////////////////////////////////////////////////////////////////////////////
// status fuzzers
// /////////////////////////////////////////////////////////////////////////////

func FuzzStatusResponse(forkDigest *[]byte, finalizedRoot *[]byte, finalizedEpoch *primitives.Epoch, headRoot *[]byte, HeadSlot *primitives.Slot) {
	// forkDigest: [4]byte
	// finalizedRoot: [32]byte
	// finalizedEpoch: *primitives.Epoch
	// headRoot: *[]byte
	// HeadSlot: *primitives.Slot

	log.Info("FUZZ - FuzzStatusResponse BEFORE:")
	log.Info("\tforkDigest", forkDigest)
	log.Info("\tfinalizedRoot", finalizedRoot)
	log.Info("\tfinalizedEpoch", finalizedEpoch)
	log.Info("\theadRoot", headRoot)
	log.Info("\tHeadSlot", HeadSlot)

	// 20% chance to fuzz each field
	if rand.Intn(100) < 20 {
		log.Info("Mutating forkDigest")
		MutateNBytes(forkDigest, 4)
	}
	if rand.Intn(100) < 20 {
		// TODO: put valid root minus 1
		log.Info("Mutating finalizedRoot")
		MutateNBytes(finalizedRoot, 32)
	}
	if rand.Intn(100) < 20 {
		// TODO: put valid root minus 1

		log.Info("Mutating finalizedEpoch")
		*finalizedEpoch = primitives.Epoch(RandUint64())
	}
	if rand.Intn(100) < 20 {
		// TODO: put valid root minus 1

		log.Info("Mutating headRoot")
		MutateNBytes(headRoot, 32)
	}
	if rand.Intn(100) < 20 {
		// TODO: put valid root minus 1

		log.Info("Mutating HeadSlot")
		*HeadSlot = primitives.Slot(RandUint64())
	}

	log.Info("FUZZ - FuzzStatusResponse AFTER:")
	log.Info("\tforkDigest", forkDigest)
	log.Info("\tfinalizedRoot", finalizedRoot)
	log.Info("\tfinalizedEpoch", finalizedEpoch)
	log.Info("\theadRoot", headRoot)
	log.Info("\tHeadSlot", HeadSlot)
}