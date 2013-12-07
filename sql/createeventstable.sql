
CREATE TABLE histri.events (
    id serial PRIMARY KEY,
    timeutc timestamp NOT NULL,
    event_type varchar(200) NOT NULL CHECK (event_type <> ''),
    ext_ref varchar(200),
    data json
);
CREATE INDEX histri_events_timeutc_idx 
    ON histri.events (timeutc ASC,
                      event_type NULLS LAST,
                      ext_ref NULLS LAST)
