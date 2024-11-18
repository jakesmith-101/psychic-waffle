package db

import "context"

func PostFuncs() error {
	var err error

	_, err = PgxPool.Exec(context.Background(), `
	CREATE OR REPLACE FUNCTION public.set_slug_from_name() RETURNS trigger
		LANGUAGE plpgsql
		AS $$
	DECLARE
		base_slug TEXT;
		final_slug TEXT;
		counter INTEGER := 1;
	BEGIN
		base_slug := NEW.slug;
		final_slug := base_slug;

		-- Loop to ensure uniqueness of the slug
		LOOP
			-- Check if the slug already exists in the table
			IF EXISTS (SELECT 1 FROM "my_table_name" WHERE slug = final_slug AND id != COALESCE(NEW.id, 0)) THEN
				-- If it exists, append a numeric suffix and increment the counter
				final_slug := base_slug || '-' || counter;
				counter := counter + 1;
			ELSE
				-- If it's unique, exit the loop
				EXIT;
			END IF;
		END LOOP;

		-- Set the unique slug to the 'slug' field of the NEW record
		NEW.slug := final_slug;
		RETURN NEW;
	END
	$$;
	`)
	if err != nil {
		return err
	}

	_, err = PgxPool.Exec(context.Background(), `
	CREATE TRIGGER set_slug_from_name
	BEFORE INSERT OR UPDATE
	ON "posts"
	FOR EACH ROW
	EXECUTE FUNCTION public.set_slug_from_name();
	`)
	if err != nil {
		return err
	}

	return err
}
