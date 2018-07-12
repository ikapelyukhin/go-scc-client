package scc_client

import (
  "fmt"
  "errors"
  "gopkg.in/resty.v1"
)

type Credentials struct {
  Id int `json:"id"`
  Login string `json:"login"`
  Password string `json:"password"`
}

type Product struct {
  Identifier          string        `json:"identifier"`
  Version             string        `json:"version"`
  Arch                string        `json:"arch"`
}

type Service struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Product Product `json:"product"`
}

func AnnounceSystem(regcode string) (*Credentials, error) {
  var credentials Credentials

  resp, err := resty.R().SetHeader("Accept", "application/json").
    SetAuthToken(regcode).
    SetBody(`{}`).
    SetResult(&credentials).
    Post("https://scc.suse.com/connect/subscriptions/systems")

  if (err != nil) {
    return nil, err
  }
  
  if !(resp.StatusCode() >= 200 && resp.StatusCode() < 300) {
    return nil, errors.New(fmt.Sprintf("Request failed with error code %v", resp.StatusCode()))
  }
  
  return &credentials, err
}

func DeregisterSystem(login string, password string) error {
  resp, err := resty.R().SetHeader("Accept", "application/json").
    SetBasicAuth(login, password).
    Delete("https://scc.suse.com/connect/systems")
    
  if (err != nil) {
    return err
  }
  
  if !(resp.StatusCode() >= 200 && resp.StatusCode() < 300) {
    return errors.New(fmt.Sprintf("Request failed with error code %v", resp.StatusCode()))
  }
  
  return nil
}

func GetServices(login string, password string) error {
  resp, err := resty.R().SetHeader("Accept", "application/json").
    SetBasicAuth(login, password).
    Get("https://scc.suse.com/connect/systems/services")
  
  if (err != nil) {
    return err
  }
  
  if !(resp.StatusCode() >= 200 && resp.StatusCode() < 300) {
    return errors.New(fmt.Sprintf("Request failed with error code %v", resp.StatusCode()))
  }
  
  return nil
}

func RegisterProduct(login string, password string, identifier string, version string, arch string, regcode string) (*Service, error) {
  var service Service
  
  resp, err := resty.R().SetHeader("Accept", "application/json").
    SetBasicAuth(login, password).
    SetBody(map [string] interface {} {"identifier": identifier, "version": version, "arch": arch, "token": regcode }).
    SetResult(&service).
    Post("https://scc.suse.com/connect/systems/products")
  
  if (err != nil) {
    return nil, err
  }
  
  if !(resp.StatusCode() >= 200 && resp.StatusCode() < 300) {
    return nil, errors.New(fmt.Sprintf("Request failed with error code %v", resp.StatusCode()))
  }
  
  return &service, err
}

func GetProduct(login string, password string, identifier string, version string, arch string) (*Product, error) {
  var product Product
  
  resp, err := resty.R().SetHeader("Accept", "application/json").
    SetBasicAuth(login, password).
    SetQueryParams(map [string] string {"identifier": identifier, "version": version, "arch": arch }).
    SetResult(&product).
    Get("https://scc.suse.com/connect/systems/products")

  if (err != nil) {
    return nil, err
  }
  
  if !(resp.StatusCode() >= 200 && resp.StatusCode() < 300) {
    return nil, errors.New(fmt.Sprintf("Request failed with error code %v", resp.StatusCode()))
  }
  
  return &product, err
}

func DeactivateProduct(login string, password string, identifier string, version string, arch string) error {
  resp, err := resty.R().SetHeader("Accept", "application/json").
    SetBasicAuth(login, password).
    SetBody(map [string] interface {} {"identifier": identifier, "version": version, "arch": arch }).
    Delete("https://scc.suse.com/connect/systems/products")

    fmt.Printf("Error: %v\n", err)
    fmt.Printf("Response Status Code: %v\n", resp.StatusCode())
    fmt.Printf("Response Body: %v\n\n", resp)
  if (err != nil) {
    return err
  }
  
  if !(resp.StatusCode() >= 200 && resp.StatusCode() < 300) {
    return errors.New(fmt.Sprintf("Request failed with error code %v", resp.StatusCode()))
  }
  
  return nil
}
