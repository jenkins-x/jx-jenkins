# jx-jenkins

[![Documentation](https://godoc.org/github.com/jenkins-x/jx-jenkins?status.svg)](https://pkg.go.dev/mod/github.com/jenkins-x/jx-jenkins)
[![Go Report Card](https://goreportcard.com/badge/github.com/jenkins-x/jx-jenkins)](https://goreportcard.com/report/github.com/jenkins-x/jx-jenkins)
[![Releases](https://img.shields.io/github/release-pre/jenkins-x-labs/trigger-pipeline.svg)](https://github.com/jenkins-x/jx-jenkins/releases)
[![LICENSE](https://img.shields.io/github/license/jenkins-x-labs/trigger-pipeline.svg)](https://github.com/jenkins-x/jx-jenkins/blob/master/LICENSE)
[![Slack Status](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://slack.k8s.io/)

`jx-jenkins` is a small command line tool for working with [Jenkins](https://jenkins.io/) servers on kubernetes.

It helps bridge the gap between [Jenkins](https://jenkins.io/) and [Jenkins X](https://jenkins-x.io/) around developer UX and ChatOps.


## Using Trigger Pipeline CLI

Download the binary and put it on your `$PATH` or use the trigger-pipeline container image in a pipeline step in Jenkins or Jenkins X / Tekton.

To trigger a pipeline in the default Jenkins server:

``` 
jx jenkins trigger
```

If you have multiple Jenkins server custom resources in the current namespace then specify the Jenkins custom resource name on the command line:

``` 
jx jenkins trigger  --jenkins myJenkinsServer
```

If you are not sure what Jenkins names are available in the current namespace run:

``` 
kubectl get jenkins
```

Which will list all of the available Jenkins custom resource names.

If you make a mistake and use a name which does not exist the `tp` executable will give you a meaninful error message togehter with listing all the available Jenkins custom resource names.


### Using Environment Variables

To make it easier to configure inside pipelines you can specify the Jenkins instance name via the `$TRIGGER_JENKINS_SERVER` environment variable

```   
export TRIGGER_JENKINS_SERVER="someJenkinsCrdName"
jx jenkins trigger
```
 
For more information type: 

``` 
jx jenkins trigger --help
```       

## Adding Jenkins Servers

`trigger-pipeline` can automatically discover Jenkins servers created via the [Jenkins Operator](https://jenkinsci.github.io/kubernetes-operator/).

In addition you can register any Jenkins servers you wish to the Jenkins Server Registry via the `jx jenkins add` command.

To add a new Jenkins server with a guided wizard:

```
jx jenkins add 
```

If you already know the name, URL, username and API Token then you can use:

```
jx jenkins add 
```

### Removing Jenkins Servers

You can remove a Jenkins server via:

``` 
jx jenkins remove
```

Note that this only removes it from the registry; it doesn't affect the actual Jenkins Server.

## Listing the available Jenkins Servers

To list the servers you can use try:

``` 
jx jenkins list
```

## How it works

To maintain a registry of Jenkins Servers `trigger-pipeline` uses a Kubernetes `Secret` for each Jenkins Server with details of the URL, username and API Token 

## Known issues

If you see this error when trying to trigger a pipeline:

``` 
403 No valid crumb was included in the request
```

Then until we figure out a better workaround you need to go into `Manage Jenkins` -> `Configure Global Security` then make sure you uncheck `Prevent Cross Site Request Forgery exploits` 