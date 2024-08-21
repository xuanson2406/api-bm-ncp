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
	server      = ""
	vpcId       = ""
	bearerToken = ""
)

type requestBodyCreate struct {
	RegionId   string `json:"regionId,omitempty"`  // Region id (hanoi-vn)
	ClusterId  int32  `json:"clusterId,omitempty"` // Cluster id
	ServerType string `json:"serverType"`          // Server type ID as uuid
	Names      string `json:"names"`               // Define name of machines, it will be understanded as number of machines
	OS         string `json:"os"`                  // OS type that machines will be installed (ubuntu/centos..)
	RaidType   string `json:"raidType,omitempty"`  // Raid type that machine will be configured as a storage type (level-1) (Not using right now)
	SshKey     string `json:"sshKey,omitempty"`    // sshKey that user already imported on portal, it must be an identify sshkey on portal
	UserData   string `json:"userData,omitempty"`  // post user defines script after os installed
	Distro     string `json:"distro,omitempty"`    // Json dump with keys name and hwe_kernel. Name as distro name, hwe_kernel as architecture
}

type requestBodyPut struct { // Update a machine (Turn ON/OFF machine)
	Name  string `json:"name,omitempty"`  // Name of machine that will be update (Null means no modification)
	Power string `json:"power,omitempty"` // Power type of mechine that will be applied (Null means no modification) [ on, off, null ]
}
type responseBodyCreateAccepted struct { // Created accepted (202)
	Data      Msg    `json:"data"`
	IsSuccess bool   `json:"isSuccess"`
	Error     string `json:"error,omitempty"`
}

type Msg struct {
	Message string `json:"message"`
}

type responseBodyCreateError struct { // BadRequest (400) Unauthorized (401) InternalServerError (500) NotFound (404)
	IsSuccess bool   `json:"isSuccess"`
	Error     string `json:"error,omitempty"` // (optional) Reasons
}

type responseBody struct { // Update/Get/Delete a machine successfully  (200)  BadRequest (400) Unauthorized (401) InternalServerError (500) NotFound (404)
	IsSuccess bool   `json:"isSuccess"`
	Error     string `json:"error,omitempty"` // (optional) Reasons
}

func main() {

}

// Create a machine with user's predefination
func CreateMachine(server, vpcId, bearerToken string, reqCreate requestBodyCreate) error {
	portalServer := "https://" + server + "/vmware/vpc/" + vpcId + "/hpc/server/create"
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
		fmt.Printf("Data: %s\n", result.Data.Message)
		fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
		fmt.Printf("Error: %s\n", result.Error)
		return nil
	}
	var result responseBodyCreateError
	json.Unmarshal(body, &result)
	fmt.Printf("Error: %s\n", result.Error)
	fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
	return fmt.Errorf("error to create machine: [%v]", result.Error)
}

// Update a machine state
func UpdateMachine(server, vpcId, bearerToken, systemId string, reqPut requestBodyPut) error {
	portalServer := "https://" + server + "/vpc/" + vpcId + "/hpc/machines/" + systemId
	bearer := "Bearer " + bearerToken
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	putBody, _ := json.Marshal(reqPut)
	reqBody := bytes.NewBuffer(putBody)
	clienthttp := &http.Client{Transport: tr}
	req, _ := http.NewRequest("PUT", portalServer, reqBody)
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := clienthttp.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform PUT request : [%v]", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body of PUT to Update a machine state: [%v]", err)
	}
	var result responseBody
	json.Unmarshal(body, &result)
	if resp.StatusCode == 200 {
		fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
		fmt.Printf("Error: %s\n", result.Error)
		return nil
	}
	fmt.Printf("Error: %s\n", result.Error)
	fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
	return fmt.Errorf("error to update machine: [%v]", result.Error)
}

// Get machine detail with machine ID
func GetMachine(server, vpcId, bearerToken, systemId string) error {
	portalServer := "https://" + server + "/vpc/" + vpcId + "/hpc/machines/" + systemId
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
	var result responseBody
	json.Unmarshal(body, &result)
	if resp.StatusCode == 200 {
		fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
		fmt.Printf("Error: %s\n", result.Error)
		return nil
	}
	fmt.Printf("Error: %s\n", result.Error)
	fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
	return fmt.Errorf("error to GET machine: [%v]", result.Error)
}

// Release a machine with machine ID
func DeleteMachine(server, vpcId, bearerToken, systemId string) error {
	portalServer := "https://" + server + "/vpc/" + vpcId + "/hpc/machines/" + systemId
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
	var result responseBody
	json.Unmarshal(body, &result)
	if resp.StatusCode == 200 {
		fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
		fmt.Printf("Error: %s\n", result.Error)
		return nil
	}
	fmt.Printf("Error: %s\n", result.Error)
	fmt.Printf("IsSuccess: %t\n", result.IsSuccess)
	return fmt.Errorf("error to DELETE machine: [%v]", result.Error)
}
