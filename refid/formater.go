package refid

type RefIdFormatter struct {
	ID      string `json:"id"`
	Flgused int    `json:"flgused"`
}

func FormatRefId(refid RefId) RefIdFormatter {
	formatter := RefIdFormatter{
		ID:      refid.ID,
		Flgused: refid.Flgused,
	}

	return formatter
}
