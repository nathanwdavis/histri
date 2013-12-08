
CREATE OR REPLACE FUNCTION histri.insert_event
    (
        t timestamp,
        et varchar(200),
        er varchar(200),
        d json
    ) RETURNS int AS $$

    INSERT INTO histri.events (timeutc, event_type, ext_ref, data)
        VALUES (t, et, er, d)
    RETURNING id;
$$ LANGUAGE SQL;
