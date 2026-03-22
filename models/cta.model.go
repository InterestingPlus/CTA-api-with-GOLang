package models

type CtaType string

const (
	BookDemo    CtaType = "book_demo"
	TalkToSales CtaType = "talk_to_sales"
	Career      CtaType = "career"
	Contact     CtaType = "contact"
)

type Cta struct {
	Type    []CtaType `json:"type"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Message string    `json:"message"`
}
