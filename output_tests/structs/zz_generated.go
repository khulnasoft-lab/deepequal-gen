//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

// Code generated by deepequal-gen. DO NOT EDIT.

package structs

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Inner) DeepEqual(other *Inner) bool {
	if other == nil {
		return false
	}

	if in.Byte != other.Byte {
		return false
	}
	if in.Int16 != other.Int16 {
		return false
	}
	if in.Int32 != other.Int32 {
		return false
	}
	if in.Int64 != other.Int64 {
		return false
	}
	if in.Uint8 != other.Uint8 {
		return false
	}
	if in.Uint16 != other.Uint16 {
		return false
	}
	if in.Uint32 != other.Uint32 {
		return false
	}
	if in.Uint64 != other.Uint64 {
		return false
	}
	if in.Float32 != other.Float32 {
		return false
	}
	if in.Float64 != other.Float64 {
		return false
	}
	if in.String != other.String {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Ttest) DeepEqual(other *Ttest) bool {
	if other == nil {
		return false
	}

	if in.Inner1 != other.Inner1 {
		return false
	}

	if in.Inner2 != other.Inner2 {
		return false
	}

	return true
}
