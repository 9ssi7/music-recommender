package genre

type messages struct {
	CreateFailed string
	FetchFailed  string
	ParseFailed  string
	DeleteFailed string
}

var Messages = messages{
	CreateFailed: "genre_create_failed",
	FetchFailed:  "genre_fetch_failed",
	ParseFailed:  "genre_parse_failed",
	DeleteFailed: "genre_delete_failed",
}
