// Code generated by protoc-gen-goext. DO NOT EDIT.

package sqlserver

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *SQLServerConfig2017Std) SetMaxDegreeOfParallelism(v *wrapperspb.Int64Value) {
	m.MaxDegreeOfParallelism = v
}

func (m *SQLServerConfig2017Std) SetCostThresholdForParallelism(v *wrapperspb.Int64Value) {
	m.CostThresholdForParallelism = v
}

func (m *SQLServerConfig2017Std) SetAuditLevel(v *wrapperspb.Int64Value) {
	m.AuditLevel = v
}

func (m *SQLServerConfig2017Std) SetFillFactorPercent(v *wrapperspb.Int64Value) {
	m.FillFactorPercent = v
}

func (m *SQLServerConfig2017Std) SetOptimizeForAdHocWorkloads(v *wrapperspb.BoolValue) {
	m.OptimizeForAdHocWorkloads = v
}

func (m *SQLServerConfigSet2017Std) SetEffectiveConfig(v *SQLServerConfig2017Std) {
	m.EffectiveConfig = v
}

func (m *SQLServerConfigSet2017Std) SetUserConfig(v *SQLServerConfig2017Std) {
	m.UserConfig = v
}

func (m *SQLServerConfigSet2017Std) SetDefaultConfig(v *SQLServerConfig2017Std) {
	m.DefaultConfig = v
}

func (m *SQLServerConfig2017Ent) SetMaxDegreeOfParallelism(v *wrapperspb.Int64Value) {
	m.MaxDegreeOfParallelism = v
}

func (m *SQLServerConfig2017Ent) SetCostThresholdForParallelism(v *wrapperspb.Int64Value) {
	m.CostThresholdForParallelism = v
}

func (m *SQLServerConfig2017Ent) SetAuditLevel(v *wrapperspb.Int64Value) {
	m.AuditLevel = v
}

func (m *SQLServerConfig2017Ent) SetFillFactorPercent(v *wrapperspb.Int64Value) {
	m.FillFactorPercent = v
}

func (m *SQLServerConfig2017Ent) SetOptimizeForAdHocWorkloads(v *wrapperspb.BoolValue) {
	m.OptimizeForAdHocWorkloads = v
}

func (m *SQLServerConfigSet2017Ent) SetEffectiveConfig(v *SQLServerConfig2017Ent) {
	m.EffectiveConfig = v
}

func (m *SQLServerConfigSet2017Ent) SetUserConfig(v *SQLServerConfig2017Ent) {
	m.UserConfig = v
}

func (m *SQLServerConfigSet2017Ent) SetDefaultConfig(v *SQLServerConfig2017Ent) {
	m.DefaultConfig = v
}
