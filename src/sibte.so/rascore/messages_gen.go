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
			var zcmr uint32
			zcmr, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zcmr > 0 {
				zcmr--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "BaseMessage":
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
						case "EventName":
							z.RecipientMessage.BaseMessage.EventName, err = dc.ReadString()
							if err != nil {
								return
							}
						case "Id":
							z.RecipientMessage.BaseMessage.Id, err = dc.ReadUint64()
							if err != nil {
								return
							}
						case "UTCTimestamp":
							z.RecipientMessage.BaseMessage.UTCTimestamp, err = dc.ReadInt64()
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
					z.RecipientMessage.To, err = dc.ReadString()
					if err != nil {
						return
					}
				case "From":
					z.RecipientMessage.From, err = dc.ReadString()
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
func (z *ChatMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "RecipientMessage"
	// map header, size 3
	// write "BaseMessage"
	// map header, size 3
	// write "EventName"
	err = en.Append(0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.RecipientMessage.BaseMessage.EventName)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.RecipientMessage.BaseMessage.Id)
	if err != nil {
		return
	}
	// write "UTCTimestamp"
	err = en.Append(0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.RecipientMessage.BaseMessage.UTCTimestamp)
	if err != nil {
		return
	}
	// write "To"
	err = en.Append(0xa2, 0x54, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.RecipientMessage.To)
	if err != nil {
		return
	}
	// write "From"
	err = en.Append(0xa4, 0x46, 0x72, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.RecipientMessage.From)
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
	// map header, size 3
	// string "BaseMessage"
	// map header, size 3
	// string "EventName"
	o = append(o, 0x82, 0xb0, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xab, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x83, 0xa9, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.RecipientMessage.BaseMessage.EventName)
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.RecipientMessage.BaseMessage.Id)
	// string "UTCTimestamp"
	o = append(o, 0xac, 0x55, 0x54, 0x43, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.RecipientMessage.BaseMessage.UTCTimestamp)
	// string "To"
	o = append(o, 0xa2, 0x54, 0x6f)
	o = msgp.AppendString(o, z.RecipientMessage.To)
	// string "From"
	o = append(o, 0xa4, 0x46, 0x72, 0x6f, 0x6d)
	o = msgp.AppendString(o, z.RecipientMessage.From)
	// string "Message"
	o = append(o, 0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Message)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ChatMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "RecipientMessage":
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
							z.RecipientMessage.BaseMessage.EventName, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "Id":
							z.RecipientMessage.BaseMessage.Id, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								return
							}
						case "UTCTimestamp":
							z.RecipientMessage.BaseMessage.UTCTimestamp, bts, err = msgp.ReadInt64Bytes(bts)
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
					z.RecipientMessage.To, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "From":
					z.RecipientMessage.From, bts, err = msgp.ReadStringBytes(bts)
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
func (z *ChatMessage) Msgsize() (s int) {
	s = 1 + 17 + 1 + 12 + 1 + 10 + msgp.StringPrefixSize + len(z.RecipientMessage.BaseMessage.EventName) + 3 + msgp.Uint64Size + 13 + msgp.Int64Size + 3 + msgp.StringPrefixSize + len(z.RecipientMessage.To) + 5 + msgp.StringPrefixSize + len(z.RecipientMessage.From) + 8 + msgp.StringPrefixSize + len(z.Message)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ErrorMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
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
	var zdaf uint32
	zdaf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdaf > 0 {
		zdaf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zpks uint32
			zpks, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zpks > 0 {
				zpks--
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
	var zcxo uint32
	zcxo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zeff uint32
			zeff, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zeff > 0 {
				zeff--
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
			var zrsw uint32
			zrsw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rooms) >= int(zrsw) {
				z.Rooms = (z.Rooms)[:zrsw]
			} else {
				z.Rooms = make([]string, zrsw)
			}
			for zjfb := range z.Rooms {
				z.Rooms[zjfb], err = dc.ReadString()
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
	for zjfb := range z.Rooms {
		err = en.WriteString(z.Rooms[zjfb])
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
	for zjfb := range z.Rooms {
		o = msgp.AppendString(o, z.Rooms[zjfb])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HandshakeMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxpk uint32
	zxpk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxpk > 0 {
		zxpk--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
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
			var zobc uint32
			zobc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rooms) >= int(zobc) {
				z.Rooms = (z.Rooms)[:zobc]
			} else {
				z.Rooms = make([]string, zobc)
			}
			for zjfb := range z.Rooms {
				z.Rooms[zjfb], bts, err = msgp.ReadStringBytes(bts)
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
	for zjfb := range z.Rooms {
		s += msgp.StringPrefixSize + len(z.Rooms[zjfb])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NickMessage) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "BaseMessage":
			var zqyh uint32
			zqyh, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zqyh > 0 {
				zqyh--
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
	var zyzr uint32
	zyzr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zyzr > 0 {
		zyzr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zywj uint32
			zywj, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zywj > 0 {
				zywj--
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
	var zjpj uint32
	zjpj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
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
		case "BaseMessage":
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
	var zawn uint32
	zawn, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zawn > 0 {
		zawn--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "BaseMessage":
			var zwel uint32
			zwel, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zwel > 0 {
				zwel--
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
