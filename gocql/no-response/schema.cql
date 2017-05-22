CREATE KEYSPACE ks WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1}  AND durable_writes = true;

CREATE TYPE ks.url (
    scheme text,
    host text,
    path text,
    query text,
    fragment text
);
