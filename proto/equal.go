// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"bytes"
	"math"
	"reflect"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Equal reports whether two messages are equal,
// by recursively comparing the fields of the message.
//
//   - Bytes fields are equal if they contain identical bytes.
//     Empty bytes (regardless of nil-ness) are considered equal.
//
//   - Floating-point fields are equal if they contain the same value.
//     Unlike the == operator, a NaN is equal to another NaN.
//
//   - Other scalar fields are equal if they contain the same value.
//
//   - Message fields are equal if they have
//     the same set of populated known and extension field values, and
//     the same set of unknown fields values.
//
//   - Lists are equal if they are the same length and
//     each corresponding element is equal.
//
//   - Maps are equal if they have the same set of keys and
//     the corresponding value for each key is equal.
//
// An invalid message is not equal to a valid message.
// An invalid message is only equal to another invalid message of the
// same type. An invalid message often corresponds to a nil pointer
// of the concrete message type. For example, (*pb.M)(nil) is not equal
// to &pb.M{}.
// If two valid messages marshal to the same bytes under deterministic
// serialization, then Equal is guaranteed to report true.
func Equal(x, y Message) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}
	if reflect.TypeOf(x).Kind() == reflect.Ptr && x == y {
		// Avoid an expensive comparison if both inputs are identical pointers.
		return true
	}
	mx := x.ProtoReflect()
	my := y.ProtoReflect()
	if mx.IsValid() != my.IsValid() {
		return false
	}
	return equalMessage(mx, my)
}

// equalMessage compares two messages.
func equalMessage(mx, my protoreflect.Message) bool {
	if mx.Descriptor() != my.Descriptor() {
		return false
	}

	nx := 0
	equal := true
	mx.Range(func(fd protoreflect.FieldDescriptor, vx protoreflect.Value) bool {
		nx++
		vy := my.Get(fd)
		equal = my.Has(fd) && equalField(fd, vx, vy)
		return equal
	})
	if !equal {
		return false
	}
	ny := 0
	my.Range(func(fd protoreflect.FieldDescriptor, vx protoreflect.Value) bool {
		ny++
		return true
	})
	if nx != ny {
		return false
	}

	return equalUnknown(mx.GetUnknown(), my.GetUnknown())
}

// equalField compares two fields.
func equalField(fd protoreflect.FieldDescriptor, x, y protoreflect.Value) bool {
	switch {
	case fd.IsList():
		return equalList(fd, x.List(), y.List())
	case fd.IsMap():
		return equalMap(fd, x.Map(), y.Map())
	default:
		return equalValue(fd, x, y)
	}
}

// equalMap compares two maps.
func equalMap(fd protoreflect.FieldDescriptor, x, y protoreflect.Map) bool {
	if x.Len() != y.Len() {
		return false
	}
	equal := true
	x.Range(func(k protoreflect.MapKey, vx protoreflect.Value) bool {
		vy := y.Get(k)
		equal = y.Has(k) && equalValue(fd.MapValue(), vx, vy)
		return equal
	})
	return equal
}

// equalList compares two lists.
func equalList(fd protoreflect.FieldDescriptor, x, y protoreflect.List) bool {
	if x.Len() != y.Len() {
		return false
	}
	for i := x.Len() - 1; i >= 0; i-- {
		if !equalValue(fd, x.Get(i), y.Get(i)) {
			return false
		}
	}
	return true
}

// equalValue compares two singular values.
func equalValue(fd protoreflect.FieldDescriptor, x, y protoreflect.Value) bool {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		return x.Bool() == y.Bool()
	case protoreflect.EnumKind:
		return x.Enum() == y.Enum()
		case protoreflect.Int32Kind, protoreflect.Sint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return x.Int() == y.Int()
		case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		return x.Uint() == y.Uint()
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		fx := x.Float()
		fy := y.Float()
		if math.IsNaN(fx) || math.IsNaN(fy) {
			return math.IsNaN(fx) && math.IsNaN(fy)
		}
		return fx == fy
	case protoreflect.StringKind:
		if psx, ok := x.Interface().(protoreflect.ProtoStringer); ok {
			sx, err := psx.ProtoString()
			if err != nil {
				panic(err)
			}
			sy, err := y.Interface().(protoreflect.ProtoStringer).ProtoString()
			if err != nil {
				panic(err)
			}
			return sx == sy
		}
		return x.String() == y.String()
	case protoreflect.BytesKind:
		return bytes.Equal(x.Bytes(), y.Bytes())
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return equalMessage(x.Message(), y.Message())
	default:
		return x.Interface() == y.Interface()
	}
}

// equalUnknown compares unknown fields by direct comparison on the raw bytes
// of each individual field number.
func equalUnknown(x, y protoreflect.RawFields) bool {
	if len(x) != len(y) {
		return false
	}
	if bytes.Equal([]byte(x), []byte(y)) {
		return true
	}

	mx := make(map[protoreflect.FieldNumber]protoreflect.RawFields)
	my := make(map[protoreflect.FieldNumber]protoreflect.RawFields)
	for len(x) > 0 {
		fnum, _, n := protowire.ConsumeField(x)
		mx[fnum] = append(mx[fnum], x[:n]...)
		x = x[n:]
	}
	for len(y) > 0 {
		fnum, _, n := protowire.ConsumeField(y)
		my[fnum] = append(my[fnum], y[:n]...)
		y = y[n:]
	}
	return reflect.DeepEqual(mx, my)
}
