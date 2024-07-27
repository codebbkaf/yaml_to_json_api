convert yaml to json
```
curl --location 'http://localhost:8080/yaml-to-json' \
--header 'Content-Type: application/json' \
--data '{
  "yaml": "version: '\''3.1'\''\nservices:\n  mongo:\n    image: mongo:latest\n    container_name: mongodb\n    ports:\n      - \"27017:27017\"\n    environment:\n      MONGO_INITDB_ROOT_USERNAME: root\n      MONGO_INITDB_ROOT_PASSWORD: example\n    volumes:\n      - mongo_data:/data/db\n\nvolumes:\n  mongo_data:\n    driver: local\n"
}'

```

convert properties to json
```
curl --location 'http://localhost:8080/yaml-to-json' \
--header 'Content-Type: application/json' \
--data '{"properties":"main.config.enable=true\nmain.config.maxSize=100\nmain.config.minSize=10\ndatabase.type=mysql\ndatabase.host=localhost\ndatabase.port=3306\ndatabase.user=root\ndatabase.password=rootPassword\nfeatures.login.enable=true\nfeatures.search.enable=false\nfeatures.dashboard.enable=true\napi.timeout=30\napi.retryTimes=5\nclient.defaultSetting.language=en\nclient.defaultSetting.region=US\nlogging.level=DEBUG\nlogging.file.path=/var/log/app.log\nlogging.file.rotation.size=10MB\nlogging.console.enable=true\nserver.port=8080\nserver.mode=production\nsecurity.ssl.enabled=true\nsecurity.encryption.keys=ABCDEFGHIJ"}'
```
