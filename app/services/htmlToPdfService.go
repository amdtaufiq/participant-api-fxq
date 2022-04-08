package services

import (
	"bytes"
	"html/template"
	"participant-api/app/entities"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type IHtmlToPDFService interface {
	GenerateNameTag(data entities.Participant) ([]byte, error)
	GenerateCertificate(data entities.Participant) ([]byte, error)
}

type htmlToPDFService struct{}

// const path = "./wkhtmltopdf"
const path = "/app/bin/wkhtmltopdf"

func HtmlToPDFService() *htmlToPDFService {
	return &htmlToPDFService{}
}

func (p htmlToPDFService) GenerateNameTag(participant entities.Participant) ([]byte, error) {
	var templ *template.Template
	var err error
	currentTime := time.Now()

	data := content{
		ID:           currentTime.Format("20060102150405"),
		FullName:     participant.FullName,
		BusinessName: participant.BusinessName,
		Date:         currentTime.Format("02 Jan, 2006"),
	}
	// use Go's default HTML template generation tools to generate your HTML
	if templ, err = template.ParseFiles("templates/name-tag.html"); err != nil {
		return nil, err
	}
	// apply the parsed HTML template data and keep the result in a Buffer
	var body bytes.Buffer
	if err = templ.Execute(&body, data); err != nil {
		return nil, err
	}

	//set path
	wkhtmltopdf.SetPath(path)

	// initalize a wkhtmltopdf generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// read the HTML page as a PDF page
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	// enable this if the HTML file contains local references such as images, CSS, etc.
	page.EnableLocalFileAccess.Set(true)

	// add the page to your generator
	pdfg.AddPage(page)

	// manipulate page attributes as needed
	pdfg.MarginLeft.Set(1)
	pdfg.MarginRight.Set(2)
	pdfg.MarginTop.Set(2)
	pdfg.Dpi.Set(350)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA7)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	// magic
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}

func (p htmlToPDFService) GenerateCertificate(participant entities.Participant) ([]byte, error) {
	var templ *template.Template
	var err error
	currentTime := time.Now()

	data := content{
		ID:           currentTime.Format("20060102150405"),
		FullName:     participant.FullName,
		BusinessName: participant.BusinessName,
		Date:         currentTime.Format("02 Jan, 2006"),
	}
	// use Go's default HTML template generation tools to generate your HTML
	if templ, err = template.ParseFiles("templates/certificate.html"); err != nil {
		return nil, err
	}
	// apply the parsed HTML template data and keep the result in a Buffer
	var body bytes.Buffer
	if err = templ.Execute(&body, data); err != nil {
		return nil, err
	}

	//set path
	wkhtmltopdf.SetPath(path)

	// initalize a wkhtmltopdf generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// read the HTML page as a PDF page
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	// enable this if the HTML file contains local references such as images, CSS, etc.
	page.EnableLocalFileAccess.Set(true)

	// add the page to your generator
	pdfg.AddPage(page)

	// manipulate page attributes as needed
	pdfg.MarginLeft.Set(5)
	pdfg.MarginRight.Set(5)
	pdfg.MarginTop.Set(5)
	pdfg.MarginBottom.Set(5)
	pdfg.Dpi.Set(350)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	// magic
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}

type content struct {
	ID           string
	FullName     string
	BusinessName string
	Date         string
}
