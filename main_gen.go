package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import "github.com/tinylib/msgp/msgp"

// DecodeMsg implements msgp.Decodable
func (z *Player) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "id":
			z.ID, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
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
func (z Player) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "id"
	err = en.Append(0x82, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.ID)
	if err != nil {
		return
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Player) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "id"
	o = append(o, 0x82, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt(o, z.ID)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Player) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "id":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
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
func (z Player) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PlayerComment) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "id":
			z.ID, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "comment":
			z.Comment, err = dc.ReadString()
			if err != nil {
				return
			}
		case "player_id":
			z.PlayerID, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "created_at":
			z.CreatedAt, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "updated_at":
			z.UpdatedAt, err = dc.ReadInt64()
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
func (z *PlayerComment) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "id"
	err = en.Append(0x85, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.ID)
	if err != nil {
		return
	}
	// write "comment"
	err = en.Append(0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Comment)
	if err != nil {
		return
	}
	// write "player_id"
	err = en.Append(0xa9, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PlayerID)
	if err != nil {
		return
	}
	// write "created_at"
	err = en.Append(0xaa, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.CreatedAt)
	if err != nil {
		return
	}
	// write "updated_at"
	err = en.Append(0xaa, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.UpdatedAt)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PlayerComment) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "id"
	o = append(o, 0x85, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt(o, z.ID)
	// string "comment"
	o = append(o, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Comment)
	// string "player_id"
	o = append(o, 0xa9, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64)
	o = msgp.AppendInt(o, z.PlayerID)
	// string "created_at"
	o = append(o, 0xaa, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74)
	o = msgp.AppendInt64(o, z.CreatedAt)
	// string "updated_at"
	o = append(o, 0xaa, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74)
	o = msgp.AppendInt64(o, z.UpdatedAt)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PlayerComment) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "id":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "comment":
			z.Comment, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "player_id":
			z.PlayerID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "created_at":
			z.CreatedAt, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "updated_at":
			z.UpdatedAt, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *PlayerComment) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 8 + msgp.StringPrefixSize + len(z.Comment) + 10 + msgp.IntSize + 11 + msgp.Int64Size + 11 + msgp.Int64Size
	return
}
