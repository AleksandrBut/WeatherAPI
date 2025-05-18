-- +goose Up
-- +goose StatementBegin
create type frequency_type as enum ('hourly', 'daily');

create table subscription
(
    id        serial primary key,
    cityName  varchar(255)   not null,
    email     varchar(255)   not null,
    frequency frequency_type not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table subscription;
drop type frequency_type;
-- +goose StatementEnd
