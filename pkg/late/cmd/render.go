package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/blang/late/pkg/late/dm"
	"github.com/blang/late/pkg/late/tmpl"
	"github.com/pkg/errors"

	log "github.com/kubernetes/klog"
	"github.com/spf13/cobra"
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render a template based on given data",
	Long:  `Based on the data given in any supported format, the template is rendered and written to stdout by default.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.V(1).Info("Render called")
		renderOpts, err := getRenderOptions(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not process render options", err)
			return
		}

		// Handle file vs stdin
		dataFile, err := cmd.PersistentFlags().GetString("input")
		if err != nil {
			fmt.Fprintln(os.Stderr, "No datafile parameter")
			return
		}
		var dataReader io.Reader
		if dataFile == "" {
			log.V(1).Info("No datafile parameter, using Stdin")
			dataReader = os.Stdin
		} else {
			reader, err := os.Open(dataFile)
			log.V(1).Info("Using datafile parameter instead of Stdin")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not read datafile %s: %s", dataFile, err)
				return
			}
			dataReader = reader
		}

		dataFormat, err := cmd.PersistentFlags().GetString("data-format")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read dataformat", err)
			return
		}

		var dataParser dm.DataParser
		var data interface{}

		switch dataFormat {
		case "json":
			dataParser = &dm.JsonDataModel{}
			err = dataParser.Parse(dataReader, &data)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not parse data: ", err)
				return
			}

		case "yaml":
			dataParser = &dm.YamlDataModel{}
			err = dataParser.Parse(dataReader, &data)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not parse data: ", err)
				return
			}
		}

		templateFile, err := cmd.PersistentFlags().GetString("template")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not use template parameter")
			return
		}

		err = tmpl.Render(data, templateFile, os.Stdout, renderOpts)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not render template: ", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
	renderCmd.PersistentFlags().StringP("template", "t", "template.tmpl", "Template file")
	renderCmd.PersistentFlags().StringP("input", "i", "", "Data file, otherwise stdout")
	renderCmd.PersistentFlags().StringP("data-format", "f", "json", "Data format, json or yaml")
	renderCmd.PersistentFlags().String("delimiterLeft", "{{", "Template syntax left delimiter")
	renderCmd.PersistentFlags().String("delimiterRight", "}}", "Template syntax left delimiter")
}

func getRenderOptions(cmd *cobra.Command) (*tmpl.RenderOptions, error) {
	opt := tmpl.DefaultRenderOptions()
	delimLeft, err := cmd.PersistentFlags().GetString("delimiterLeft")
	if err != nil {
		return nil, errors.Wrap(err, "Could not use delimiterLeft")
	}
	opt.DelimiterLeft = delimLeft
	delimRight, err := cmd.PersistentFlags().GetString("delimiterRight")
	if err != nil {
		return nil, errors.Wrap(err, "Could not use delimiterRight")
	}
	opt.DelimiterRight = delimRight
	return opt, nil
}
