// Code generated by protoc-gen-goext. DO NOT EDIT.

package compute

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type PlacementGroup_PlacementStrategy = isPlacementGroup_PlacementStrategy

func (m *PlacementGroup) SetPlacementStrategy(v PlacementGroup_PlacementStrategy) {
	m.PlacementStrategy = v
}

func (m *PlacementGroup) SetId(v string) {
	m.Id = v
}

func (m *PlacementGroup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *PlacementGroup) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *PlacementGroup) SetName(v string) {
	m.Name = v
}

func (m *PlacementGroup) SetDescription(v string) {
	m.Description = v
}

func (m *PlacementGroup) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *PlacementGroup) SetSpreadPlacementStrategy(v *SpreadPlacementStrategy) {
	m.PlacementStrategy = &PlacementGroup_SpreadPlacementStrategy{
		SpreadPlacementStrategy: v,
	}
}
