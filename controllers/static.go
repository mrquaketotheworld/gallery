package controllers

import (
	"github.com/mrquaketotheworld/gallery/views"
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
		tpl.Execute(w, questions)
	}
}
