//go:build !ignore_autogenerated
// +build !ignore_autogenerated

//SPDX-License-Identifier:Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPool) DeepCopyInto(out *AddressPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPool.
func (in *AddressPool) DeepCopy() *AddressPool {
	if in == nil {
		return nil
	}
	out := new(AddressPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPoolList) DeepCopyInto(out *AddressPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddressPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPoolList.
func (in *AddressPoolList) DeepCopy() *AddressPoolList {
	if in == nil {
		return nil
	}
	out := new(AddressPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPoolSpec) DeepCopyInto(out *AddressPoolSpec) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AutoAssign != nil {
		in, out := &in.AutoAssign, &out.AutoAssign
		*out = new(bool)
		**out = **in
	}
	if in.BGPAdvertisements != nil {
		in, out := &in.BGPAdvertisements, &out.BGPAdvertisements
		*out = make([]BgpAdvertisement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPoolSpec.
func (in *AddressPoolSpec) DeepCopy() *AddressPoolSpec {
	if in == nil {
		return nil
	}
	out := new(AddressPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPoolStatus) DeepCopyInto(out *AddressPoolStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPoolStatus.
func (in *AddressPoolStatus) DeepCopy() *AddressPoolStatus {
	if in == nil {
		return nil
	}
	out := new(AddressPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BgpAdvertisement) DeepCopyInto(out *BgpAdvertisement) {
	*out = *in
	if in.AggregationLength != nil {
		in, out := &in.AggregationLength, &out.AggregationLength
		*out = new(int32)
		**out = **in
	}
	if in.AggregationLengthV6 != nil {
		in, out := &in.AggregationLengthV6, &out.AggregationLengthV6
		*out = new(int32)
		**out = **in
	}
	if in.Communities != nil {
		in, out := &in.Communities, &out.Communities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BgpAdvertisement.
func (in *BgpAdvertisement) DeepCopy() *BgpAdvertisement {
	if in == nil {
		return nil
	}
	out := new(BgpAdvertisement)
	in.DeepCopyInto(out)
	return out
}