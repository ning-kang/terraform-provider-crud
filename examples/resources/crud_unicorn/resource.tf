terraform {
  required_providers {
    crud = {
      source = "crud.com/crud/crud"
    }
  }
}

provider "crud" {
  endpoint = "https://crudcrud.com/api/d91f56d5de624701a2c1e3ecbc8cc29a"
}

resource "crud_unicorn" "second" {
  name   = "second"
  age    = 5
  colour = "green"
}

output "crud_unicorns" {
  value = crud_unicorn.second.id
}