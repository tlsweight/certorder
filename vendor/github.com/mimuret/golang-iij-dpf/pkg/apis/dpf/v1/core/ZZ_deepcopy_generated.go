//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
MIT License

Copyright (c) 2021 Manabu Sonoda

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package core

import (
	api "github.com/mimuret/golang-iij-dpf/pkg/api"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttributeMeta) DeepCopyInto(out *AttributeMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttributeMeta.
func (in *AttributeMeta) DeepCopy() *AttributeMeta {
	if in == nil {
		return nil
	}
	out := new(AttributeMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Contract) DeepCopyInto(out *Contract) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Contract.
func (in *Contract) DeepCopy() *Contract {
	if in == nil {
		return nil
	}
	out := new(Contract)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Contract) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContractList) DeepCopyInto(out *ContractList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Contract, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContractList.
func (in *ContractList) DeepCopy() *ContractList {
	if in == nil {
		return nil
	}
	out := new(ContractList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *ContractList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Delegation) DeepCopyInto(out *Delegation) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	in.DelegationRequestedAt.DeepCopyInto(&out.DelegationRequestedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Delegation.
func (in *Delegation) DeepCopy() *Delegation {
	if in == nil {
		return nil
	}
	out := new(Delegation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Delegation) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DelegationApply) DeepCopyInto(out *DelegationApply) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.ZoneIDs != nil {
		in, out := &in.ZoneIDs, &out.ZoneIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DelegationApply.
func (in *DelegationApply) DeepCopy() *DelegationApply {
	if in == nil {
		return nil
	}
	out := new(DelegationApply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *DelegationApply) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DelegationList) DeepCopyInto(out *DelegationList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Delegation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DelegationList.
func (in *DelegationList) DeepCopy() *DelegationList {
	if in == nil {
		return nil
	}
	out := new(DelegationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *DelegationList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Job) DeepCopyInto(out *Job) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Job.
func (in *Job) DeepCopy() *Job {
	if in == nil {
		return nil
	}
	out := new(Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Job) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LBDomain) DeepCopyInto(out *LBDomain) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LBDomain.
func (in *LBDomain) DeepCopy() *LBDomain {
	if in == nil {
		return nil
	}
	out := new(LBDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *LBDomain) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LBDomainList) DeepCopyInto(out *LBDomainList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LBDomain, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LBDomainList.
func (in *LBDomainList) DeepCopy() *LBDomainList {
	if in == nil {
		return nil
	}
	out := new(LBDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *LBDomainList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Log) DeepCopyInto(out *Log) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Log.
func (in *Log) DeepCopy() *Log {
	if in == nil {
		return nil
	}
	out := new(Log)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Zone) DeepCopyInto(out *Zone) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Zone.
func (in *Zone) DeepCopy() *Zone {
	if in == nil {
		return nil
	}
	out := new(Zone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Zone) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneList) DeepCopyInto(out *ZoneList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Zone, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneList.
func (in *ZoneList) DeepCopy() *ZoneList {
	if in == nil {
		return nil
	}
	out := new(ZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *ZoneList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
