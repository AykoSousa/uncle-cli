package cmd

import (
	"fmt"
	"os"

	"github.com/AykoSousa/uncle-cli/cmd/core/gemini"
	"github.com/AykoSousa/uncle-cli/cmd/core/groq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	llm string
	API_KEY string
	model string
	userPrompt string
)

func init() {
	// Initialize config
	cobra.OnInitialize(initConfig)

	// Flags
	rootCmd.Flags().StringVarP(&API_KEY, "api-key", "k", "", "YOUR_API_KEY")
	rootCmd.Flags().StringVarP(&model, "model", "m", "", "AI model, like: gemini-2.0-flash")
	rootCmd.Flags().StringVarP(&llm, "llm", "l", "", "LLM name, ex: gemini, chatgpt, ...")
	rootCmd.Flags().StringVarP(&userPrompt, "prompt", "p", "", "What is DevOps culture?")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "./uncle-config.json")

	// Required Flags
	rootCmd.MarkFlagRequired("prompt")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigFile(".\\uncle-config.json")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

var rootCmd = &cobra.Command{
	Use:   "uncle",
	Short: "A smart cli to help you all day",
	Long: `A smart CLI to help you get quick answers via the terminal, without needing to access another page to do it.
		Documentation is available at https://github.com/AykoSousa/uncle-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get defaults configs
		if llm == "" {
			llm = viper.GetString("llm")
		}
		if model == "" {
			model = viper.GetString("model")
		}
		if API_KEY == "" {
			API_KEY = viper.GetString("apiKey")
		}

		// Call LLM
		if llm == "gemini" {
			gemini.Gemini(API_KEY, model, userPrompt)
		}

		if llm == "groq" {
			fmt.Println("You can check availables models of groq here: https://console.groq.com/docs/models")
			groq.Groq(API_KEY, model, userPrompt)
		}
	},
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Fprintln(os.Stderr, err)
	  os.Exit(1)
	}
}