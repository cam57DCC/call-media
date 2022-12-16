create table requests
(
    id serial not null,
    url varchar(512) not null,
    response_code int null,
    response text
);

create unique index requests_id_uindex
	on requests (id);