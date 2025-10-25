package utils

var CmdGetNodes = `kubectl get nodes -o json | jq '.items | map({name: .metadata.annotations["k3s.io/hostname"], internal_ip: .metadata.annotations["k3s.io/internal-ip"], cpus: (.status.capacity.cpu | tonumber), memory: (.status.capacity.memory | match("\\d+").string | tonumber), memory_unit: (.status.capacity.memory | sub("[0-9]"; ""; "g")), storage: (.status.capacity["ephemeral-storage"] | match("\\d+").string | tonumber), storage_unit: (.status.capacity["ephemeral-storage"] | sub("[0-9]"; ""; "g")), arch: .status.nodeInfo.architecture, os_image: .status.nodeInfo.osImage, kernel_version: .status.nodeInfo.kernelVersion, is_control_plane: (.metadata.labels["node-role.kubernetes.io/control-plane"] | not | not) })'`

var CmdGetPodKinds = `kubectl get all -A -o json | jq '.items | map(select(.metadata.namespace != "kube-system") | {kind: .kind}) | group_by(.kind) | map({ kind: .[0].kind, count: length})'`

var CmdGetPods = `kubectl get all -A -o json | jq '.items | map(select(.metadata.namespace != "kube-system") | {kind: .kind, uid: .metadata.uid, namespace: .metadata.namespace, name: .metadata.name, app: .metadata.labels.app, ref: .metadata.ownerReferences, status: .spec.phase, node: .spec.nodeName, subdomain: .spec.subdomain, host: .spec.hostname, priority: .spec.priority, status: .status.phase, host_ip: .status.hostIP })'`
