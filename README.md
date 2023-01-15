# Introduction
Blockchains, including Ethereum, are optimized for immutability and auditability.
An adverse side-effect of this optimization is that the data on Ethereum blockchain is not easily searchable.
The Ethereum Indexer application scans the blocks and listens for new blocks and stores the data in a relational database for efficient searching.
The application will allow easy search for transactions, tokens, account activity and interactions.

At present, the application supports Ethereum, and it can be extended to support other blockchain such as Bitcoin.

# Start Postgres
Reference: https://hub.docker.com/_/postgres
Set postgres password in docker/postgres-compose.yml
docker-compose -f docker/postgres-compose..yml up
Import schema in the database from postgres/db.sql

# Start Ethereum
docker-compose -f docker/eth-dev-compose.yml up
