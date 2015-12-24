drop keyspace observr;
create keyspace observr WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 3};

use observr;

create table visits(
    id                      UUID,
    remote_addr             VARCHAR,
    method                  VARCHAR,
    user_agent              VARCHAR,
    status_code             INT,
    host                    VARCHAR,
    protocol                VARCHAR,
    path                    VARCHAR,
    data                    VARCHAR,
    headers                 VARCHAR,
    query_string            VARCHAR,
    location_id             UUID,
    language_id             UUID,
    group_id                UUID,
    url_id                  UUID,
    project_id              UUID,
    visitor_id              UUID,
    visitor_device_id       UUID,
    visitor_address_id      UUID,
    referer_id              UUID,
    device_id               UUID,
    browser_id              UUID,
    platform_id             UUID,
    created_at              TIMEUUID,
    PRIMARY KEY (id, created_at)
) WITH CLUSTERING ORDER BY (created_at DESC);


create table urls(
    id                      UUID PRIMARY KEY,
    key                     VARCHAR,
    path                    VARCHAR,
    host                    VARCHAR,
    project_id              UUID,
    group_id                UUID,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT,
    seen_uniq_count         INT,
);


create table groups(
    id                      UUID PRIMARY KEY,
    key                     VARCHAR,
    project_id              UUID,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT,
);

create table referers(
    id                      UUID PRIMARY KEY,
    host                    VARCHAR,
    path                    VARCHAR,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT
);


create table visitor_devices(
    id                      UUID PRIMARY KEY,
    visitor_id              UUID,
    visitor_address_id      UUID,
    device_id               UUID,
    browser_id              UUID,
    platform_id             UUID,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT
);

create table devices(
    id                      UUID PRIMARY KEY,
    key                     VARCHAR,
    value                   VARCHAR,
    seen_count              INT,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    device_family_id        UUID
);

create table languages(
    id                      UUID PRIMARY KEY,
    key                     VARCHAR,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT
);

create table countries(
    id                      UUID PRIMARY KEY,
    code                    VARCHAR,
);

create table locations(
    id                      UUID PRIMARY KEY,
    country_id              UUID,
    city                    VARCHAR,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT
);

create table browsers(
    id                      UUID PRIMARY KEY,
    key                     VARCHAR,
    version                 VARCHAR,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT
);

create table platforms(
    id                      UUID PRIMARY KEY,
    key                     VARCHAR,
    version                 VARCHAR,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT
);

create table visitors(
    id                      UUID PRIMARY KEY,
    key                     VARCHAR,
    value                   TEXT,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT
);

create table visitor_addresses(
    id                      UUID PRIMARY KEY,
    key                     VARCHAR,
    visitor_id              UUID,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    seen_count              INT
);
