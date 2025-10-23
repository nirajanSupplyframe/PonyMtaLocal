
CREATE table IF NOT EXISTS domain(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT unique not null
);

create table IF NOT EXISTS mta_status_text (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    short text unique not null,
    support text unique not null,
    customer text unique not null default ''
);


CREATE TABLE  IF NOT EXISTS application (
    id integer primary key autoincrement,
    apikey text unique not null,
    name text not null,
    domain_id int not null,
    foreign key (domain_id) references domain(id)

);

create table IF NOT EXISTS email_status (
    id text not null primary key,
    state TEXT not null check ( state IN ('enqueuing','sending','sent','failed')),
    postfix_id text not null,
    status mta_status not null check ( state IN ('unknown','sent','deferred','bounced','expired')),
    status_text_id integer not null,
    domain_id integer not null ,
    created text not null,
    updated text not null,
    log text,
    foreign key (status_text_id) references mta_status_text(id),
    foreign key (domain_id) references domain(id)
);

create table if not exists post_request
(
    id text not null primary key,
    message text not null,
    created_on text not null
);

create table if not exists get_request(
    request_id text not null primary key,
    status
);

-- create table if not exists post_response(
--
-- );