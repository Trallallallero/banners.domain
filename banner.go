package domain

type Banner struct {
	ZoneId            int `db:"zoneId"`
	Id                int
	Date              string
	Background_color  string
	Background_image  string
	Button_text       string
	Button_color      string
	Button_background string
	Title             string
	Description       string
	Text_align        string
	Link              string
	Link_isExternal   string `db:"link_isExternal"`
}

func NewBanner(zoneId, id int, date, background_color, background_image, button_text, button_color, button_background, title,
	description, text_align, link, link_isExternal string) Banner {
	banner := Banner{zoneId, id, date, background_color, background_image, button_text, button_color, button_background,
		title, description, text_align, link, link_isExternal}
	return banner
}
