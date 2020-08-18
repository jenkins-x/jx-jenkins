### Linux

```shell
curl -L https://github.com/jenkins-x/jx-jenkins/releases/download/v{{.Version}}/jx-jenkins-linux-amd64.tar.gz | tar xzv 
sudo mv jx-jenkins /usr/local/bin
```

### macOS

```shell
curl -L  https://github.com/jenkins-x/jx-jenkins/releases/download/v{{.Version}}/jx-jenkins-darwin-amd64.tar.gz | tar xzv
sudo mv jx-jenkins /usr/local/bin
```

