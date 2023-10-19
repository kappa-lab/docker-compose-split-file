# docker-compose-split-file
同一containerに複数docker-composeファイルで同居させる

.envをdocker-compose.ymlと同階層に置くことでcontainerNameを指定し、同一containerに配置できる
```
COMPOSE_PROJECT_NAME=docker-compose-split-file
```
