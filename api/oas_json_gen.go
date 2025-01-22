// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Item) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Item) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("shortDescription")
		e.Str(s.ShortDescription)
	}
	{
		e.FieldStart("price")
		e.Str(s.Price)
	}
}

var jsonFieldsNameOfItem = [2]string{
	0: "shortDescription",
	1: "price",
}

// Decode decodes Item from json.
func (s *Item) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Item to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "shortDescription":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.ShortDescription = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"shortDescription\"")
			}
		case "price":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Price = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"price\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Item")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfItem) {
					name = jsonFieldsNameOfItem[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Item) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Item) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int64 as json.
func (o OptInt64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int64(int64(o.Value))
}

// Decode decodes int64 from json.
func (o *OptInt64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt64 to nil")
	}
	o.Set = true
	v, err := d.Int64()
	if err != nil {
		return err
	}
	o.Value = int64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Receipt) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Receipt) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("retailer")
		e.Str(s.Retailer)
	}
	{
		e.FieldStart("purchaseDate")
		json.EncodeDate(e, s.PurchaseDate)
	}
	{
		e.FieldStart("purchaseTime")
		json.EncodeTime(e, s.PurchaseTime)
	}
	{
		e.FieldStart("items")
		e.ArrStart()
		for _, elem := range s.Items {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("total")
		e.Str(s.Total)
	}
}

var jsonFieldsNameOfReceipt = [5]string{
	0: "retailer",
	1: "purchaseDate",
	2: "purchaseTime",
	3: "items",
	4: "total",
}

// Decode decodes Receipt from json.
func (s *Receipt) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Receipt to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "retailer":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Retailer = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"retailer\"")
			}
		case "purchaseDate":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := json.DecodeDate(d)
				s.PurchaseDate = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"purchaseDate\"")
			}
		case "purchaseTime":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := json.DecodeTime(d)
				s.PurchaseTime = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"purchaseTime\"")
			}
		case "items":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				s.Items = make([]Item, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Item
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Items = append(s.Items, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"items\"")
			}
		case "total":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Str()
				s.Total = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Receipt")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfReceipt) {
					name = jsonFieldsNameOfReceipt[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Receipt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Receipt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ReceiptsIDPointsGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ReceiptsIDPointsGetOK) encodeFields(e *jx.Encoder) {
	{
		if s.Points.Set {
			e.FieldStart("points")
			s.Points.Encode(e)
		}
	}
}

var jsonFieldsNameOfReceiptsIDPointsGetOK = [1]string{
	0: "points",
}

// Decode decodes ReceiptsIDPointsGetOK from json.
func (s *ReceiptsIDPointsGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReceiptsIDPointsGetOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "points":
			if err := func() error {
				s.Points.Reset()
				if err := s.Points.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"points\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ReceiptsIDPointsGetOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReceiptsIDPointsGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReceiptsIDPointsGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ReceiptsProcessPostOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ReceiptsProcessPostOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Str(s.ID)
	}
}

var jsonFieldsNameOfReceiptsProcessPostOK = [1]string{
	0: "id",
}

// Decode decodes ReceiptsProcessPostOK from json.
func (s *ReceiptsProcessPostOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReceiptsProcessPostOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.ID = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ReceiptsProcessPostOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfReceiptsProcessPostOK) {
					name = jsonFieldsNameOfReceiptsProcessPostOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReceiptsProcessPostOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReceiptsProcessPostOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
