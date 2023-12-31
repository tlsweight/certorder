// Code generated by protoc-gen-goext. DO NOT EDIT.

package billing

func (m *ListCustomersRequest) SetResellerId(v string) {
	m.ResellerId = v
}

func (m *ListCustomersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListCustomersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListCustomersResponse) SetCustomers(v []*Customer) {
	m.Customers = v
}

func (m *ListCustomersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *InviteCustomerRequest) SetResellerId(v string) {
	m.ResellerId = v
}

func (m *InviteCustomerRequest) SetName(v string) {
	m.Name = v
}

func (m *InviteCustomerRequest) SetInvitationEmail(v string) {
	m.InvitationEmail = v
}

func (m *InviteCustomerRequest) SetPerson(v *CustomerPerson) {
	m.Person = v
}

func (m *ActivateCustomerRequest) SetCustomerId(v string) {
	m.CustomerId = v
}

func (m *SuspendCustomerRequest) SetCustomerId(v string) {
	m.CustomerId = v
}

func (m *CustomerMetadata) SetResellerId(v string) {
	m.ResellerId = v
}

func (m *CustomerMetadata) SetCustomerId(v string) {
	m.CustomerId = v
}
