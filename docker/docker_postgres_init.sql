-- Check if the database exists
SELECT datname FROM pg_database WHERE datname = 'aggreflow-db';

-- If the database doesn't exist, create it
DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'aggreflow-db') THEN
    CREATE DATABASE aggreflow-db;
  END IF;
END $$;
