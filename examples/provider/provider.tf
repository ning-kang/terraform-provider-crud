terraform {
  required_providers {
    crud = {
      source = "crud.com/crud/crud"
    }
  }
}

provider "crud" {
  endpoint = "https://crudcrud.com/api/de7cbd95b4c9466cbbbf73a407f5afc3"
}