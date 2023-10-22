// Code generated by protoc-gen-goext. DO NOT EDIT.

package cdn

func (m *Origin) SetId(v int64) {
	m.Id = v
}

func (m *Origin) SetOriginGroupId(v int64) {
	m.OriginGroupId = v
}

func (m *Origin) SetSource(v string) {
	m.Source = v
}

func (m *Origin) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *Origin) SetBackup(v bool) {
	m.Backup = v
}

func (m *Origin) SetMeta(v *OriginMeta) {
	m.Meta = v
}

func (m *OriginParams) SetSource(v string) {
	m.Source = v
}

func (m *OriginParams) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *OriginParams) SetBackup(v bool) {
	m.Backup = v
}

func (m *OriginParams) SetMeta(v *OriginMeta) {
	m.Meta = v
}

type OriginMeta_OriginMetaVariant = isOriginMeta_OriginMetaVariant

func (m *OriginMeta) SetOriginMetaVariant(v OriginMeta_OriginMetaVariant) {
	m.OriginMetaVariant = v
}

func (m *OriginMeta) SetCommon(v *OriginNamedMeta) {
	m.OriginMetaVariant = &OriginMeta_Common{
		Common: v,
	}
}

func (m *OriginMeta) SetBucket(v *OriginNamedMeta) {
	m.OriginMetaVariant = &OriginMeta_Bucket{
		Bucket: v,
	}
}

func (m *OriginMeta) SetWebsite(v *OriginNamedMeta) {
	m.OriginMetaVariant = &OriginMeta_Website{
		Website: v,
	}
}

func (m *OriginMeta) SetBalancer(v *OriginBalancerMeta) {
	m.OriginMetaVariant = &OriginMeta_Balancer{
		Balancer: v,
	}
}

func (m *OriginNamedMeta) SetName(v string) {
	m.Name = v
}

func (m *OriginBalancerMeta) SetId(v string) {
	m.Id = v
}
