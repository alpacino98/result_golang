package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func makeResultRequests(goRoutine int) {
	if goRoutine == 0 {
		getProjekti()
		getOnas()
		getKariera()
		getBlog()
	} else if goRoutine == 1 {
		go getProjekti()
		getOnas()
		getKariera()
		getBlog()
	} else if goRoutine == 2 {
		go getProjekti()
		go getOnas()
		getKariera()
		getBlog()
	} else if goRoutine == 3 {
		go getProjekti()
		go getOnas()
		go getKariera()
		getBlog()
	} else if goRoutine == 4 {
		go getProjekti()
		go getOnas()
		go getKariera()
		go getBlog()
	}
}

func getProjekti() {
	doc, err := goquery.NewDocument(results.Websites[0].Url)
	if err != nil {
		results.Websites[0].Unsuccess = results.Websites[0].Unsuccess + 1
		results.Total_requests.Unsuccess = results.Total_requests.Unsuccess + 1
		log.Fatalln(err)
		lastResults["projeckti_success"] = false
	}
	results.Websites[0].Success = results.Websites[0].Success + 1
	results.Total_requests.Success = results.Total_requests.Success + 1

	pageTitle := doc.Find("title").Contents().Text()
	results.Websites[0].Title = pageTitle
	lastResults["projeckti_success"] = true
}

func getOnas() {
	doc, err := goquery.NewDocument(results.Websites[1].Url)
	if err != nil {
		results.Websites[1].Unsuccess = results.Websites[1].Unsuccess + 1
		results.Total_requests.Unsuccess = results.Total_requests.Unsuccess + 1
		log.Fatalln(err)
		lastResults["onas_success"] = false
	}
	results.Websites[1].Success = results.Websites[1].Success + 1
	results.Total_requests.Success = results.Total_requests.Success + 1
	pageTitle := doc.Find("title").Contents().Text()
	results.Websites[1].Title = pageTitle
	lastResults["onas_success"] = true
}

func getKariera() {
	doc, err := goquery.NewDocument(results.Websites[2].Url)
	if err != nil {
		results.Websites[2].Unsuccess = results.Websites[2].Unsuccess + 1
		results.Total_requests.Unsuccess = results.Total_requests.Unsuccess + 1
		log.Fatalln(err)
		lastResults["kariera_success"] = false
	}
	results.Websites[2].Success = results.Websites[2].Success + 1
	results.Total_requests.Success = results.Total_requests.Success + 1

	pageTitle := doc.Find("title").Contents().Text()
	results.Websites[2].Title = pageTitle
	lastResults["kariera_success"] = true
}

func getBlog() {
	doc, err := goquery.NewDocument(results.Websites[3].Url)
	if err != nil {
		results.Websites[3].Unsuccess = results.Websites[3].Unsuccess + 1
		results.Total_requests.Unsuccess = results.Total_requests.Unsuccess + 1
		log.Fatalln(err)
		lastResults["blog_success"] = false
	}
	results.Websites[3].Success = results.Websites[3].Success + 1
	results.Total_requests.Success = results.Total_requests.Success + 1

	pageTitle := doc.Find("title").Contents().Text()
	results.Websites[3].Title = pageTitle
	lastResults["blog_success"] = true
}
