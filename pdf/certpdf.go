package pdf

import (
	"fmt"
	"generate-certificat/cert"
	"github.com/jung-kurt/gofpdf"
	"os"
	"path"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutputDir: outputdir,
	}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, true)
	pdf.AddPage()

	p.background(pdf)
	p.header(cert, pdf)

	pdf.Ln(30)

	p.body(cert, pdf)
	p.footer(pdf)

	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	pathJoined := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(pathJoined)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate %v", pathJoined)
	return nil
}

func (p *PdfSaver) background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions(
		"img/background.png",
		0,
		0,
		pageWidth,
		pageHeight,
		false,
		opts,
		0,
		"",
	)
}

func (p *PdfSaver) header(cert cert.Cert, pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/csops.png"
	pdf.ImageOptions(
		filename,
		x+margin,
		20,
		imageWidth,
		0,
		false,
		opts,
		0,
		"",
	)

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(
		filename,
		x-margin,
		20,
		imageWidth,
		0,
		false,
		opts,
		0,
		"",
	)

	pdf.SetFont("Helvetica", "", 35)
	pdf.WriteAligned(0, 50, cert.LabelCompletion, "C")
}

func (p *PdfSaver) body(cert cert.Cert, pdf *gofpdf.Fpdf) {
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	// Student name
	pdf.SetFont("Times", "B", 35)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	// Body Participation
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	// Date
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")
}


func (p *PdfSaver) footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	imageWidth := 50.0
	filename := "img/stamp.png"
	pageWidth, pageHeight := pdf.GetPageSize()
	x := pageWidth - imageWidth - 20.0
	y := pageHeight - imageWidth - 10.0

	pdf.ImageOptions(
		filename,
		x,
		y,
		imageWidth,
		0,
		false,
		opts,
		0,
		"",
	)
}
