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
	server      = "console-api-stg.fptcloud.net"
	vpcId       = "4ef9197e-90a9-4ce4-8df0-bb172db0a720"
	bearerToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE3MjU5MzkzNzkuNTE1MDgyMSwic3ViIjp7ImlkIjoiZDc5YjEwOWUtZWE4MS00ZmE0LWI3NzEtNjZlZWU0ZWJjYjZkIiwiZW1haWwiOiJsb2NucDI1QGZwdC5jb20udm4iLCJqdGkiOiIyZWQxNWUzMy1lOWMzLTRjMjctYWI0ZS01ZGEyY2FiZTgwOTAifSwiZXhwIjoxNzI2MDM3Mzc5LjUxNTA4MjF9.qmUAwYoV8FuP6BqOcyN0RX9mnrM3c2nnxy_v19OrXl8"
	serverType  = "646204d4-1995-484a-b469-277619022d22"
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
	reqCreate := requestBodyCreate{
		RegionId:   "hanoi-vn",
		Names:      []string{"pass-check-api"},
		RaidType:   "",
		ServerType: serverType,
		OS:         "ubuntu",
		SshKey:     string(jsonDataKey),
		Distro:     string(jsonDataDistro),
		ClusterId:  2,
		UserData: `
#!/bin/bash

# New hostname
NEW_HOSTNAME="sondx12"

# Change the hostname in /etc/hostname
echo "$NEW_HOSTNAME" > /etc/hostname

# Update /etc/hosts to reflect the new hostname
sed -i "s/127.0.1.1 .*/127.0.1.1 $NEW_HOSTNAME/g" /etc/hosts

# Set the new hostname for the current session
hostnamectl set-hostname "$NEW_HOSTNAME"

# Print a message
echo "Hostname successfully changed to $NEW_HOSTNAME. A reboot may be required for full effect."
reboot`,
	}

	err = CreateMachine(server, vpcId, bearerToken, reqCreate)
	if err != nil {
		fmt.Printf("error when get detail server in vpc: [%v]", err)
	}

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
