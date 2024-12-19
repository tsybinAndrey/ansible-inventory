package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

func buildXrayConfigFromTemplate() {
	var (
		secretsPath, xrayTemplatePath, xrayConfigPath string
	)

	flag.StringVar(&secretsPath, "secrets", "./secrets.yml", "path to ansible-vault secrets file")
	flag.StringVar(&xrayTemplatePath, "xray-config-template", "./xray/reality_config_simple.template.json", "path to xray config template")
	flag.StringVar(&xrayConfigPath, "xray-config-save", "./xray/reality_config_simple.json", "where you want to save finall config")

	cmd := exec.Command("ansible-vault", "view", secretsPath, "--ask-vault-password")

	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Failed to run ansible-vault: %v", err)
	}

	var vaultConfig map[string]interface{}

	err = yaml.Unmarshal(output, &vaultConfig)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	file, err := os.Open(xrayTemplatePath)

	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}

	defer file.Close()

	var xrayConfig map[string]interface{}

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&xrayConfig); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	// template config
	inbounds := xrayConfig["inbounds"].([]interface{})
	settings := inbounds[0].(map[string]interface{})["settings"].(map[string]interface{})
	streamSettings := inbounds[0].(map[string]interface{})["streamSettings"].(map[string]interface{})
	realitySettings := streamSettings["realitySettings"].(map[string]interface{})
	// vault config
	reality := vaultConfig["xray"].(map[string]interface{})["reality"].(map[string]interface{})
	clients := vaultConfig["clients"].([]interface{})

	realitySettings["privateKey"] = reality["privateKey"].(string)
	realitySettings["shortIds"] = append(realitySettings["shortIds"].([]interface{}), reality["shortId"].(string))
	settings["clients"] = clients

	jsonData, err := json.MarshalIndent(xrayConfig, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	resultFile, err := os.Create(xrayConfigPath)

	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer resultFile.Close()

	_, err = resultFile.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to write JSON to file: %v", err)
	}

	fmt.Printf("Config has been written to %s", xrayConfigPath)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command passed. Expected command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build-config":
		buildXrayConfigFromTemplate()
	case "add-user":
		fmt.Println("Here will be add user command")
	default:
		log.Fatal("No command passed. Expected command")
	}
}
