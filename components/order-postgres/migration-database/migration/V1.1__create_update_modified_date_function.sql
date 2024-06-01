CREATE OR REPLACE FUNCTION update_modified_at() RETURNS TRIGGER
    LANGUAGE plpgsql AS
$$
BEGIN
    NEW.modified_at := now();
    RETURN NEW;
END;
$$;