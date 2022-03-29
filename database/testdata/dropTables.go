package testdata

import (
	"PowerShare/database"
	"log"
)

func CleanDB() {
	sqlStatement :=
		`DO $$ 
		  DECLARE 
			r RECORD;
		BEGIN
		  FOR r IN 
			(
			  SELECT table_name 
			  FROM information_schema.tables 
			  WHERE table_schema=current_schema()
			) 
		  LOOP
			 EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.table_name) || ' CASCADE';
		  END LOOP;
		END $$ ;`

	_, err := database.DB.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	log.Println("DB cleaned")
}
