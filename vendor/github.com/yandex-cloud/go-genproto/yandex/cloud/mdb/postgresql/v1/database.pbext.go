// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

func (m *Database) SetName(v string) {
	m.Name = v
}

func (m *Database) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *Database) SetOwner(v string) {
	m.Owner = v
}

func (m *Database) SetLcCollate(v string) {
	m.LcCollate = v
}

func (m *Database) SetLcCtype(v string) {
	m.LcCtype = v
}

func (m *Database) SetExtensions(v []*Extension) {
	m.Extensions = v
}

func (m *Database) SetTemplateDb(v string) {
	m.TemplateDb = v
}

func (m *Extension) SetName(v string) {
	m.Name = v
}

func (m *Extension) SetVersion(v string) {
	m.Version = v
}

func (m *DatabaseSpec) SetName(v string) {
	m.Name = v
}

func (m *DatabaseSpec) SetOwner(v string) {
	m.Owner = v
}

func (m *DatabaseSpec) SetLcCollate(v string) {
	m.LcCollate = v
}

func (m *DatabaseSpec) SetLcCtype(v string) {
	m.LcCtype = v
}

func (m *DatabaseSpec) SetExtensions(v []*Extension) {
	m.Extensions = v
}

func (m *DatabaseSpec) SetTemplateDb(v string) {
	m.TemplateDb = v
}
