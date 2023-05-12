terraform {
  required_providers {
    crud = {
      source = "crud.com/crud/unicorns"
    }
  }
}

provider "crud" {
   endpoint = "https://crudcrud.com/api/b93947cfec7840c9aba7f57e2bae87e8/unicorns"
}
