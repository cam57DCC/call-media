create table response_headers
(
    id serial not null,
    request_id bigint(20) unsigned not null,
    header varchar(512) not null,
    header_value text not null,
    constraint response_headers_pk
        primary key (id),
    constraint response_headers_requests_id_fk
        foreign key (request_id) references requests (id)
            on delete cascade
);