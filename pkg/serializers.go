package sia

import "math/big"

func (s *sia) SerializeInt8ArrayItem(item int8) Sia {
	s.AddInt8(item)
	return s
}

func (s *sia) SerializeInt16ArrayItem(item int16) Sia {
	s.AddInt16(item)
	return s
}

func (s *sia) SerializeInt32ArrayItem(item int32) Sia {
	s.AddInt32(item)
	return s
}

func (s *sia) SerializeInt64ArrayItem(item int64) Sia {
	s.AddInt64(item)
	return s
}

func (s *sia) SerializeUInt8ArrayItem(item uint8) Sia {
	s.AddUInt8(item)
	return s
}

func (s *sia) SerializeUInt16ArrayItem(item uint16) Sia {
	s.AddUInt16(item)
	return s
}

func (s *sia) SerializeUInt32ArrayItem(item uint32) Sia {
	s.AddUInt32(item)
	return s
}

func (s *sia) SerializeUInt64ArrayItem(item uint64) Sia {
	s.AddUInt64(item)
	return s
}

func (s *sia) SerializeString8ArrayItem(item string) Sia {
	s.AddString8(item)
	return s
}

func (s *sia) SerializeString16ArrayItem(item string) Sia {
	s.AddString16(item)
	return s
}

func (s *sia) SerializeString32ArrayItem(item string) Sia {
	s.AddString32(item)
	return s
}

func (s *sia) SerializeString64ArrayItem(item string) Sia {
	s.AddString64(item)
	return s
}

func (s *sia) SerializeByteArray8ArrayItem(item []byte) Sia {
	s.AddByteArray8(item)
	return s
}

func (s *sia) SerializeByteArray16ArrayItem(item []byte) Sia {
	s.AddByteArray16(item)
	return s
}

func (s *sia) SerializeByteArray32ArrayItem(item []byte) Sia {
	s.AddByteArray32(item)
	return s
}

func (s *sia) SerializeByteArray64ArrayItem(item []byte) Sia {
	s.AddByteArray64(item)
	return s
}

func (s *sia) SerializeBoolArrayItem(item bool) Sia {
	s.AddBool(item)
	return s
}

func (s *sia) SerializeBigIntArrayItem(item *big.Int) Sia {
	s.AddBigInt(item)
	return s
}
