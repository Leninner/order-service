-- Create the order user with password authentication
CREATE ROLE order_role WITH LOGIN PASSWORD 'pa55word';

-- Grant necessary privileges to order user
GRANT CONNECT ON DATABASE order_db TO order_role;
GRANT USAGE ON SCHEMA public TO order_role;
GRANT CREATE ON SCHEMA public TO order_role;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO order_role;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO order_role;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO order_role;

-- Set default privileges for future objects
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO order_role;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO order_role;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON FUNCTIONS TO order_role;

-- Create the citext extension
CREATE EXTENSION IF NOT EXISTS citext;

-- Create pg_stat_statements extension for query monitoring
CREATE EXTENSION IF NOT EXISTS pg_stat_statements; 