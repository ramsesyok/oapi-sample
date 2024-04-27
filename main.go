/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "landmarks/cmd"

//go:generate oapi-codegen --config=oapi-codegen-echo.yaml ./api/openapi/landmarks.yaml
//go:generate oapi-codegen --config=oapi-codegen-models.yaml ./api/openapi/landmarks.yaml

func main() {
	cmd.Execute()
}
