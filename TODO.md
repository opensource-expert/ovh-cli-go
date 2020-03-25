# TODO list

## GET convert JSON input to QueryString ?

```
  myovh_cli GET /domain/zone/$domain/record \
    "{ \"subDomain\" : \"$subdomain\", \"fieldType\" : \"$fieldType\" }"
```

```
  local url="/domain/zone/$domain/record?subDomain=$subdomain&fieldType=$fieldType"
  ovh-cli GET "$url" | jq -r '.[0]'
```


## move deploy.sh to its own project

## automate all build to its own public cloud VM
