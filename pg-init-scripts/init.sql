CREATE DATABASE test_task;
CREATE USER test_task WITH PASSWORD 'test_task';
GRANT ALL PRIVILEGES ON DATABASE test_task TO test_task;

\c test_task
GRANT USAGE ON SCHEMA public TO test_task;
GRANT CREATE ON SCHEMA public TO test_task;