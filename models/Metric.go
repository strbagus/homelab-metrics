package models

type DiskType struct {
	Filesystem string `json:"filesystem"`
	Size       int64  `json:"size"`
	Used       int64  `json:"used"`
	Available  int64  `json:"available"`
	Usage      string `json:"usage"`
	Mountpoint string `json:"mountpoint"`
	Hostname   string `json:"hostname"`
}

type ResourcesCategoryType struct {
	Category string   `json:"category"`
	Slug     string   `json:"slug"`
	Types    []string `json:"types"`
}

type ResourcesKindType struct {
	ID        string `json:"id"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Kind      string `json:"kind"`
}
