package rascore

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BaseMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "EventName":
			z.EventName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Id":
			z.Id, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "UTCTimestamp":
			z.UTCTimestamp, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z BaseMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "EventName"
	err = en.Append(0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.EventName)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.Id)
	if err != nil {
		return
	}
	// write "UTCTimestamp"
	err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.UTCTimestamp)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BaseMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "EventName"
	o = append(o, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.EventName)
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.Id)
	// string "UTCTimestamp"
	o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.UTCTimestamp)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BaseMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "EventName":
			z.EventName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Id":
			z.Id, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "UTCTimestamp":
			z.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z BaseMessage) Msgsize() (s int) {
	s = 1 + 10 + msgp.StringPrefixSize + len(z.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ChatMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbai uint32
	zbai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "RecipientMessage":
			err = z.RecipientMessage.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "Message":
			z.Message, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ChatMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "RecipientMessage"
	err = en.Append(0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = z.RecipientMessage.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "Message"
	err = en.Append(0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Message)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ChatMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "RecipientMessage"
	o = append(o, 0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o, err = z.RecipientMessage.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "Message"
	o = append(o, 0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Message)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ChatMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "RecipientMessage":
			bts, err = z.RecipientMessage.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "Message":
			z.Message, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ChatMessage) Msgsize() (s int) {
	s = 1 + 17 + z.RecipientMessage.Msgsize() + 8 + msgp.StringPrefixSize + len(z.Message)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ErrorMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zwht uint32
			zwht, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zwht > 0 {
				zwht--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, err = dc.ReadString()
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, err = dc.ReadUint64()
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, err = dc.ReadInt64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "Type":
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Error":
			z.Error, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Body":
			z.Body, err = dc.ReadIntf()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ErrorMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "BaseMessage"
	// map header, size 3
	// write "EventName"
	err = en.Append(0x84, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.BaseMessage.EventName)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.BaseMessage.Id)
	if err != nil {
		return
	}
	// write "UTCTimestamp"
	err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.BaseMessage.UTCTimestamp)
	if err != nil {
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Type)
	if err != nil {
		return
	}
	// write "Error"
	err = en.Append(0xa5, 0x45, 0x72, 0x72, 0x6f, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Error)
	if err != nil {
		return
	}
	// write "Body"
	err = en.Append(0xa4, 0x42, 0x6f, 0x64, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteIntf(z.Body)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ErrorMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "BaseMessage"
	// map header, size 3
	// string "EventName"
	o = append(o, 0x84, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.BaseMessage.EventName)
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.BaseMessage.Id)
	// string "UTCTimestamp"
	o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.BaseMessage.UTCTimestamp)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Error"
	o = append(o, 0xa5, 0x45, 0x72, 0x72, 0x6f, 0x72)
	o = msgp.AppendString(o, z.Error)
	// string "Body"
	o = append(o, 0xa4, 0x42, 0x6f, 0x64, 0x79)
	o, err = msgp.AppendIntf(o, z.Body)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ErrorMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zcua uint32
			zcua, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zcua > 0 {
				zcua--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Error":
			z.Error, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Body":
			z.Body, bts, err = msgp.ReadIntfBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ErrorMessage) Msgsize() (s int) {
	s = 1 + 12 + 1 + 10 + msgp.StringPrefixSize + len(z.BaseMessage.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Type) + 6 + msgp.StringPrefixSize + len(z.Error) + 5 + msgp.GuessSize(z.Body)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *HandshakeMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zlqf uint32
	zlqf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zlqf > 0 {
		zlqf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zdaf uint32
			zdaf, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zdaf > 0 {
				zdaf--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, err = dc.ReadString()
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, err = dc.ReadUint64()
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, err = dc.ReadInt64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "Nick":
			z.Nick, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Rooms":
			var zpks uint32
			zpks, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rooms) >= int(zpks) {
				z.Rooms = (z.Rooms)[:zpks]
			} else {
				z.Rooms = make([]string, zpks)
			}
			for zxhx := range z.Rooms {
				z.Rooms[zxhx], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *HandshakeMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "BaseMessage"
	// map header, size 3
	// write "EventName"
	err = en.Append(0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.BaseMessage.EventName)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.BaseMessage.Id)
	if err != nil {
		return
	}
	// write "UTCTimestamp"
	err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.BaseMessage.UTCTimestamp)
	if err != nil {
		return
	}
	// write "Nick"
	err = en.Append(0xa4, 0x4e, 0x69, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nick)
	if err != nil {
		return
	}
	// write "Rooms"
	err = en.Append(0xa5, 0x52, 0x6f, 0x6f, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Rooms)))
	if err != nil {
		return
	}
	for zxhx := range z.Rooms {
		err = en.WriteString(z.Rooms[zxhx])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *HandshakeMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "BaseMessage"
	// map header, size 3
	// string "EventName"
	o = append(o, 0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.BaseMessage.EventName)
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.BaseMessage.Id)
	// string "UTCTimestamp"
	o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.BaseMessage.UTCTimestamp)
	// string "Nick"
	o = append(o, 0xa4, 0x4e, 0x69, 0x63, 0x6b)
	o = msgp.AppendString(o, z.Nick)
	// string "Rooms"
	o = append(o, 0xa5, 0x52, 0x6f, 0x6f, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rooms)))
	for zxhx := range z.Rooms {
		o = msgp.AppendString(o, z.Rooms[zxhx])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HandshakeMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zcxo uint32
			zcxo, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zcxo > 0 {
				zcxo--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "Nick":
			z.Nick, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Rooms":
			var zeff uint32
			zeff, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rooms) >= int(zeff) {
				z.Rooms = (z.Rooms)[:zeff]
			} else {
				z.Rooms = make([]string, zeff)
			}
			for zxhx := range z.Rooms {
				z.Rooms[zxhx], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *HandshakeMessage) Msgsize() (s int) {
	s = 1 + 12 + 1 + 10 + msgp.StringPrefixSize + len(z.BaseMessage.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Nick) + 6 + msgp.ArrayHeaderSize
	for zxhx := range z.Rooms {
		s += msgp.StringPrefixSize + len(z.Rooms[zxhx])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NickMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrsw uint32
	zrsw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zxpk uint32
			zxpk, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zxpk > 0 {
				zxpk--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, err = dc.ReadString()
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, err = dc.ReadUint64()
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, err = dc.ReadInt64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "OldNick":
			z.OldNick, err = dc.ReadString()
			if err != nil {
				return
			}
		case "NewNick":
			z.NewNick, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *NickMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "BaseMessage"
	// map header, size 3
	// write "EventName"
	err = en.Append(0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.BaseMessage.EventName)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.BaseMessage.Id)
	if err != nil {
		return
	}
	// write "UTCTimestamp"
	err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.BaseMessage.UTCTimestamp)
	if err != nil {
		return
	}
	// write "OldNick"
	err = en.Append(0xa7, 0x4f, 0x6c, 0x64, 0x4e, 0x69, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteString(z.OldNick)
	if err != nil {
		return
	}
	// write "NewNick"
	err = en.Append(0xa7, 0x4e, 0x65, 0x77, 0x4e, 0x69, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteString(z.NewNick)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NickMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "BaseMessage"
	// map header, size 3
	// string "EventName"
	o = append(o, 0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.BaseMessage.EventName)
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.BaseMessage.Id)
	// string "UTCTimestamp"
	o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.BaseMessage.UTCTimestamp)
	// string "OldNick"
	o = append(o, 0xa7, 0x4f, 0x6c, 0x64, 0x4e, 0x69, 0x63, 0x6b)
	o = msgp.AppendString(o, z.OldNick)
	// string "NewNick"
	o = append(o, 0xa7, 0x4e, 0x65, 0x77, 0x4e, 0x69, 0x63, 0x6b)
	o = msgp.AppendString(o, z.NewNick)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NickMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdnj uint32
	zdnj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdnj > 0 {
		zdnj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zobc uint32
			zobc, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zobc > 0 {
				zobc--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "OldNick":
			z.OldNick, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "NewNick":
			z.NewNick, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NickMessage) Msgsize() (s int) {
	s = 1 + 12 + 1 + 10 + msgp.StringPrefixSize + len(z.BaseMessage.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size + 8 + msgp.StringPrefixSize + len(z.OldNick) + 8 + msgp.StringPrefixSize + len(z.NewNick)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PingMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsnv uint32
	zsnv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsnv > 0 {
		zsnv--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zkgt uint32
			zkgt, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zkgt > 0 {
				zkgt--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, err = dc.ReadString()
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, err = dc.ReadUint64()
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, err = dc.ReadInt64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "Type":
			z.Type, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PingMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "BaseMessage"
	// map header, size 3
	// write "EventName"
	err = en.Append(0x82, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.BaseMessage.EventName)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.BaseMessage.Id)
	if err != nil {
		return
	}
	// write "UTCTimestamp"
	err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.BaseMessage.UTCTimestamp)
	if err != nil {
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Type)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PingMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "BaseMessage"
	// map header, size 3
	// string "EventName"
	o = append(o, 0x82, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.BaseMessage.EventName)
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.BaseMessage.Id)
	// string "UTCTimestamp"
	o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.BaseMessage.UTCTimestamp)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, z.Type)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PingMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zema uint32
	zema, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zema > 0 {
		zema--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zpez uint32
			zpez, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zpez > 0 {
				zpez--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "Type":
			z.Type, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PingMessage) Msgsize() (s int) {
	s = 1 + 12 + 1 + 10 + msgp.StringPrefixSize + len(z.BaseMessage.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size + 5 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RecipientContentMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqke uint32
	zqke, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqke > 0 {
		zqke--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "RecipientMessage":
			err = z.RecipientMessage.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "Message":
			z.Message, err = dc.ReadIntf()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RecipientContentMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "RecipientMessage"
	err = en.Append(0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = z.RecipientMessage.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "Message"
	err = en.Append(0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteIntf(z.Message)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RecipientContentMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "RecipientMessage"
	o = append(o, 0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o, err = z.RecipientMessage.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "Message"
	o = append(o, 0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o, err = msgp.AppendIntf(o, z.Message)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RecipientContentMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zqyh uint32
	zqyh, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqyh > 0 {
		zqyh--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "RecipientMessage":
			bts, err = z.RecipientMessage.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "Message":
			z.Message, bts, err = msgp.ReadIntfBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RecipientContentMessage) Msgsize() (s int) {
	s = 1 + 17 + z.RecipientMessage.Msgsize() + 8 + msgp.GuessSize(z.Message)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RecipientMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zyzr uint32
	zyzr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zyzr > 0 {
		zyzr--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zywj uint32
			zywj, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zywj > 0 {
				zywj--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, err = dc.ReadString()
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, err = dc.ReadUint64()
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, err = dc.ReadInt64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "To":
			z.To, err = dc.ReadString()
			if err != nil {
				return
			}
		case "From":
			z.From, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RecipientMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "BaseMessage"
	// map header, size 3
	// write "EventName"
	err = en.Append(0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.BaseMessage.EventName)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.BaseMessage.Id)
	if err != nil {
		return
	}
	// write "UTCTimestamp"
	err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.BaseMessage.UTCTimestamp)
	if err != nil {
		return
	}
	// write "To"
	err = en.Append(0xa2, 0x54, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.To)
	if err != nil {
		return
	}
	// write "From"
	err = en.Append(0xa4, 0x46, 0x72, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.From)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RecipientMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "BaseMessage"
	// map header, size 3
	// string "EventName"
	o = append(o, 0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.BaseMessage.EventName)
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.BaseMessage.Id)
	// string "UTCTimestamp"
	o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.BaseMessage.UTCTimestamp)
	// string "To"
	o = append(o, 0xa2, 0x54, 0x6f)
	o = msgp.AppendString(o, z.To)
	// string "From"
	o = append(o, 0xa4, 0x46, 0x72, 0x6f, 0x6d)
	o = msgp.AppendString(o, z.From)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RecipientMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjpj uint32
	zjpj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zzpf uint32
			zzpf, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zzpf > 0 {
				zzpf--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "To":
			z.To, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "From":
			z.From, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RecipientMessage) Msgsize() (s int) {
	s = 1 + 12 + 1 + 10 + msgp.StringPrefixSize + len(z.BaseMessage.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size + 3 + msgp.StringPrefixSize + len(z.To) + 5 + msgp.StringPrefixSize + len(z.From)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StringMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrfe uint32
	zrfe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrfe > 0 {
		zrfe--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zgmo uint32
			zgmo, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zgmo > 0 {
				zgmo--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, err = dc.ReadString()
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, err = dc.ReadUint64()
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, err = dc.ReadInt64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "Message":
			z.Message, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *StringMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "BaseMessage"
	// map header, size 3
	// write "EventName"
	err = en.Append(0x82, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.BaseMessage.EventName)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.BaseMessage.Id)
	if err != nil {
		return
	}
	// write "UTCTimestamp"
	err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.BaseMessage.UTCTimestamp)
	if err != nil {
		return
	}
	// write "Message"
	err = en.Append(0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Message)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *StringMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "BaseMessage"
	// map header, size 3
	// string "EventName"
	o = append(o, 0x82, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.BaseMessage.EventName)
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.BaseMessage.Id)
	// string "UTCTimestamp"
	o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.BaseMessage.UTCTimestamp)
	// string "Message"
	o = append(o, 0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Message)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StringMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztaf uint32
	ztaf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztaf > 0 {
		ztaf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zeth uint32
			zeth, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zeth > 0 {
				zeth--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "EventName":
					z.BaseMessage.EventName, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "Id":
					z.BaseMessage.Id, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						return
					}
				case "UTCTimestamp":
					z.BaseMessage.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "Message":
			z.Message, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StringMessage) Msgsize() (s int) {
	s = 1 + 12 + 1 + 10 + msgp.StringPrefixSize + len(z.BaseMessage.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size + 8 + msgp.StringPrefixSize + len(z.Message)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *compositeMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsbz uint32
	zsbz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsbz > 0 {
		zsbz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Base":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Base = nil
			} else {
				if z.Base == nil {
					z.Base = new(BaseMessage)
				}
				var zrjx uint32
				zrjx, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zrjx > 0 {
					zrjx--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "EventName":
						z.Base.EventName, err = dc.ReadString()
						if err != nil {
							return
						}
					case "Id":
						z.Base.Id, err = dc.ReadUint64()
						if err != nil {
							return
						}
					case "UTCTimestamp":
						z.Base.UTCTimestamp, err = dc.ReadInt64()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		case "Ping":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Ping = nil
			} else {
				if z.Ping == nil {
					z.Ping = new(PingMessage)
				}
				err = z.Ping.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Handshake":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Handshake = nil
			} else {
				if z.Handshake == nil {
					z.Handshake = new(HandshakeMessage)
				}
				err = z.Handshake.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Recipient":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Recipient = nil
			} else {
				if z.Recipient == nil {
					z.Recipient = new(RecipientMessage)
				}
				err = z.Recipient.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "RecipientContent":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.RecipientContent = nil
			} else {
				if z.RecipientContent == nil {
					z.RecipientContent = new(RecipientContentMessage)
				}
				var zawn uint32
				zawn, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zawn > 0 {
					zawn--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "RecipientMessage":
						err = z.RecipientContent.RecipientMessage.DecodeMsg(dc)
						if err != nil {
							return
						}
					case "Message":
						z.RecipientContent.Message, err = dc.ReadIntf()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		case "Nick":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Nick = nil
			} else {
				if z.Nick == nil {
					z.Nick = new(NickMessage)
				}
				var zwel uint32
				zwel, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zwel > 0 {
					zwel--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "BaseMessage":
						var zrbe uint32
						zrbe, err = dc.ReadMapHeader()
						if err != nil {
							return
						}
						for zrbe > 0 {
							zrbe--
							field, err = dc.ReadMapKeyPtr()
							if err != nil {
								return
							}
							switch msgp.UnsafeString(field) {
							case "EventName":
								z.Nick.BaseMessage.EventName, err = dc.ReadString()
								if err != nil {
									return
								}
							case "Id":
								z.Nick.BaseMessage.Id, err = dc.ReadUint64()
								if err != nil {
									return
								}
							case "UTCTimestamp":
								z.Nick.BaseMessage.UTCTimestamp, err = dc.ReadInt64()
								if err != nil {
									return
								}
							default:
								err = dc.Skip()
								if err != nil {
									return
								}
							}
						}
					case "OldNick":
						z.Nick.OldNick, err = dc.ReadString()
						if err != nil {
							return
						}
					case "NewNick":
						z.Nick.NewNick, err = dc.ReadString()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		case "String":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.String = nil
			} else {
				if z.String == nil {
					z.String = new(StringMessage)
				}
				err = z.String.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Error":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Error = nil
			} else {
				if z.Error == nil {
					z.Error = new(ErrorMessage)
				}
				err = z.Error.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Chat":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Chat = nil
			} else {
				if z.Chat == nil {
					z.Chat = new(ChatMessage)
				}
				var zmfd uint32
				zmfd, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zmfd > 0 {
					zmfd--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "RecipientMessage":
						err = z.Chat.RecipientMessage.DecodeMsg(dc)
						if err != nil {
							return
						}
					case "Message":
						z.Chat.Message, err = dc.ReadString()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *compositeMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "Base"
	err = en.Append(0x89, 0xa4, 0x42, 0x61, 0x73, 0x65)
	if err != nil {
		return err
	}
	if z.Base == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 3
		// write "EventName"
		err = en.Append(0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Base.EventName)
		if err != nil {
			return
		}
		// write "Id"
		err = en.Append(0xa2, 0x49, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteUint64(z.Base.Id)
		if err != nil {
			return
		}
		// write "UTCTimestamp"
		err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.Base.UTCTimestamp)
		if err != nil {
			return
		}
	}
	// write "Ping"
	err = en.Append(0xa4, 0x50, 0x69, 0x6e, 0x67)
	if err != nil {
		return err
	}
	if z.Ping == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Ping.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Handshake"
	err = en.Append(0xa9, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65)
	if err != nil {
		return err
	}
	if z.Handshake == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Handshake.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Recipient"
	err = en.Append(0xa9, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	if z.Recipient == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Recipient.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "RecipientContent"
	err = en.Append(0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	if z.RecipientContent == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 2
		// write "RecipientMessage"
		err = en.Append(0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
		if err != nil {
			return err
		}
		err = z.RecipientContent.RecipientMessage.EncodeMsg(en)
		if err != nil {
			return
		}
		// write "Message"
		err = en.Append(0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteIntf(z.RecipientContent.Message)
		if err != nil {
			return
		}
	}
	// write "Nick"
	err = en.Append(0xa4, 0x4e, 0x69, 0x63, 0x6b)
	if err != nil {
		return err
	}
	if z.Nick == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 3
		// write "BaseMessage"
		// map header, size 3
		// write "EventName"
		err = en.Append(0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Nick.BaseMessage.EventName)
		if err != nil {
			return
		}
		// write "Id"
		err = en.Append(0xa2, 0x49, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteUint64(z.Nick.BaseMessage.Id)
		if err != nil {
			return
		}
		// write "UTCTimestamp"
		err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.Nick.BaseMessage.UTCTimestamp)
		if err != nil {
			return
		}
		// write "OldNick"
		err = en.Append(0xa7, 0x4f, 0x6c, 0x64, 0x4e, 0x69, 0x63, 0x6b)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Nick.OldNick)
		if err != nil {
			return
		}
		// write "NewNick"
		err = en.Append(0xa7, 0x4e, 0x65, 0x77, 0x4e, 0x69, 0x63, 0x6b)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Nick.NewNick)
		if err != nil {
			return
		}
	}
	// write "String"
	err = en.Append(0xa6, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67)
	if err != nil {
		return err
	}
	if z.String == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.String.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Error"
	err = en.Append(0xa5, 0x45, 0x72, 0x72, 0x6f, 0x72)
	if err != nil {
		return err
	}
	if z.Error == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Error.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Chat"
	err = en.Append(0xa4, 0x43, 0x68, 0x61, 0x74)
	if err != nil {
		return err
	}
	if z.Chat == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 2
		// write "RecipientMessage"
		err = en.Append(0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
		if err != nil {
			return err
		}
		err = z.Chat.RecipientMessage.EncodeMsg(en)
		if err != nil {
			return
		}
		// write "Message"
		err = en.Append(0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Chat.Message)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *compositeMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "Base"
	o = append(o, 0x89, 0xa4, 0x42, 0x61, 0x73, 0x65)
	if z.Base == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "EventName"
		o = append(o, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Base.EventName)
		// string "Id"
		o = append(o, 0xa2, 0x49, 0x64)
		o = msgp.AppendUint64(o, z.Base.Id)
		// string "UTCTimestamp"
		o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
		o = msgp.AppendInt64(o, z.Base.UTCTimestamp)
	}
	// string "Ping"
	o = append(o, 0xa4, 0x50, 0x69, 0x6e, 0x67)
	if z.Ping == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Ping.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Handshake"
	o = append(o, 0xa9, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65)
	if z.Handshake == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Handshake.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Recipient"
	o = append(o, 0xa9, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74)
	if z.Recipient == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Recipient.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "RecipientContent"
	o = append(o, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	if z.RecipientContent == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 2
		// string "RecipientMessage"
		o = append(o, 0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
		o, err = z.RecipientContent.RecipientMessage.MarshalMsg(o)
		if err != nil {
			return
		}
		// string "Message"
		o = append(o, 0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
		o, err = msgp.AppendIntf(o, z.RecipientContent.Message)
		if err != nil {
			return
		}
	}
	// string "Nick"
	o = append(o, 0xa4, 0x4e, 0x69, 0x63, 0x6b)
	if z.Nick == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "BaseMessage"
		// map header, size 3
		// string "EventName"
		o = append(o, 0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Nick.BaseMessage.EventName)
		// string "Id"
		o = append(o, 0xa2, 0x49, 0x64)
		o = msgp.AppendUint64(o, z.Nick.BaseMessage.Id)
		// string "UTCTimestamp"
		o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
		o = msgp.AppendInt64(o, z.Nick.BaseMessage.UTCTimestamp)
		// string "OldNick"
		o = append(o, 0xa7, 0x4f, 0x6c, 0x64, 0x4e, 0x69, 0x63, 0x6b)
		o = msgp.AppendString(o, z.Nick.OldNick)
		// string "NewNick"
		o = append(o, 0xa7, 0x4e, 0x65, 0x77, 0x4e, 0x69, 0x63, 0x6b)
		o = msgp.AppendString(o, z.Nick.NewNick)
	}
	// string "String"
	o = append(o, 0xa6, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67)
	if z.String == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.String.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Error"
	o = append(o, 0xa5, 0x45, 0x72, 0x72, 0x6f, 0x72)
	if z.Error == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Error.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Chat"
	o = append(o, 0xa4, 0x43, 0x68, 0x61, 0x74)
	if z.Chat == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 2
		// string "RecipientMessage"
		o = append(o, 0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
		o, err = z.Chat.RecipientMessage.MarshalMsg(o)
		if err != nil {
			return
		}
		// string "Message"
		o = append(o, 0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
		o = msgp.AppendString(o, z.Chat.Message)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *compositeMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzdc uint32
	zzdc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzdc > 0 {
		zzdc--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Base":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Base = nil
			} else {
				if z.Base == nil {
					z.Base = new(BaseMessage)
				}
				var zelx uint32
				zelx, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zelx > 0 {
					zelx--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "EventName":
						z.Base.EventName, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Id":
						z.Base.Id, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							return
						}
					case "UTCTimestamp":
						z.Base.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Ping":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Ping = nil
			} else {
				if z.Ping == nil {
					z.Ping = new(PingMessage)
				}
				bts, err = z.Ping.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Handshake":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Handshake = nil
			} else {
				if z.Handshake == nil {
					z.Handshake = new(HandshakeMessage)
				}
				bts, err = z.Handshake.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Recipient":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Recipient = nil
			} else {
				if z.Recipient == nil {
					z.Recipient = new(RecipientMessage)
				}
				bts, err = z.Recipient.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "RecipientContent":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.RecipientContent = nil
			} else {
				if z.RecipientContent == nil {
					z.RecipientContent = new(RecipientContentMessage)
				}
				var zbal uint32
				zbal, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zbal > 0 {
					zbal--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "RecipientMessage":
						bts, err = z.RecipientContent.RecipientMessage.UnmarshalMsg(bts)
						if err != nil {
							return
						}
					case "Message":
						z.RecipientContent.Message, bts, err = msgp.ReadIntfBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Nick":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Nick = nil
			} else {
				if z.Nick == nil {
					z.Nick = new(NickMessage)
				}
				var zjqz uint32
				zjqz, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zjqz > 0 {
					zjqz--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "BaseMessage":
						var zkct uint32
						zkct, bts, err = msgp.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
						for zkct > 0 {
							zkct--
							field, bts, err = msgp.ReadMapKeyZC(bts)
							if err != nil {
								return
							}
							switch msgp.UnsafeString(field) {
							case "EventName":
								z.Nick.BaseMessage.EventName, bts, err = msgp.ReadStringBytes(bts)
								if err != nil {
									return
								}
							case "Id":
								z.Nick.BaseMessage.Id, bts, err = msgp.ReadUint64Bytes(bts)
								if err != nil {
									return
								}
							case "UTCTimestamp":
								z.Nick.BaseMessage.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
								if err != nil {
									return
								}
							default:
								bts, err = msgp.Skip(bts)
								if err != nil {
									return
								}
							}
						}
					case "OldNick":
						z.Nick.OldNick, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "NewNick":
						z.Nick.NewNick, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "String":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.String = nil
			} else {
				if z.String == nil {
					z.String = new(StringMessage)
				}
				bts, err = z.String.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Error":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Error = nil
			} else {
				if z.Error == nil {
					z.Error = new(ErrorMessage)
				}
				bts, err = z.Error.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Chat":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Chat = nil
			} else {
				if z.Chat == nil {
					z.Chat = new(ChatMessage)
				}
				var ztmt uint32
				ztmt, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for ztmt > 0 {
					ztmt--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "RecipientMessage":
						bts, err = z.Chat.RecipientMessage.UnmarshalMsg(bts)
						if err != nil {
							return
						}
					case "Message":
						z.Chat.Message, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *compositeMessage) Msgsize() (s int) {
	s = 1 + 5
	if z.Base == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 10 + msgp.StringPrefixSize + len(z.Base.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size
	}
	s += 5
	if z.Ping == nil {
		s += msgp.NilSize
	} else {
		s += z.Ping.Msgsize()
	}
	s += 10
	if z.Handshake == nil {
		s += msgp.NilSize
	} else {
		s += z.Handshake.Msgsize()
	}
	s += 10
	if z.Recipient == nil {
		s += msgp.NilSize
	} else {
		s += z.Recipient.Msgsize()
	}
	s += 17
	if z.RecipientContent == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 17 + z.RecipientContent.RecipientMessage.Msgsize() + 8 + msgp.GuessSize(z.RecipientContent.Message)
	}
	s += 5
	if z.Nick == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 12 + 1 + 10 + msgp.StringPrefixSize + len(z.Nick.BaseMessage.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size + 8 + msgp.StringPrefixSize + len(z.Nick.OldNick) + 8 + msgp.StringPrefixSize + len(z.Nick.NewNick)
	}
	s += 7
	if z.String == nil {
		s += msgp.NilSize
	} else {
		s += z.String.Msgsize()
	}
	s += 6
	if z.Error == nil {
		s += msgp.NilSize
	} else {
		s += z.Error.Msgsize()
	}
	s += 5
	if z.Chat == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 17 + z.Chat.RecipientMessage.Msgsize() + 8 + msgp.StringPrefixSize + len(z.Chat.Message)
	}
	return
}
