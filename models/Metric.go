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
