/*
Copyright © 2024 blacktop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	text    bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "what-dis <IMAGE>",
	Short: "Describe an image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(filepath.Clean(args[0]))
		if err != nil {
			log.Fatal(fmt.Errorf("could not read image: %w", err))
		}

		cli, err := api.ClientFromEnvironment()
		if err != nil {
			log.Fatal(fmt.Errorf("could not create ollama client: %w", err))
		}

		systemMsg := "Provide very brief, concise responses"
		if verbose {
			systemMsg = "Provide very clear and detailed responses"
		}

		messages := []api.Message{
			{
				Role:    "system",
				Content: systemMsg,
			},
			{
				Role:    "user",
				Content: "What is in this image?",
			},
			{
				Images: []api.ImageData{data},
			},
		}

		var respFunc api.ChatResponseFunc

		if text {
			respFunc = func(cr api.ChatResponse) error {
				fmt.Println(cr.Message.Content)
				return nil
			}
		} else {
			respFunc = func(cr api.ChatResponse) error {
				cmd := exec.Command("/usr/bin/say", "--rate=200")
				cmd.Stdin = strings.NewReader(cr.Message.Content)
				if err := cmd.Run(); err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err)
				}
				return nil
			}
		}

		if err := cli.Chat(context.Background(), &api.ChatRequest{
			Model:    "llama3.2-vision",
			Messages: messages,
			Stream:   new(bool),
		}, respFunc); err != nil {
			log.Fatal(fmt.Errorf("could not chat with ollama: %w", err))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "V", false, "More detailed output")
	rootCmd.Flags().BoolVarP(&text, "text", "t", false, "Output as text")
}
