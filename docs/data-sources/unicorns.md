---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "crud_unicorns Data Source - terraform-provider-crud"
subcategory: ""
description: |-
  
---

# crud_unicorns (Data Source)



## Example Usage

```terraform
data "crud_unicorns" "all" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `unicorns` (Attributes List) (see [below for nested schema](#nestedatt--unicorns))

<a id="nestedatt--unicorns"></a>
### Nested Schema for `unicorns`

Read-Only:

- `age` (Number)
- `colour` (String)
- `id` (String)
- `name` (String)


