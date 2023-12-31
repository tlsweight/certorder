// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *User) SetName(v string) {
	m.Name = v
}

func (m *User) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *User) SetPermissions(v []*Permission) {
	m.Permissions = v
}

func (m *User) SetConnLimit(v int64) {
	m.ConnLimit = v
}

func (m *User) SetSettings(v *UserSettings) {
	m.Settings = v
}

func (m *User) SetLogin(v *wrapperspb.BoolValue) {
	m.Login = v
}

func (m *User) SetGrants(v []string) {
	m.Grants = v
}

func (m *Permission) SetDatabaseName(v string) {
	m.DatabaseName = v
}

func (m *UserSpec) SetName(v string) {
	m.Name = v
}

func (m *UserSpec) SetPassword(v string) {
	m.Password = v
}

func (m *UserSpec) SetPermissions(v []*Permission) {
	m.Permissions = v
}

func (m *UserSpec) SetConnLimit(v *wrapperspb.Int64Value) {
	m.ConnLimit = v
}

func (m *UserSpec) SetSettings(v *UserSettings) {
	m.Settings = v
}

func (m *UserSpec) SetLogin(v *wrapperspb.BoolValue) {
	m.Login = v
}

func (m *UserSpec) SetGrants(v []string) {
	m.Grants = v
}

func (m *UserSettings) SetDefaultTransactionIsolation(v UserSettings_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *UserSettings) SetLockTimeout(v *wrapperspb.Int64Value) {
	m.LockTimeout = v
}

func (m *UserSettings) SetLogMinDurationStatement(v *wrapperspb.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *UserSettings) SetSynchronousCommit(v UserSettings_SynchronousCommit) {
	m.SynchronousCommit = v
}

func (m *UserSettings) SetTempFileLimit(v *wrapperspb.Int64Value) {
	m.TempFileLimit = v
}

func (m *UserSettings) SetLogStatement(v UserSettings_LogStatement) {
	m.LogStatement = v
}
