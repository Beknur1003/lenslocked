package controllers

import (
	"html/template"
	"lens/views"
	"net/http"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free trial for 30 days on any paid plans.",
		},
		{
			Question: "Can I cancel my subscription at any time?",
			Answer:   "Yes! Send us an email and we'll process your request no questions asked.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We offer support 24/7. Please email us if you have any questions!",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
