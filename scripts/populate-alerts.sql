-- Adjust the function definition if necessary
CREATE OR REPLACE FUNCTION random_string(length INT) RETURNS TEXT AS $$
BEGIN
RETURN (
    SELECT string_agg(substring('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789' FROM (random() * 62)::int + 1 FOR 1), '')
    FROM generate_series(1, length)
);
END;
$$ LANGUAGE plpgsql;

-- Ensure you have a sequence for generating incremental IDs
CREATE SEQUENCE IF NOT EXISTS pseudo_uuid_seq START 1;

-- Populate the table with random values and controlled create_time and pseudo-UUID
DO $$
    DECLARE
i INT := 0;
        baseTime TIMESTAMP := NOW() - INTERVAL '30 days';
        timeIncrement INTERVAL := '30 days'::INTERVAL / 100;
BEGIN
        WHILE i < 100 LOOP
                INSERT INTO alerts (id, description, create_time, resolve_time, severity, floor, zone, source, federation, subsystem, ack_time, ack_author_name, ack_author_email, ack_author_id)
                VALUES (
                           uuid_generate_v4(), -- Generate a true UUID
                           random_string(10),
                           baseTime + (i * timeIncrement), -- Increment time for each record
                           null,
                           CASE
                               WHEN random() < 0.25 THEN 9 -- INFO
                               WHEN random() < 0.5 THEN 13 -- WARNING
                               WHEN random() < 0.75 THEN 17 -- SEVERE
                               ELSE 21 -- LIFE_SAFETY
                               END,
                           'Room 1',
                           'Room 1',
                           random_string(5),
                           random_string(5),
                           random_string(5),
                           null,
                           null,
                           null,
                           null
                       );
                i := i + 1;
END LOOP;
END $$;

select *
from alerts;

delete
from alerts
where id in (select id
             from alerts
             order by create_time
    limit 1000);

