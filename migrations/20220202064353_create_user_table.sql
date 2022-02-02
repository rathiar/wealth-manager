-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users (
	id int NULL,
	first_name varchar NOT NULL,
	"password" varchar NOT NULL,
	"role" int NOT NULL,
	last_name varchar NOT NULL,
	email varchar NOT NULL
);
CREATE UNIQUE INDEX users_email_idx ON public.users (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.users;
-- +goose StatementEnd
