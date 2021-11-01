package main

import (
	"fmt"
	argocd "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

type Application struct {
}

// Render return deployment struct
func (a *Application) Render() *argocd.Application {
	return nil
}

func main() {
	fmt.Println("hello argocd!")
}
