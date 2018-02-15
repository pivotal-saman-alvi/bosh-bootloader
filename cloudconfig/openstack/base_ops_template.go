package openstack

const (
	BaseOps = `---
- type: replace
  path: /azs
  value:
  - name: z1
    cloud_properties:
      availability_zone: ((az))
  - name: z2
    cloud_properties:
      availability_zone: ((az))
  - name: z3
    cloud_properties:
      availability_zone: ((az))

- type: replace
  path: /compilation
  value:
    workers: 5
    reuse_compilation_vms: true
    az: z1
    vm_type: default
    network: default

- type: replace
  path: /disk_types/name=default/disk_size?
  value: 3000

- type: replace
  path: /networks
  value:
  - name: default
    type: manual
    subnets:
    - range: ((internal_cidr))
      gateway: ((internal_gw))
      azs: [z1, z2, z3]
      dns: [8.8.8.8]
      reserved: [((jumpbox__internal_ip))]
      cloud_properties:
        net_id: ((net_id))

- type: replace
  path: /vm_extensions/-
  value:
    name: cf-router-network-properties

- type: replace
  path: /vm_extensions/-
  value:
    name: cf-tcp-router-network-properties

- type: replace
  path: /vm_extensions/-
  value:
    name: diego-ssh-proxy-network-properties
`
)