package sia

import "math/big"

func (s *sia) DeserializeInt8ArrayItem() int8 {
	return s.ReadInt8()
}

func (s *sia) DeserializeInt16ArrayItem() int16 {
	return s.ReadInt16()
}

func (s *sia) DeserializeInt32ArrayItem() int32 {
	return s.ReadInt32()
}

func (s *sia) DeserializeInt64ArrayItem() int64 {
	return s.ReadInt64()
}

func (s *sia) DeserializeUInt8ArrayItem() uint8 {
	return s.ReadUInt8()
}

func (s *sia) DeserializeUInt16ArrayItem() uint16 {
	return s.ReadUInt16()
}

func (s *sia) DeserializeUInt32ArrayItem() uint32 {
	return s.ReadUInt32()
}

func (s *sia) DeserializeUInt64ArrayItem() uint64 {
	return s.ReadUInt64()
}

func (s *sia) DeserializeString8ArrayItem() string {
	return s.ReadString8()
}

func (s *sia) DeserializeString16ArrayItem() string {
	return s.ReadString16()
}

func (s *sia) DeserializeString32ArrayItem() string {
	return s.ReadString32()
}

func (s *sia) DeserializeString64ArrayItem() string {
	return s.ReadString64()
}

func (s *sia) DeserializeByteArray8ArrayItem() []byte {
	return s.ReadByteArray8()
}

func (s *sia) DeserializeByteArray16ArrayItem() []byte {
	return s.ReadByteArray16()
}

func (s *sia) DeserializeByteArray32ArrayItem() []byte {
	return s.ReadByteArray32()
}

func (s *sia) DeserializeByteArray64ArrayItem() []byte {
	return s.ReadByteArray64()
}

func (s *sia) DeserializeBoolArrayItem() bool {
	return s.ReadBool()
}

func (s *sia) DeserializeBigIntArrayItem() *big.Int {
	return s.ReadBigInt()
}
