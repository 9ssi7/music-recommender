package user

type messages struct {
	EmailAlreadyExists string
	CreateFailed       string
	FetchFailed        string
	ParseFailed        string
}

var Messages = messages{
	EmailAlreadyExists: "user_email_already_exists",
	CreateFailed:       "user_create_failed",
	FetchFailed:        "user_fetch_failed",
	ParseFailed:        "user_parse_failed",
}
