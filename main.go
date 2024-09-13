package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	server        = "console-api-stg.fptcloud.net"
	vpcId         = "4ef9197e-90a9-4ce4-8df0-bb172db0a720"
	bearerToken   = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE3MjU5MzkzNzkuNTE1MDgyMSwic3ViIjp7ImlkIjoiZDc5YjEwOWUtZWE4MS00ZmE0LWI3NzEtNjZlZWU0ZWJjYjZkIiwiZW1haWwiOiJsb2NucDI1QGZwdC5jb20udm4iLCJqdGkiOiIyZWQxNWUzMy1lOWMzLTRjMjctYWI0ZS01ZGEyY2FiZTgwOTAifSwiZXhwIjoxNzI2MDM3Mzc5LjUxNTA4MjF9.qmUAwYoV8FuP6BqOcyN0RX9mnrM3c2nnxy_v19OrXl8"
	serverType    = "646204d4-1995-484a-b469-277619022d22"
	scriptContent = `#!/bin/bash
mkdir -p /etc/cloud/cloud.cfg.d/
cat <<EOF > /etc/cloud/cloud.cfg.d/custom-networking.cfg
network:
config: disabled
EOF
chmod 0644 /etc/cloud/cloud.cfg.d/custom-networking.cfg


mkdir -p '/var/lib/cloud-config-downloader/credentials'
cat << EOF | base64 -d > '/var/lib/cloud-config-downloader/credentials/server'
aHR0cHM6Ly9hcGkuY2hlY2stZHBhYXMtdXZlZnZ6N28ueHBsYXQtdmRjLmludGVybmFsLnN0Zy5tLmZrZS5mcHRjbG91ZC5jb20=
EOF
chmod '0644' '/var/lib/cloud-config-downloader/credentials/server'

mkdir -p '/var/lib/cloud-config-downloader/credentials'
cat << EOF | base64 -d > '/var/lib/cloud-config-downloader/credentials/ca.crt'
LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM1akNDQWM2Z0F3SUJBZ0lRTnVVZTI4c2tCYXM2SXE0OHdrTVZSREFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWVGdzB5TkRBNU1UQXdPRFEwTURCYUZ3MHpOREE1TVRBd09EUTBNREJhTUEweApDekFKQmdOVkJBTVRBbU5oTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF5MG02ClcvdjBxSG9wbDc2ZDhTNUVYd29ublBLRzNqZ3VyOTZHbGxxQmFEcDJpS3lsWktmYzA3NFhnZGZEdGZBNjVQeFoKKzM5Y3pSNFg0ZGNzUkhLbHpXTFd3YzhzSG5xSlc2aWIrS01wckM3WFI0b1BTT1pGSjhJa0w3TkdCeHkvU25TbwpMN2dyNXNxbFUzOUlCWUdGZU83SUpSeHB6ZWFVM29zaUZCSXUzSmx6czVyT2VHQVhOdXZPSW94TkFab2EzbmNKCm10d2djTUd1YUN4aVZiYzYwZThOOTlKRFRkNmZTb1dqOUR1TUo5Mm95dUVTYmo2WkdPVzl5SzNUTS9lZFJmUXgKMlUrbUkzdFdacnpmazFZQnd5bm9BTjMxNXlCS3BQUEduTi8vb0xEKzRYUllpWUw2ZHBsbVFhSHNiQmpDNjlhYQp1Z2xWUEs5YW13RE83UFNnWVFJREFRQUJvMEl3UURBT0JnTlZIUThCQWY4RUJBTUNBYVl3RHdZRFZSMFRBUUgvCkJBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVWJTY1RrTHBic0tHRWphOFBOTnJpVEdvSjRFQXdEUVlKS29aSWh2Y04KQVFFTEJRQURnZ0VCQUxZb1hOMENCMXZ0TDdhVnBVMTdUK0prZmU1MEZ6dEJ1R1lzR1RuQWp5QlVkRDJYeFdpMQpiRTFmU3Jmem15OUVmRnorUnc4Z2RjNmJ6UjJPekJkZm5wK1FQZUQ4dXZjcWlneDIvVzFuYlBpVXV3eFc0dlRDCjVuMy9IdG52MzFCMWljQzhtZUhZV1MwektrcGVXUTBIVm1jSndjVk1iMi9oRXhaMEx5cTlEMjFzSUkyQmdjdFcKK0QrakhrWkxsNjRBSWdVNFBSclloRnc3eGMra2s0bUp2TER4MkNQU2JUMGxqZDZiYTNyYjVES2NMN2NFT3RLTwpiMzF6enR6K0V6andGMUo4VCsxcVJwR1R4THZSbUZqL2h6VzJjR0xvYlFYY2JoY2V5UStyMXhnNXdETFdNSE4vCkZaZVVnMEJrK1FQTnV5Mk4rSlBHczl1ZysrVVZmZkFTMEVrPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
EOF
chmod '0644' '/var/lib/cloud-config-downloader/credentials/ca.crt'

mkdir -p '/var/lib/cloud-config-downloader'
cat << EOF | base64 -d > '/var/lib/cloud-config-downloader/download-cloud-config.sh'
IyEvYmluL2Jhc2gKCnNldCAtbyBlcnJleGl0CnNldCAtbyBub3Vuc2V0CnNldCAtbyBwaXBlZmFpbAoKewpTRUNSRVRfTkFNRT0iY2xvdWQtY29uZmlnLXdvcmtlci1zcDA5cDlmci1mNzZiNiIKVE9LRU5fU0VDUkVUX05BTUU9ImNsb3VkLWNvbmZpZy1kb3dubG9hZGVyIgoKUEFUSF9DTE9VRENPTkZJR19ET1dOTE9BREVSX0NMSUVOVF9DRVJUPSIvdmFyL2xpYi9jbG91ZC1jb25maWctZG93bmxvYWRlci9jcmVkZW50aWFscy9jbGllbnQuY3J0IgpQQVRIX0NMT1VEQ09ORklHX0RPV05MT0FERVJfQ0xJRU5UX0tFWT0iL3Zhci9saWIvY2xvdWQtY29uZmlnLWRvd25sb2FkZXIvY3JlZGVudGlhbHMvY2xpZW50LmtleSIKUEFUSF9DTE9VRENPTkZJR19ET1dOTE9BREVSX1RPS0VOPSIvdmFyL2xpYi9jbG91ZC1jb25maWctZG93bmxvYWRlci9jcmVkZW50aWFscy90b2tlbiIKUEFUSF9CT09UU1RSQVBfVE9LRU49Ii92YXIvbGliL2Nsb3VkLWNvbmZpZy1kb3dubG9hZGVyL2NyZWRlbnRpYWxzL2Jvb3RzdHJhcC10b2tlbiIKUEFUSF9FWEVDVVRPUl9TQ1JJUFQ9Ii92YXIvbGliL2Nsb3VkLWNvbmZpZy1kb3dubG9hZGVyL2Rvd25sb2Fkcy9leGVjdXRlLWNsb3VkLWNvbmZpZy5zaCIKUEFUSF9FWEVDVVRPUl9TQ1JJUFRfQ0hFQ0tTVU09Ii92YXIvbGliL2Nsb3VkLWNvbmZpZy1kb3dubG9hZGVyL2Rvd25sb2Fkcy9leGVjdXRlLWNsb3VkLWNvbmZpZy1jaGVja3N1bSIKCm1rZGlyIC1wICIvdmFyL2xpYi9jbG91ZC1jb25maWctZG93bmxvYWRlci9kb3dubG9hZHMiCgpmdW5jdGlvbiByZWFkU2VjcmV0KCkgewogIHdnZXQgXAogICAgLXFPLSBcCiAgICAtLWNhLWNlcnRpZmljYXRlICIvdmFyL2xpYi9jbG91ZC1jb25maWctZG93bmxvYWRlci9jcmVkZW50aWFscy9jYS5jcnQiIFwKICAgICIke0A6Mn0iICIkKGNhdCAiL3Zhci9saWIvY2xvdWQtY29uZmlnLWRvd25sb2FkZXIvY3JlZGVudGlhbHMvc2VydmVyIikvYXBpL3YxL25hbWVzcGFjZXMva3ViZS1zeXN0ZW0vc2VjcmV0cy8kMSIKfQoKZnVuY3Rpb24gcmVhZFNlY3JldEZ1bGwoKSB7CiAgcmVhZFNlY3JldCAiJDEiICItLWhlYWRlcj1BY2NlcHQ6IGFwcGxpY2F0aW9uL3lhbWwiICIke0A6Mn0iCn0KCmZ1bmN0aW9uIHJlYWRTZWNyZXRNZXRhKCkgewogIHJlYWRTZWNyZXQgIiQxIiAiLS1oZWFkZXI9QWNjZXB0OiBhcHBsaWNhdGlvbi95YW1sO2FzPVBhcnRpYWxPYmplY3RNZXRhZGF0YTtnPW1ldGEuazhzLmlvO3Y9djEsYXBwbGljYXRpb24veWFtbDthcz1QYXJ0aWFsT2JqZWN0TWV0YWRhdGE7Zz1tZXRhLms4cy5pbzt2PXYxIiAiJHtAOjJ9Igp9CgpmdW5jdGlvbiByZWFkU2VjcmV0TWV0YVdpdGhUb2tlbigpIHsKICByZWFkU2VjcmV0TWV0YSAiJDEiICItLWhlYWRlcj1BdXRob3JpemF0aW9uOiBCZWFyZXIgJDIiCn0KCmZ1bmN0aW9uIHJlYWRTZWNyZXRXaXRoVG9rZW4oKSB7CiAgcmVhZFNlY3JldEZ1bGwgIiQxIiAiLS1oZWFkZXI9QXV0aG9yaXphdGlvbjogQmVhcmVyICQyIgp9CgpmdW5jdGlvbiByZWFkU2VjcmV0V2l0aENsaWVudENlcnRpZmljYXRlKCkgewogIHJlYWRTZWNyZXRGdWxsICIkMSIgIi0tY2VydGlmaWNhdGU9JFBBVEhfQ0xPVURDT05GSUdfRE9XTkxPQURFUl9DTElFTlRfQ0VSVCIgIi0tcHJpdmF0ZS1rZXk9JFBBVEhfQ0xPVURDT05GSUdfRE9XTkxPQURFUl9DTElFTlRfS0VZIgp9CgpmdW5jdGlvbiBleHRyYWN0RGF0YUtleUZyb21TZWNyZXQoKSB7CiAgZWNobyAiJDEiIHwgc2VkIC1ybiAicy8gICQyOiAoLiopL1wxL3AiIHwgYmFzZTY0IC1kCn0KCmZ1bmN0aW9uIGV4dHJhY3RDaGVja3N1bUZyb21TZWNyZXQoKSB7CiAgZWNobyAiJDEiIHwgc2VkIC1ybiAncy8gICAgY2hlY2tzdW1cL2RhdGEtc2NyaXB0OiAoLiopL1wxL3AnIHwgc2VkIC1lICdzL14iLy8nIC1lICdzLyIkLy8nCn0KCmZ1bmN0aW9uIHdyaXRlVG9EaXNrU2FmZWx5KCkgewogIGxvY2FsIGRhdGE9IiQxIgogIGxvY2FsIGZpbGVfcGF0aD0iJDIiCgogIGlmIGVjaG8gIiRkYXRhIiA+ICIkZmlsZV9wYXRoLnRtcCIgJiYgKCBbWyAhIC1mICIkZmlsZV9wYXRoIiBdXSB8fCAhIGRpZmYgIiRmaWxlX3BhdGgiICIkZmlsZV9wYXRoLnRtcCI+L2Rldi9udWxsICk7IHRoZW4KICAgIG12ICIkZmlsZV9wYXRoLnRtcCIgIiRmaWxlX3BhdGgiCiAgZWxpZiBbWyAtZiAiJGZpbGVfcGF0aC50bXAiIF1dOyB0aGVuCiAgICBybSAtZiAiJGZpbGVfcGF0aC50bXAiCiAgZmkKfQoKIyBkb3dubG9hZCBzaG9vdCBhY2Nlc3MgdG9rZW4gZm9yIGNsb3VkLWNvbmZpZy1kb3dubG9hZGVyCmlmIFtbIC1mICIkUEFUSF9DTE9VRENPTkZJR19ET1dOTE9BREVSX1RPS0VOIiBdXTsgdGhlbgogIGlmICEgU0VDUkVUPSIkKHJlYWRTZWNyZXRXaXRoVG9rZW4gIiRUT0tFTl9TRUNSRVRfTkFNRSIgIiQoY2F0ICIkUEFUSF9DTE9VRENPTkZJR19ET1dOTE9BREVSX1RPS0VOIikiKSI7IHRoZW4KICAgIGVjaG8gIkNvdWxkIG5vdCByZXRyaWV2ZSB0aGUgc2hvb3QgYWNjZXNzIHNlY3JldCB3aXRoIG5hbWUgJFRPS0VOX1NFQ1JFVF9OQU1FIHdpdGggZXhpc3RpbmcgdG9rZW4iCiAgICBleGl0IDEKICBmaQplbHNlCiAgaWYgW1sgLWYgIiRQQVRIX0JPT1RTVFJBUF9UT0tFTiIgXV07IHRoZW4KICAgIGlmICEgU0VDUkVUPSIkKHJlYWRTZWNyZXRXaXRoVG9rZW4gIiRUT0tFTl9TRUNSRVRfTkFNRSIgIiQoY2F0ICIkUEFUSF9CT09UU1RSQVBfVE9LRU4iKSIpIjsgdGhlbgogICAgICBlY2hvICJDb3VsZCBub3QgcmV0cmlldmUgdGhlIHNob290IGFjY2VzcyBzZWNyZXQgd2l0aCBuYW1lICRUT0tFTl9TRUNSRVRfTkFNRSB3aXRoIGJvb3RzdHJhcCB0b2tlbiIKICAgICAgZXhpdCAxCiAgICBmaQogIGVsc2UKICAgIGlmICEgU0VDUkVUPSIkKHJlYWRTZWNyZXRXaXRoQ2xpZW50Q2VydGlmaWNhdGUgIiRUT0tFTl9TRUNSRVRfTkFNRSIpIjsgdGhlbgogICAgICBlY2hvICJDb3VsZCBub3QgcmV0cmlldmUgdGhlIHNob290IGFjY2VzcyBzZWNyZXQgd2l0aCBuYW1lICRUT0tFTl9TRUNSRVRfTkFNRSB3aXRoIGNsaWVudCBjZXJ0aWZpY2F0ZSIKICAgICAgZXhpdCAxCiAgICBmaQogIGZpCmZpCgpUT0tFTj0iJChleHRyYWN0RGF0YUtleUZyb21TZWNyZXQgIiRTRUNSRVQiICJ0b2tlbiIpIgppZiBbWyAteiAiJFRPS0VOIiBdXTsgdGhlbgogIGVjaG8gIlRva2VuIGluIHNob290IGFjY2VzcyBzZWNyZXQgJFRPS0VOX1NFQ1JFVF9OQU1FIGlzIGVtcHR5IgogIGV4aXQgMQpmaQp3cml0ZVRvRGlza1NhZmVseSAiJFRPS0VOIiAiJFBBVEhfQ0xPVURDT05GSUdfRE9XTkxPQURFUl9UT0tFTiIKCiMgZG93bmxvYWQgYW5kIHJ1biB0aGUgY2xvdWQgY29uZmlnIGV4ZWN1dGlvbiBzY3JpcHQKaWYgISBTRUNSRVRfTUVUQT0iJChyZWFkU2VjcmV0TWV0YVdpdGhUb2tlbiAiJFNFQ1JFVF9OQU1FIiAiJFRPS0VOIikiOyB0aGVuCiAgZWNobyAiQ291bGQgbm90IHJldHJpZXZlIHRoZSBtZXRhZGF0YSBpbiBzZWNyZXQgd2l0aCBuYW1lICRTRUNSRVRfTkFNRSIKICBleGl0IDEKZmkKTkVXX0NIRUNLU1VNPSIkKGV4dHJhY3RDaGVja3N1bUZyb21TZWNyZXQgIiRTRUNSRVRfTUVUQSIpIgoKT0xEX0NIRUNLU1VNPSI8bm9uZT4iCmlmIFtbIC1mICIkUEFUSF9FWEVDVVRPUl9TQ1JJUFRfQ0hFQ0tTVU0iIF1dOyB0aGVuCiAgT0xEX0NIRUNLU1VNPSIkKGNhdCAiJFBBVEhfRVhFQ1VUT1JfU0NSSVBUX0NIRUNLU1VNIikiCmZpCgppZiBbWyAiJE5FV19DSEVDS1NVTSIgIT0gIiRPTERfQ0hFQ0tTVU0iIF1dOyB0aGVuCiAgZWNobyAiQ2hlY2tzdW0gb2YgY2xvdWQgY29uZmlnIHNjcmlwdCBoYXMgY2hhbmdlZCBjb21wYXJlZCB0byB3aGF0IEkgaGFkIGRvd25sb2FkZWQgZWFybGllciAobmV3OiAkTkVXX0NIRUNLU1VNLCBvbGQ6ICRPTERfQ0hFQ0tTVU0pLiBGZXRjaGluZyBuZXcgc2NyaXB0Li4uIgoKICBpZiAhIFNFQ1JFVD0iJChyZWFkU2VjcmV0V2l0aFRva2VuICIkU0VDUkVUX05BTUUiICIkVE9LRU4iKSI7IHRoZW4KICAgIGVjaG8gIkNvdWxkIG5vdCByZXRyaWV2ZSB0aGUgY2xvdWQgY29uZmlnIHNjcmlwdCBpbiBzZWNyZXQgd2l0aCBuYW1lICRTRUNSRVRfTkFNRSIKICAgIGV4aXQgMQogIGZpCgogIFNDUklQVD0iJChleHRyYWN0RGF0YUtleUZyb21TZWNyZXQgIiRTRUNSRVQiICJzY3JpcHQiKSIKICBpZiBbWyAteiAiJFNDUklQVCIgXV07IHRoZW4KICAgIGVjaG8gIlNjcmlwdCBpbiBjbG91ZCBjb25maWcgc2VjcmV0ICRTRUNSRVQgaXMgZW1wdHkiCiAgICBleGl0IDEKICBmaQoKICB3cml0ZVRvRGlza1NhZmVseSAiJFNDUklQVCIgIiRQQVRIX0VYRUNVVE9SX1NDUklQVCIgJiYgY2htb2QgK3ggIiRQQVRIX0VYRUNVVE9SX1NDUklQVCIKICB3cml0ZVRvRGlza1NhZmVseSAiJChleHRyYWN0Q2hlY2tzdW1Gcm9tU2VjcmV0ICIkU0VDUkVUIikiICIkUEFUSF9FWEVDVVRPUl9TQ1JJUFRfQ0hFQ0tTVU0iCmZpCgoiJFBBVEhfRVhFQ1VUT1JfU0NSSVBUIgpleGl0ICQ/Cn0K
EOF
chmod '0744' '/var/lib/cloud-config-downloader/download-cloud-config.sh'

mkdir -p '/var/lib/cloud-config-downloader/credentials'
cat << EOF > '/var/lib/cloud-config-downloader/credentials/bootstrap-token'
6c7238.o3kvzscw95znbjfk
EOF
chmod '0644' '/var/lib/cloud-config-downloader/credentials/bootstrap-token'

cat << EOF | base64 -d > '/etc/systemd/system/cloud-config-downloader.service'
W1VuaXRdCkRlc2NyaXB0aW9uPURvd25sb2FkcyB0aGUgYWN0dWFsIGNsb3VkIGNvbmZpZyBmcm9tIHRoZSBTaG9vdCBBUEkgc2VydmVyIGFuZCBleGVjdXRlcyBpdApBZnRlcj1kb2NrZXIuc2VydmljZSBkb2NrZXIuc29ja2V0CldhbnRzPWRvY2tlci5zb2NrZXQKW1NlcnZpY2VdClJlc3RhcnQ9YWx3YXlzClJlc3RhcnRTZWM9MzAKUnVudGltZU1heFNlYz0xMjAwCkVudmlyb25tZW50RmlsZT0vZXRjL2Vudmlyb25tZW50CkV4ZWNTdGFydD0vdmFyL2xpYi9jbG91ZC1jb25maWctZG93bmxvYWRlci9kb3dubG9hZC1jbG91ZC1jb25maWcuc2gKW0luc3RhbGxdCldhbnRlZEJ5PW11bHRpLXVzZXIudGFyZ2V0
EOF

until apt-get update -qq && apt-get install --no-upgrade -qqy containerd runc docker.io socat nfs-common logrotate jq policykit-1; do sleep 1; done
ln -s /usr/bin/docker /bin/docker
if [ ! -s /etc/containerd/config.toml ]; then
mkdir -p /etc/containerd/
containerd config default > /etc/containerd/config.toml
chmod 0644 /etc/containerd/config.toml
fi

mkdir -p /etc/systemd/system/containerd.service.d
cat <<EOF > /etc/systemd/system/containerd.service.d/11-exec_config.conf
[Service]
ExecStart=
ExecStart=/usr/bin/containerd --config=/etc/containerd/config.toml
EOF
chmod 0644 /etc/systemd/system/containerd.service.d/11-exec_config.conf

systemctl daemon-reload
systemctl enable containerd && systemctl restart containerd
systemctl enable docker && systemctl restart docker
systemctl enable cloud-config-downloader && systemctl restart cloud-config-downloader
rm -rf /root/.bash_history
history -c
`
)

type requestBodyCreate struct {
	RegionId   string   `json:"regionId,omitempty"`  // Region id (hanoi-vn)
	ClusterId  int32    `json:"clusterId,omitempty"` // Cluster id
	ServerType string   `json:"serverType"`          // Server type ID as uuid
	Names      []string `json:"names"`               // Define name of machines, it will be understanded as number of machines
	OS         string   `json:"os"`                  // OS type that machines will be installed (ubuntu/centos..)
	RaidType   string   `json:"raidType,omitempty"`  // Raid type that machine will be configured as a storage type (level-1) (Not using right now)
	SshKey     string   `json:"sshKey,omitempty"`    // sshKey that user already imported on portal, it must be an identify sshkey on portal
	UserData   string   `json:"userData,omitempty"`  // post user defines script after os installed
	Distro     string   `json:"distro,omitempty"`    // Json dump with keys name and hwe_kernel. Name as distro name, hwe_kernel as architecture
}

type sshkey struct {
	Id         string `json:"id"`         // ID of sshKey - required
	Name       string `json:"name"`       // Name of sshKey to display in portal - required (Default: fke-bm-sshKey)
	Public_key string `json:"public_key"` // PublicKey is imported to server
}

type distro struct {
	Name       string `json:"name"`       // Name of distro
	Hwe_kernel string `json:"hwe_kernel"` // architecture
}

type requestBodyPut struct { // Update a machine (Turn ON/OFF machine)
	Name  string `json:"name,omitempty"`  // Name of machine that will be update (Null means no modification)
	Power string `json:"power,omitempty"` // Power type of mechine that will be applied (Null means no modification) [ on, off, null ]
}
type responseBodyCreateAccepted struct { // Created accepted (202)
	Data  string `json:"data"`
	Error string `json:"message,omitempty"`
}

type responseBodyCreateError struct { // BadRequest (400) Unauthorized (401) InternalServerError (500) NotFound (404)
	// IsSuccess bool   `json:"isSuccess"`
	Error string `json:"message,omitempty"` // (optional) Reasons
}

type responseBodyList struct { // Update/Get/Delete a machine successfully  (200)  BadRequest (400) Unauthorized (401) InternalServerError (500) NotFound (404)
	Data      []serverDetail `json:"data"`
	Total     int            `json:"total"`
	IsSuccess bool           `json:"isSuccess"`
	Error     string         `json:"error,omitempty"` // (optional) Reasons
}

type responseBodyDetail struct {
	Data    serverDetail `json:"data"`
	Message string       `json:"message"`
}

type responseBodyDelete struct {
	Data    serverDetail `json:"data"`
	Message string       `json:"message"`
	Status  bool         `json:"status"`
}

type serverDetail struct {
	Id              string `json:"id"`              // Server's id
	Status          string `json:"status"`          // Server's status
	Created_at      string `json:"created_at"`      // Server's created time
	Updated_at      string `json:"updated_at"`      // Server's updated time
	Name            string `json:"name"`            // Name of server
	Vpc_id          string `json:"vpc_id"`          // Vpc id
	Hpc_pool_id     int    `json:"hpc_pool_id"`     // Hpc Pool id
	Pool_name       string `json:"pool_name"`       // pool name that server belong to
	Hpc_server_type string `json:"hpc_server_type"` // Server Type / Resource type
	Region          string `json:"region"`          // Server's region
	Os              string `json:"os"`              // Installed OS
	Subnet          string `json:"subnet"`
	Private_ip      string `json:"private_ip"` // Private IP address
	Public_ip       string `json:"public_ip"`  // Public IP address
	Meta_data       string `json:"meta_data"`  // Server meta data
}

func main() {
	bash := fmt.Sprintf(`#cloud-config
user: xplat
password: sondx12@123
chpasswd: {expire: False}
ssh_pwauth: True
hostname: sondx12
manage_etc_hosts: true
write_files:
  - path: /root/custom_network_config.sh
    permissions: '0755'
    content: |
      %s

runcmd:
  - /bin/bash /root/custom_network_config.sh`, scriptContent)
	fmt.Printf("%s", bash)
	// err := ListServer(server, vpcId, bearerToken)
	// if err != nil {
	// 	fmt.Printf("error when listing server in vpc: [%v]", err)
	// }

	PublicKey := sshkey{
		Id:         "aaaa-bbbb-cccc-dddd",
		Name:       "fke-pass",
		Public_key: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCjPnuMpUNzG5cztSjEFTC29SHU9RDJ/QKDlpfbtOsl6D/9sIdVoH1QaC98d6IUnumogfQvPqP38dvvau7lkHpp6+stfGsfLBm6HpUHeNTD11YdIddUJOZpnkFghOc2TIpTTnFvw4FYEPDGE0zHqsmhcWLPHfjyG9FJTd4pUinE6/V9BGQ/9O4zZTo+XNajoolewNK6giTWk1L86/SNoFVE9leZs2g85EYOztuP+U0t8kSoMVxXJygi2WuuAonQxNIIMn2WbaGevLISu93kACQelO+lYDAW2T4NEL2l26CYgQGOUGUq/9bU/6EG4WL7u7kd9b6T5sGaNvvS8M0wth9f",
	}
	jsonDataKey, err := json.Marshal(PublicKey)
	if err != nil {
		fmt.Printf("error: [%v]", err)
	}
	Distro := distro{
		Name:       "ubuntu/jammy",
		Hwe_kernel: "amd64/hwe-22.04",
	}

	jsonDataDistro, err := json.Marshal(Distro)
	if err != nil {
		fmt.Printf("error: [%v]", err)
	}
	fmt.Printf("sshKey: %s\n", string(jsonDataKey))
	fmt.Printf("Distro: %s\n", string(jsonDataDistro))
	// 	reqCreate := requestBodyCreate{
	// 		RegionId:   "hanoi-vn",
	// 		Names:      []string{"sondx12"},
	// 		RaidType:   "",
	// 		ServerType: serverType,
	// 		OS:         "ubuntu",
	// 		SshKey:     string(jsonDataKey),
	// 		Distro:     string(jsonDataDistro),
	// 		ClusterId:  2,
	// 		UserData: `
	// #!/bin/bash

	// # New hostname
	// NEW_HOSTNAME="sondx12"

	// # Change the hostname in /etc/hostname
	// echo "$NEW_HOSTNAME" > /etc/hostname

	// # Update /etc/hosts to reflect the new hostname
	// sed -i "s/127.0.1.1 .*/127.0.1.1 $NEW_HOSTNAME/g" /etc/hosts

	// # Set the new hostname for the current session
	// hostnamectl set-hostname "$NEW_HOSTNAME"

	// # Print a message
	// echo "Hostname successfully changed to $NEW_HOSTNAME. A reboot may be required for full effect."
	// reboot`,
	// 	}

	// 	err = CreateMachine(server, vpcId, bearerToken, reqCreate)
	// 	if err != nil {
	// 		fmt.Printf("error when get detail server in vpc: [%v]", err)
	// 	}

	// err = GetMachine(server, vpcId, bearerToken, "665a5173-a48c-42ee-9415-788725a581a7")
	// if err != nil {
	// 	fmt.Printf("error when get detail server in vpc: [%v]", err)
	// }

	// err = DeleteMachine(server, vpcId, bearerToken, "c32aaad1-91c5-4c88-b780-f357cb9be90d")
	// if err != nil {
	// 	fmt.Printf("error when delete server in vpc: [%v]", err)
	// }
	// err = PowerOffServer(server, vpcId, bearerToken, "3e935fca-9082-45b0-9da2-ce9b01af40ab")
	// if err != nil {
	// 	fmt.Printf("error when power off server in vpc: [%v]", err)
	// }
}

// Create a machine with user's predefination
func CreateMachine(server, vpcId, bearerToken string, reqCreate requestBodyCreate) error {
	portalServer := "https://" + server + "/api/v1/vmware/vpc/" + vpcId + "/hpc/server/create"
	bearer := "Bearer " + bearerToken
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	createBody, _ := json.Marshal(reqCreate)
	reqBody := bytes.NewBuffer(createBody)
	clienthttp := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", portalServer, reqBody)
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	resp, err := clienthttp.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform POST request : [%v]", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body of GET placement policy request: [%v]", err)
	}
	if resp.StatusCode == 202 {
		var result responseBodyCreateAccepted
		json.Unmarshal(body, &result)
		fmt.Printf("Data: %s\n", result.Data)
		fmt.Printf("Error: %s\n", result.Error)
		return nil
	}
	var result responseBodyCreateError
	json.Unmarshal(body, &result)
	fmt.Printf("Error: %v\n", resp.StatusCode)
	fmt.Printf("Message: %s\n", string(body))
	return fmt.Errorf("error to create machine: [%v]", result.Error)
}

// PowerOff server BM
func PowerOffServer(server, vpcId, bearerToken, systemId string) error {
	portalServer := "https://" + server + "/api/v1/vmware/vpc/" + vpcId + "/hpc/server/" + systemId + "/powered-off"
	bearer := "Bearer " + bearerToken
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// putBody, _ := json.Marshal(reqPut)
	// reqBody := bytes.NewBuffer(putBody)
	clienthttp := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", portalServer, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := clienthttp.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform POST request : [%v]", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body of POST to poweroff a server state: [%v]", err)
	}
	var result responseBodyCreateAccepted
	json.Unmarshal(body, &result)
	if resp.StatusCode == 202 {
		fmt.Printf("data: %s\n", result.Data)
		fmt.Printf("message: %s\n", result.Error)
		return nil
	}
	fmt.Printf("data: %s\n", result.Data)
	// fmt.Printf("message: %s\n", result.Error)
	return fmt.Errorf("error to update machine: [%v]", result.Error)
}

// PowerOff server BM
func PowerOnServer(server, vpcId, bearerToken, systemId string) error {
	portalServer := "https://" + server + "/api/v1/vmware/vpc/" + vpcId + "/hpc/server/" + systemId + "/powered-on"
	bearer := "Bearer " + bearerToken
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clienthttp := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", portalServer, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := clienthttp.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform POST request : [%v]", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body of POST to poweron a server state: [%v]", err)
	}
	var result responseBodyCreateAccepted
	json.Unmarshal(body, &result)
	if resp.StatusCode == 202 {
		fmt.Printf("data: %s\n", result.Data)
		fmt.Printf("message: %s\n", result.Error)
		return nil
	}
	fmt.Printf("data: %s\n", result.Data)
	// fmt.Printf("message: %s\n", result.Error)
	return fmt.Errorf("error to update machine: [%v]", result.Error)
}

// Get machine detail with machine ID
func GetMachine(server, vpcId, bearerToken, systemId string) error {
	portalServer := "https://" + server + "/api/v1/vmware/vpc/" + vpcId + "/hpc/server/" + systemId + "/detail"
	bearer := "Bearer " + bearerToken
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clienthttp := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", portalServer, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := clienthttp.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform GET request : [%v]", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body of GET machine detail request : [%v]", err)
	}
	var result responseBodyDetail
	json.Unmarshal(body, &result)
	if resp.StatusCode == 200 {
		fmt.Printf("Error: %s\n", result.Message)
		fmt.Printf("data: %v\n", result.Data)
		return nil
	}
	fmt.Printf("Error: %s\n", result.Message)
	return fmt.Errorf("error to GET machine: [%v]", result.Message)
}

// Release a machine with machine ID
func DeleteMachine(server, vpcId, bearerToken, systemId string) error {
	portalServer := "https://" + server + "/api/v1/vmware/vpc/" + vpcId + "/hpc/server/" + systemId + "/delete"
	bearer := "Bearer " + bearerToken
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clienthttp := &http.Client{Transport: tr}
	req, _ := http.NewRequest("DELETE", portalServer, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := clienthttp.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform DELETE request : [%v]", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body of DELETE machine request : [%v]", err)
	}
	var result responseBodyDelete
	json.Unmarshal(body, &result)
	if resp.StatusCode == 202 {
		fmt.Printf("Message: %s\n", result.Message)
		fmt.Printf("Data: %v\n", result.Data)
		return nil
	}
	fmt.Printf("Error: %s\n", result.Message)
	// fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
	return fmt.Errorf("error to DELETE machine: [%v]", result.Message)
}

// List all server in VPC
func ListServer(server, vpcID, bearerToken string) error {
	portalServer := "https://" + server + "/api/v1/vmware/vpc/" + vpcId + "/hpc/server/list"
	bearer := "Bearer " + bearerToken
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clienthttp := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", portalServer, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := clienthttp.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform GET request : [%v]", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body of list machine detail request : [%v]", err)
	}
	var result responseBodyList
	json.Unmarshal(body, &result)
	if resp.StatusCode == 200 {
		if result.Total > 0 {
			fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
			fmt.Printf("Error: %s\n", result.Error)
			fmt.Printf("body: %v\n", result.Data[0])
		} else {
			fmt.Printf("No server found in vpc")
		}
		return nil
	}
	fmt.Printf("Error: %s\n", result.Error)
	fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
	return fmt.Errorf("error to list machine: [%v]", result.Error)
}
