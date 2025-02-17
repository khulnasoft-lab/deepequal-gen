//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

// Code generated by deepequal-gen. DO NOT EDIT.

package maps

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Ttest) DeepEqual(other *Ttest) bool {
	if other == nil {
		return false
	}

	if ((in.Byte != nil) && (other.Byte != nil)) || ((in.Byte == nil) != (other.Byte == nil)) {
		in, other := &in.Byte, &other.Byte
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Int16 != nil) && (other.Int16 != nil)) || ((in.Int16 == nil) != (other.Int16 == nil)) {
		in, other := &in.Int16, &other.Int16
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Int32 != nil) && (other.Int32 != nil)) || ((in.Int32 == nil) != (other.Int32 == nil)) {
		in, other := &in.Int32, &other.Int32
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Int64 != nil) && (other.Int64 != nil)) || ((in.Int64 == nil) != (other.Int64 == nil)) {
		in, other := &in.Int64, &other.Int64
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Uint8 != nil) && (other.Uint8 != nil)) || ((in.Uint8 == nil) != (other.Uint8 == nil)) {
		in, other := &in.Uint8, &other.Uint8
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Uint16 != nil) && (other.Uint16 != nil)) || ((in.Uint16 == nil) != (other.Uint16 == nil)) {
		in, other := &in.Uint16, &other.Uint16
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Uint32 != nil) && (other.Uint32 != nil)) || ((in.Uint32 == nil) != (other.Uint32 == nil)) {
		in, other := &in.Uint32, &other.Uint32
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Uint64 != nil) && (other.Uint64 != nil)) || ((in.Uint64 == nil) != (other.Uint64 == nil)) {
		in, other := &in.Uint64, &other.Uint64
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Float32 != nil) && (other.Float32 != nil)) || ((in.Float32 == nil) != (other.Float32 == nil)) {
		in, other := &in.Float32, &other.Float32
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Float64 != nil) && (other.Float64 != nil)) || ((in.Float64 == nil) != (other.Float64 == nil)) {
		in, other := &in.Float64, &other.Float64
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.String != nil) && (other.String != nil)) || ((in.String == nil) != (other.String == nil)) {
		in, other := &in.String, &other.String
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.StringPtr != nil) && (other.StringPtr != nil)) || ((in.StringPtr == nil) != (other.StringPtr == nil)) {
		in, other := &in.StringPtr, &other.StringPtr
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if (inValue == nil) != (otherValue == nil) || ((inValue != nil) && (otherValue != nil) && (*inValue != *otherValue)) {
						return false
					}
				}
			}
		}
	}

	if ((in.Struct != nil) && (other.Struct != nil)) || ((in.Struct == nil) != (other.Struct == nil)) {
		in, other := &in.Struct, &other.Struct
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(&otherValue) {
						return false
					}
				}
			}
		}
	}

	if ((in.StructPtr != nil) && (other.StructPtr != nil)) || ((in.StructPtr == nil) != (other.StructPtr == nil)) {
		in, other := &in.StructPtr, &other.StructPtr
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(otherValue) {
						return false
					}
				}
			}
		}
	}

	return true
}
