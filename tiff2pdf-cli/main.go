package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/companieshouse/go-tiff2pdf/tiff2pdf"
	"github.com/spf13/cobra"
)

func cmd() *cobra.Command {

	output := "output.pdf"
	config := tiff2pdf.DefaultConfig()

	c := cobra.Command{
		Use:   "t2p <input-tiff-file>",
		Short: "A tool for covert tiff to pdf",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("invalid arguments, require input tiff file")
			}
			input, err := filepath.Abs(args[0])
			if err != nil {
				return err
			}

			b, err := ioutil.ReadFile(input)
			if err != nil {
				return err
			}

			o, err := tiff2pdf.ConvertTiffToPDF(b, config, filepath.Base(input), filepath.Base(output))
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(output, o.PDF, 0664); err != nil {
				return err
			}

			return nil
		},
	}

	c.Flags().StringVarP(&output, "output", "o", output, "output file name")
	c.Flags().StringVar(&config.PageSize, "page-size", config.PageSize, "sets the PDF page size, e.g. legal, letter or A4")
	c.Flags().BoolVar(&config.FullPage, "full-page", config.FullPage, "makes the tiff image fill the PDF page")
	c.Flags().StringVar(&config.Subject, "subject", config.Subject, "the document description")
	c.Flags().StringVar(&config.Author, "author", config.Author, "the document name")
	c.Flags().StringVar(&config.Creator, "creator", config.Creator, "the image software used to create the document")
	c.Flags().StringVar(&config.Title, "title", config.Title, "the document title")

	//c.SetUsageTemplate(`Use "` + filepath.Base(os.Args[0]) + ` --help" for more information.`)

	return &c
}

func main() {
	if err := cmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
