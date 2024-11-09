package main

import (
	"fmt"
	"log"
	"os"

	"github.com/getaseww/boiler-gen/src/helpers"
	"github.com/getaseww/boiler-gen/src/templates"

	"github.com/spf13/cobra"
)

var (
	moduleName   string
	filePath     string
	templatePath string
)

// Generate model file inside the model directory
func generateModel(moduleName string) {
	modelDir := "./internal/database/models"

	err := helpers.CreateDirIfNotExist(modelDir)
	if err != nil {
		log.Fatal(err)
	}

	modelContent := templates.ModelTemplate(moduleName)

	filePath := fmt.Sprintf("%s/%s.go", modelDir, helpers.FormatModuleName(moduleName))
	helpers.CreateFile(filePath, modelContent)
	fmt.Printf("Generated model: %s\n", filePath)
}

// Generate model file inside the model directory
func generateRepository(moduleName string) {
	modelDir := "./internal/repositories"
	err := helpers.CreateDirIfNotExist(modelDir)
	if err != nil {
		log.Fatal(err)
	}

	repositoryContent := templates.RepositoryTemplate(moduleName)

	filePath := fmt.Sprintf("%s/%s.go", modelDir, helpers.FormatModuleName(moduleName))
	helpers.CreateFile(filePath, repositoryContent)
	fmt.Printf("Generated repository: %s\n", filePath)
}

// Generate model file inside the model directory
func generateService(moduleName string) {
	modelDir := "./internal/services"
	err := helpers.CreateDirIfNotExist(modelDir)
	if err != nil {
		log.Fatal(err)
	}

	serviceContent := templates.ServiceTemplate(moduleName)

	filePath := fmt.Sprintf("%s/%s.go", modelDir, helpers.FormatModuleName(moduleName))
	helpers.CreateFile(filePath, serviceContent)
	fmt.Printf("Generated model: %s\n", filePath)
}

// Generate controller file inside the controller directory
func generateHandler(moduleName string) {
	handlerDir := "./api/v1/handlers"
	err := helpers.CreateDirIfNotExist(handlerDir)
	if err != nil {
		log.Fatal(err)
	}

	handlerContent := templates.HandlerTemplate(moduleName)

	filePath := fmt.Sprintf("%s/%s.go", handlerDir, helpers.FormatModuleName(moduleName))
	helpers.CreateFile(filePath, handlerContent)
	fmt.Printf("Generated handler: %s\n", filePath)
}

// Generate route file inside the routes directory
func generateRoute(moduleName string) {
	routeDir := "./api/v1/routes"
	err := helpers.CreateDirIfNotExist(routeDir)
	if err != nil {
		log.Fatal(err)
	}

	routeContent := templates.RouteTemplate(moduleName)

	filePath := fmt.Sprintf("%s/%s.go", routeDir, helpers.FormatModuleName(moduleName))
	helpers.CreateFile(filePath, routeContent)
	fmt.Printf("Generated route: %s\n", filePath)
}

// Command for generating the model file
var modelCmd = &cobra.Command{
	Use:   "make:model",
	Short: "Generate a model file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating model for module: %s\n", moduleName)
		generateModel(moduleName)
	},
}

// Command for generating the repository file
var repositoryCmd = &cobra.Command{
	Use:   "make:repository",
	Short: "Generate a repository file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating repository for module: %s\n", moduleName)
		generateRepository(moduleName)
	},
}

// Command for generating the service file
var serviceCmd = &cobra.Command{
	Use:   "make:service",
	Short: "Generate a service file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating service for module: %s\n", moduleName)
		generateService(moduleName)
	},
}

// Command for generating the handler file
var handlerCmd = &cobra.Command{
	Use:   "make:handler",
	Short: "Generate a handler file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating handler for module: %s\n", moduleName)
		generateHandler(moduleName)
	},
}

// Command for generating the route file
var routeCmd = &cobra.Command{
	Use:   "make:route",
	Short: "Generate a route file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating route for module: %s\n", moduleName)
		generateRoute(moduleName)
	},
}

// Command for generating the full feature (model, controller, and route)
var featureCmd = &cobra.Command{
	Use:   "make:feature",
	Short: "Generate model, repository, service, handler, and route files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating all files for module: %s\n", moduleName)
		generateModel(moduleName)
		generateRepository(moduleName)
		generateService(moduleName)
		generateHandler(moduleName)
		generateRoute(moduleName)
	},
}

func main() {
	// Root command for the CLI
	var rootCmd = &cobra.Command{
		Use:   "boiler-gen",
		Short: "A CLI tool to generate boilerplate",
	}

	// Setting up the command flags and parameters
	modelCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Name of the module (required)")
	modelCmd.MarkFlagRequired("module")

	repositoryCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Name of the module (required)")
	repositoryCmd.MarkFlagRequired("module")

	serviceCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Name of the module (required)")
	serviceCmd.MarkFlagRequired("module")

	handlerCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Name of the module (required)")
	handlerCmd.MarkFlagRequired("module")

	featureCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Name of the module (required)")
	featureCmd.MarkFlagRequired("module")

	routeCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Name of the module (required)")
	routeCmd.MarkFlagRequired("module")

	// Add commands to the root
	rootCmd.AddCommand(modelCmd)
	rootCmd.AddCommand(handlerCmd)
	rootCmd.AddCommand(featureCmd)
	rootCmd.AddCommand(routeCmd)
	rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(repositoryCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
