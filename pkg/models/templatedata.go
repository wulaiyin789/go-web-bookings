package models

// https://www.udemy.com/course/building-modern-web-applications-with-go/learn/lecture/22888920
// "import cycle not allowed"
// I have packages that are importing each other and you're not allowed to do
// 						project/controllers/account
// 							^                    \
// 						 /                      \
// 					  /                        \
// 				   /                         \/
// 	project/components/mux <--- project/controllers/base

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string // message
	Warning   string // message
	Error     string // message
}
