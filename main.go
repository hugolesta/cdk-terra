package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-docker-go/docker"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here

	docker.NewDockerProvider(stack, jsii.String("docker"), &docker.DockerProviderConfig{})

    dockerImage := docker.NewImage(stack, jsii.String("nginxImage"), &docker.ImageConfig{
        Name:        jsii.String("nginx:latest"),
        KeepLocally: jsii.Bool(false),
    })

    docker.NewContainer(stack, jsii.String("nginxContainer"), &docker.ContainerConfig{
        Image: dockerImage.Latest(),
        Name:  jsii.String("tutorial"),
        Ports: &[]*docker.ContainerPorts{{
            Internal: jsii.Number(80), External: jsii.Number(8000),
        }},
    })

    return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdk-terra")

	app.Synth()
}
