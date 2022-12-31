### make launch.json and samconfig.toml
```sh
$ cp ./vscode/launch.json.sample ./vscode/launch.json
$ cp samconfig.toml.sample samconfig.toml
```

### build
```sh
$ sam build --use-container
```

### deploy
```sh
$ sam deploy
```

first time deploy
```sh
$ sam deploy --guided --capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND
```
