CREATE TABLE IF NOT EXISTS facts_resource.fact
(
    fact_id   INT GENERATED ALWAYS AS IDENTITY,
    fact_text VARCHAR,
    PRIMARY KEY (fact_id)
);