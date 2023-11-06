package pdfutils

import (
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

// Creates a new title paragraph with predefined styles
func CreateTitle(c *creator.Creator, text string, font *model.PdfFont) (*creator.Paragraph, error) {
	title := c.NewParagraph(text)
	title.SetFont(font)
	title.SetFontSize(18)
	title.SetMargins(0, 0, 0, 30)
	title.SetTextAlignment(creator.TextAlignmentCenter)
	return title, nil
}

// Creates a new subtitle paragraph with predefined styles
func CreateSubtitle(c *creator.Creator, text string, font *model.PdfFont) (*creator.Paragraph, error) {
	subtitle := c.NewParagraph(text)
	subtitle.SetFont(font)
	subtitle.SetFontSize(12)
	subtitle.SetMargins(20, 0, 10, 0)
	return subtitle, nil
}

// CreateTable creates a new table with predefined styles and populates it with the given data.
func CreateTable(c *creator.Creator, headers []string, data [][]string, fontRegular *model.PdfFont) (*creator.Table, error) {
	table := c.NewTable(len(headers)) // The number of columns is based on the number of headers
	table.SetMargins(0, 0, 30, 10)

	// Set the headers for the table.
	for _, headerText := range headers {
		paragraph := c.NewParagraph(headerText)
		paragraph.SetFont(fontRegular)
		paragraph.SetFontSize(9)
		paragraph.SetMargins(2, 2, 2, 2)
		cell := table.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		err := cell.SetContent(paragraph)
		if err != nil {
			return nil, err
		}
	}

	// Populate the table rows with data.
	for _, row := range data {
		for _, colText := range row {
			paragraph := c.NewParagraph(colText)
			paragraph.SetFont(fontRegular)
			paragraph.SetFontSize(9)
			paragraph.SetMargins(2, 2, 2, 2)
			cell := table.NewCell()
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
			cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
			cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)
			err := cell.SetContent(paragraph)
			if err != nil {
				return nil, err
			}
		}
	}

	return table, nil
}
