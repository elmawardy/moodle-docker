#!/usr/bin/env bash
echo "Restoring database"
psql -U postgres -c "CREATE DATABASE moodle"
psql -U postgres -d moodle -1 -f /tmp/database.sql
psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE moodle TO postgres"
echo "Database restored successfully"