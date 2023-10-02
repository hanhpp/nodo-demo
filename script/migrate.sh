#!/bin/bash

# Define the name of the PostgreSQL container
name="demo-db"

# Find the container name
CONTAINER_NAME=$(docker ps --format '{{.ID}} {{.Names}}' | grep "$name" | awk '{print $2}')


# Define the name of the database
DATABASE_NAME="demo-backend"

# Define the path to your SQL migration script
MIGRATION_SCRIPT="./schema/stock.sql"

# Run the migration script inside the PostgreSQL container
docker exec -i $CONTAINER_NAME psql -U postgres -d $DATABASE_NAME < $MIGRATION_SCRIPT

# Check the exit code to see if the migration was successful
if [ $? -eq 0 ]; then
  echo "Migration successful"
else
  echo "Migration failed"
fi
