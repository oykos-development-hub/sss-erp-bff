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
		cell.SetBackgroundColor(creator.ColorRGBFromHex("#F3FFF8"))
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

// CreateVerticalHeaderTableWithMap creates a new table with the first column as headers and second as data from a map.
func CreateVerticalHeaderTable(c *creator.Creator, dataMap map[string]string, fontRegular *model.PdfFont) (*creator.Table, error) {
	table := c.NewTable(2) // One for the header, one for the data.
	table.SetMargins(0, 0, 30, 10)

	// Loop through the map to set the headers and data.
	for header, value := range dataMap {
		// Header
		headerParagraph := c.NewParagraph(header)
		headerParagraph.SetFont(fontRegular)
		headerParagraph.SetFontSize(9)
		headerParagraph.SetMargins(2, 2, 2, 2)
		headerCell := table.NewCell()
		headerCell.SetBackgroundColor(creator.ColorRGBFromHex("#F3FFF8"))
		headerCell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		err := headerCell.SetContent(headerParagraph)
		if err != nil {
			return nil, err
		}

		// Data
		dataParagraph := c.NewParagraph(value)
		dataParagraph.SetFont(fontRegular)
		dataParagraph.SetFontSize(9)
		dataParagraph.SetMargins(2, 2, 2, 2)
		dataCell := table.NewCell()
		dataCell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		dataCell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
		dataCell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)
		err = dataCell.SetContent(dataParagraph)
		if err != nil {
			return nil, err
		}
	}

	return table, nil
}
