package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "How are you?",
			Answer:   "Good!",
		},
		{
			Question: "How old are you?",
			Answer:   "33",
		},
		{
			Question: "What are you doing?",
			Answer:   "Programming",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
