# Introduction
Blockchains, including Ethereum, are optimized for immutability and auditability.
An adverse side-effect of this optimization is that the data on Ethereum blockchain is not easily searchable.
The Ethereum Indexer application scans the blocks and listens for new blocks and stores the data in a relational database for efficient searching.
The application will allow easy search for transactions, tokens, account activity and interactions.

At present, the application supports Ethereum, and it can be extended to support other blockchain such as Bitcoin.

# Start Postgres
docker-compose -f postgres/stack.yml up