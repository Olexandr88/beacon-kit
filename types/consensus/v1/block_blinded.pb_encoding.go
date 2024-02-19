// Code generated by fastssz. DO NOT EDIT.
// Hash: bc1506d3f21b996dbdf984a227c006cded8ee9600b926efffd05f701835f73ac
package consensusv1

import (
	github_com_itsdevbear_bolaris_types_consensus_primitives "github.com/itsdevbear/bolaris/types/consensus/primitives"
	ssz "github.com/prysmaticlabs/fastssz"
	v1 "github.com/prysmaticlabs/prysm/v5/proto/engine/v1"
)

// MarshalSSZ ssz marshals the BlindedBeaconKitBlockCapella object
func (b *BlindedBeaconKitBlockCapella) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlindedBeaconKitBlockCapella object to a target array
func (b *BlindedBeaconKitBlockCapella) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(44)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Offset (1) 'Body'
	dst = ssz.WriteOffset(dst, offset)
	if b.Body == nil {
		b.Body = new(BlindedBeaconKitBlockBodyCapella)
	}
	offset += b.Body.SizeSSZ()

	// Field (2) 'PayloadValue'
	if size := len(b.PayloadValue); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PayloadValue", size, 32)
		return
	}
	dst = append(dst, b.PayloadValue...)

	// Field (1) 'Body'
	if dst, err = b.Body.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlindedBeaconKitBlockCapella object
func (b *BlindedBeaconKitBlockCapella) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 44 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'Slot'
	b.Slot = github_com_itsdevbear_bolaris_types_consensus_primitives.Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Offset (1) 'Body'
	if o1 = ssz.ReadOffset(buf[8:12]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 44 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'PayloadValue'
	if cap(b.PayloadValue) == 0 {
		b.PayloadValue = make([]byte, 0, len(buf[12:44]))
	}
	b.PayloadValue = append(b.PayloadValue, buf[12:44]...)

	// Field (1) 'Body'
	{
		buf = tail[o1:]
		if b.Body == nil {
			b.Body = new(BlindedBeaconKitBlockBodyCapella)
		}
		if err = b.Body.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlindedBeaconKitBlockCapella object
func (b *BlindedBeaconKitBlockCapella) SizeSSZ() (size int) {
	size = 44

	// Field (1) 'Body'
	if b.Body == nil {
		b.Body = new(BlindedBeaconKitBlockBodyCapella)
	}
	size += b.Body.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BlindedBeaconKitBlockCapella object
func (b *BlindedBeaconKitBlockCapella) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlindedBeaconKitBlockCapella object with a hasher
func (b *BlindedBeaconKitBlockCapella) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (1) 'Body'
	if err = b.Body.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'PayloadValue'
	if size := len(b.PayloadValue); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PayloadValue", size, 32)
		return
	}
	hh.PutBytes(b.PayloadValue)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BlindedBeaconKitBlockBodyCapella object
func (b *BlindedBeaconKitBlockBodyCapella) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlindedBeaconKitBlockBodyCapella object to a target array
func (b *BlindedBeaconKitBlockBodyCapella) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(136)

	// Field (0) 'RandaoReveal'
	if size := len(b.RandaoReveal); size != 96 {
		err = ssz.ErrBytesLengthFn("--.RandaoReveal", size, 96)
		return
	}
	dst = append(dst, b.RandaoReveal...)

	// Field (1) 'Graffiti'
	if size := len(b.Graffiti); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Graffiti", size, 32)
		return
	}
	dst = append(dst, b.Graffiti...)

	// Offset (2) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.Deposits); ii++ {
		offset += 4
		offset += b.Deposits[ii].SizeSSZ()
	}

	// Offset (3) 'ExecutionPayloadHeader'
	dst = ssz.WriteOffset(dst, offset)
	if b.ExecutionPayloadHeader == nil {
		b.ExecutionPayloadHeader = new(v1.ExecutionPayloadHeaderCapella)
	}
	offset += b.ExecutionPayloadHeader.SizeSSZ()

	// Field (2) 'Deposits'
	if size := len(b.Deposits); size > 16 {
		err = ssz.ErrListTooBigFn("--.Deposits", size, 16)
		return
	}
	{
		offset = 4 * len(b.Deposits)
		for ii := 0; ii < len(b.Deposits); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.Deposits[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.Deposits); ii++ {
		if dst, err = b.Deposits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (3) 'ExecutionPayloadHeader'
	if dst, err = b.ExecutionPayloadHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlindedBeaconKitBlockBodyCapella object
func (b *BlindedBeaconKitBlockBodyCapella) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 136 {
		return ssz.ErrSize
	}

	tail := buf
	var o2, o3 uint64

	// Field (0) 'RandaoReveal'
	if cap(b.RandaoReveal) == 0 {
		b.RandaoReveal = make([]byte, 0, len(buf[0:96]))
	}
	b.RandaoReveal = append(b.RandaoReveal, buf[0:96]...)

	// Field (1) 'Graffiti'
	if cap(b.Graffiti) == 0 {
		b.Graffiti = make([]byte, 0, len(buf[96:128]))
	}
	b.Graffiti = append(b.Graffiti, buf[96:128]...)

	// Offset (2) 'Deposits'
	if o2 = ssz.ReadOffset(buf[128:132]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 136 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (3) 'ExecutionPayloadHeader'
	if o3 = ssz.ReadOffset(buf[132:136]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Field (2) 'Deposits'
	{
		buf = tail[o2:o3]
		num, err := ssz.DecodeDynamicLength(buf, 16)
		if err != nil {
			return err
		}
		b.Deposits = make([]*Deposit, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.Deposits[indx] == nil {
				b.Deposits[indx] = new(Deposit)
			}
			if err = b.Deposits[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (3) 'ExecutionPayloadHeader'
	{
		buf = tail[o3:]
		if b.ExecutionPayloadHeader == nil {
			b.ExecutionPayloadHeader = new(v1.ExecutionPayloadHeaderCapella)
		}
		if err = b.ExecutionPayloadHeader.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlindedBeaconKitBlockBodyCapella object
func (b *BlindedBeaconKitBlockBodyCapella) SizeSSZ() (size int) {
	size = 136

	// Field (2) 'Deposits'
	for ii := 0; ii < len(b.Deposits); ii++ {
		size += 4
		size += b.Deposits[ii].SizeSSZ()
	}

	// Field (3) 'ExecutionPayloadHeader'
	if b.ExecutionPayloadHeader == nil {
		b.ExecutionPayloadHeader = new(v1.ExecutionPayloadHeaderCapella)
	}
	size += b.ExecutionPayloadHeader.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BlindedBeaconKitBlockBodyCapella object
func (b *BlindedBeaconKitBlockBodyCapella) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlindedBeaconKitBlockBodyCapella object with a hasher
func (b *BlindedBeaconKitBlockBodyCapella) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'RandaoReveal'
	if size := len(b.RandaoReveal); size != 96 {
		err = ssz.ErrBytesLengthFn("--.RandaoReveal", size, 96)
		return
	}
	hh.PutBytes(b.RandaoReveal)

	// Field (1) 'Graffiti'
	if size := len(b.Graffiti); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Graffiti", size, 32)
		return
	}
	hh.PutBytes(b.Graffiti)

	// Field (2) 'Deposits'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Deposits))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Deposits {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 16)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 16)
		}
	}

	// Field (3) 'ExecutionPayloadHeader'
	if err = b.ExecutionPayloadHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
