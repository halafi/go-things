## Simple gql API with pgx

### Dev

1. `$ air`

### Build

1. `$ make`

### Example queries:

```graphql
query allusers {
  users {
    name
  }
}

query alexes {
  users(name:"alex") {
    name
  }
}
```