package song

type messages struct {
	CreateFailed       string
	FetchFailed        string
	ParseFailed        string
	DeleteFailed       string
	MarkListenedFailed string
}

var Messages = messages{
	CreateFailed:       "song_create_failed",
	FetchFailed:        "song_fetch_failed",
	ParseFailed:        "song_parse_failed",
	DeleteFailed:       "song_delete_failed",
	MarkListenedFailed: "song_mark_listened_failed",
}
