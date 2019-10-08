package domain

type Zone struct {
	ZoneId       int    `db:"id"`
	DeviceId     int    `db:"deviceId"`
	PageId       int    `db:"pageId"`
	LanguageCode string `db:"languageCode"`
	Width        int
	Height       int
	Banners      []Banner
}

func NewZone(zoneId, deviceId, pageId int, languageCode string, width, height int, banners []Banner) Zone {
	z := Zone{zoneId, deviceId, pageId, languageCode, width, height, banners}
	return z
}
