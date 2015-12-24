drop keyspace observr;
create keyspace observr WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};

use observr;

create table visits(
    id                      UUID,
    host                    VARCHAR,
    path                    VARCHAR,
    remote_addr             VARCHAR,
    method                  VARCHAR,
    user_agent              VARCHAR,
    status_code             INT,
    protocol                VARCHAR,
    data                    VARCHAR,
    headers                 VARCHAR,
    cookies                 VARCHAR,
    referer                 VARCHAR,
    query_string            VARCHAR,
    project_id              UUID,
    created_at              TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE INDEX ON visits(created_at);
CREATE INDEX ON visits(method);
CREATE INDEX ON visits(status_code);
CREATE INDEX ON visits(project_id);

insert into visits (id, created_at, project_id) values (00000000-0000-0000-0000-000000000000, '2013-01-01 00:05+0000', 00000000-0000-0000-0000-000000000000);
insert into visits (id, created_at, method, project_id) values (b4d1fc4b-712e-46d9-848d-d38a728be757, '2013-01-01 01:05+0000', 'POST', e4acccf6-48cf-4d21-899a-59b8ea0eb72e);

create table tags(
    id                      UUID,
    key                     VARCHAR,
    value                   VARCHAR,
    data                    VARCHAR,
    created_at              TIMESTAMP,
    first_seen              TIMESTAMP,
    last_seen               TIMESTAMP,
    project_id              UUID,
    seen_count              INT,
    PRIMARY KEY (id)
);

CREATE INDEX ON tags(key);
CREATE INDEX ON tags(value);
CREATE INDEX ON tags(project_id);


create table visit_tags(
    visit_id                UUID,
    tag_id                  UUID,
    created_at              TIMESTAMP,
    PRIMARY KEY (visit_id, tag_id, created_at)
);


create table group_tags(
    src_tag_id              UUID,
    dst_tag_id              UUID,
    created_at              TIMESTAMP,
    PRIMARY KEY (src_tag_id, dst_tag_id, created_at)
);
