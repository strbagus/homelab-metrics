package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type NodeMetric struct {
	Name  string  `json:"name"`
	CUsed float64 `json:"cused"`
	CUnit string  `json:"cunit"`
	CPerc float64 `json:"cperc"`
	MUsed float64 `json:"mused"`
	MUnit string  `json:"munit"`
	MPerc float64 `json:"mperc"`
}

func GetMetric() []NodeMetric {
	cmd := exec.Command("kubectl", "top", "nodes", "--no-headers=true")
	var out bytes.Buffer
	cmd.Stdout = &out

	res := []NodeMetric{}

	if err := cmd.Run(); err != nil {
		fmt.Printf("[ERROR] command failed: %v\n", err)
		return res
	}

	output := strings.TrimSpace(out.String())
	if output == "" {
		return res
	}

	lines := strings.Split(output, "\n")
	res = make([]NodeMetric, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		name := fields[0]
		cpuRaw := fields[1]
		cpuPerRaw := fields[2]
		memRaw := fields[3]
		memPerRaw := fields[4]

		reNum := regexp.MustCompile(`\d+`)
		reUnit := regexp.MustCompile(`[^\d]+`)

		cc, _ := strconv.Atoi(reNum.FindString(cpuRaw))
		cu := reUnit.FindString(cpuRaw)

		cp, _ := strconv.ParseFloat(strings.Replace(cpuPerRaw, "%", "", 1), 64)

		mm, _ := strconv.Atoi(reNum.FindString(memRaw))
		mu := reUnit.FindString(memRaw)

		mp, _ := strconv.ParseFloat(strings.Replace(memPerRaw, "%", "", 1), 64)
		res = append(res, NodeMetric{
			Name:  name,
			CUsed: float64(cc),
			CUnit: cu,
			CPerc: cp,
			MUsed: float64(mm),
			MUnit: mu,
			MPerc: mp,
		})
	}

	return res
}

type Nodes struct {
	Name           string `json:"name"`
	InternalIP     string `json:"internal_ip"`
	CPUs           int16  `json:"cpus"`
	Memory         int32  `json:"memory"`
	MemoryUnit     string `json:"memory_unit"`
	Storage        int64  `json:"storage"`
	StorageUnit    string `json:"storage_unit"`
	Arch           string `json:"arch"`
	OSImage        string `json:"os_image"`
	KernelVersion  string `json:"kernel_version"`
	IsControlPlane bool   `json:"is_control_plane"`
}

type PodKinds struct {
	Kind  string `json:"kind"`
	Count int16  `json:"count"`
}
type PodRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UID  string `json:"uid"`
}
type Pods struct {
	Kind      string   `json:"kind"`
	UID       string   `json:"uid"`
	Namespace string   `json:"namespace"`
	Name      string   `json:"name"`
	App       string   `json:"app"`
	Ref       []PodRef `json:"ref"`
	Status    string   `json:"status"`
	Node      string   `json:"node"`
	Subdomain string   `json:"subdomain"`
	Host      string   `json:"host"`
	Priority  int16    `json:"priority"`
	HostIP    string   `json:"host_ip"`
}

func GetNodes() []Nodes {
	return runJSONCommand[Nodes](CmdGetNodes)
}

func GetPodKinds() []PodKinds {
	return runJSONCommand[PodKinds](CmdGetPodKinds)
}

func GetPods() []Pods {
	return runJSONCommand[Pods](CmdGetPods)
}

func runJSONCommand[T any](cmdStr string) []T {
	result := make([]T, 0)

	cmd := exec.Command("bash", "-c", cmdStr)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("[ERROR] running command %q: %v\n", cmdStr, err)
		return result
	}

	if err := json.Unmarshal(out, &result); err != nil {
		log.Printf("[ERROR] unmarshalling JSON from %q: %v\n", cmdStr, err)
	}

	return result
}
