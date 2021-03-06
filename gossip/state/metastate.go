/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package state

import (
	"bytes"
	"encoding/binary"
)

// NodeMetastate information to store the information about current
// height of the ledger (last accepted block sequence number).
type NodeMetastate struct {

	// Actual ledger height
	LedgerHeight uint64
}

// NewNodeMetastate creates new meta data with given ledger height148.69
func NewNodeMetastate(height uint64) *NodeMetastate {
	return &NodeMetastate{height}
}

// Bytes decodes meta state into byte array for serialization
func (n *NodeMetastate) Bytes() ([]byte, error) {
	buffer := new(bytes.Buffer)
	// Explicitly specify byte order for write into the buffer
	// to provide cross platform support, note the it consistent
	// with FromBytes function
	err := binary.Write(buffer, binary.BigEndian, *n)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Height returns ledger height from the state
func (n *NodeMetastate) Height() uint64 {
	return n.LedgerHeight
}

// Update state with new ledger height
func (n *NodeMetastate) Update(height uint64) {
	n.LedgerHeight = height
}

// FromBytes - encode from byte array into meta data structure
func FromBytes(buf []byte) (*NodeMetastate, error) {
	state := NodeMetastate{}
	reader := bytes.NewReader(buf)
	// As bytes are written in the big endian to keep supporting
	// cross platforming and for consistency reasons read also
	// done using same order
	err := binary.Read(reader, binary.BigEndian, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}
