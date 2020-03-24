# TODO list

## debug

`JSON_INPUT` arg doesn't work

```
PUT /cloud/project/3554d1c5d638452f94854da78af239c2/instance/8e552d6a-5d51-4680-9b86-3e7801ae8c36 '{"instanceName" : "pipo"}'
```

but it works

```
echo '{"instanceName" : "pipo"}' | ./ovh-cli --debug PUT /cloud/project/3554d1c5d638452f94854da78af239c2/instance/8e552d6a-5d51-4680-9b86-3e7801ae8c36
```

- add a test, release v0.2

## move deploy.sh to its own project

## automate all build to its own public cloud VM
