# Inca: Design Doc

This is a design outline for the initial prototype of Inca.

Inca is loosely based on the [[Tendermint][https://github.com/tendermint/tendermint/wiki/Byzantine-Consensus-Algorithm]] algorithm (Apache 2 licensed).

## Core Concepts

 - A block contains zero or more transactions.
 - Deciding the next block is a process consisting of one or more rounds.
 - A validator may sign at most one block proposal per round.

