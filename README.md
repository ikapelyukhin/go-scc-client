```
package main

import (
  "fmt"
  "os"
  "./scc_client"
)

func example() {
  regcode := os.Getenv("SCC_REGCODE")
  
  credentials, err := scc_client.AnnounceSystem(regcode)
  if (err != nil) { panic(err) }
  fmt.Printf("Registered system: %v, %v\n", credentials.Login, credentials.Password)
  
  err = scc_client.GetServices(credentials.Login, credentials.Password)
  if (err != nil) { panic(err) }
  
  service, err := scc_client.RegisterProduct(credentials.Login, credentials.Password, "SLES", "12.3", "x86_64", regcode)
  if (err != nil) { panic(err) }
  fmt.Printf("Registered product: %v, %v\n", service.Name, service.URL)
  
  product, err := scc_client.GetProduct(credentials.Login, credentials.Password, service.Product.Identifier, service.Product.Version, service.Product.Arch)
  if (err != nil) { panic(err) }
  fmt.Printf("Get product: %v %v %v\n", product.Identifier, product.Version, product.Arch)
  
  //err = scc_client.DeactivateProduct(credentials.Login, credentials.Password, product.Identifier, product.Version, product.Arch)
  //if (err != nil) { panic(err) }
  //fmt.Printf("Deactivated product\n")
      
  err = scc_client.DeregisterSystem(credentials.Login, credentials.Password)
  if (err != nil) { panic(err) }
  fmt.Printf("Deregistered system\n")
}
```
