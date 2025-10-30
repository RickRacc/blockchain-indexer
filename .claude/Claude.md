# Blockchain Indexer - Project Context

## Project Overview

The **Ethereum Blockchain Indexer** is a distributed data processing system that solves the fundamental challenge of making immutable blockchain data searchable and queryable. It continuously scans Ethereum blockchain blocks, extracts transaction data, and stores it in a PostgreSQL database for efficient querying.

### Key Purpose
Rather than reading directly from the blockchain (optimized for immutability, not search), this system:
- Continuously scans Ethereum blockchain blocks
- Listens for new blocks in real-time
- Extracts and normalizes transaction data
- Stores everything in PostgreSQL for efficient searching
- Enables fast queries on transactions, tokens, account activity, and address interactions

### Architecture Design
The system is **blockchain-agnostic** with a pluggable architecture supporting Ethereum, with extensibility for Bitcoin and other chains.

---

## Technical Architecture

### Layered Architecture Pattern

```
Entry Point (cmd/)
    ↓
Dependency Injection (di/)
    ↓
Domain Logic (indexer/) → Blockchain Client (blockchain/)
    ↓
Data Access (repository/) ← Models (model/)
    ↓
PostgreSQL Database (postgres/)
```

### Core Components

#### 1. Blockchain Interface Layer (`blockchain/`)
- Abstract `Blockchain` interface with two core methods:
  - `GetTip()` - retrieves latest block number
  - `GetBlock()` - fetches complete block with transactions
- Ethereum implementation using `go-ethereum` (geth) client at [blockchain/eth/ethereum.go](blockchain/eth/ethereum.go)

#### 2. Indexer Engine (`indexer/indexer.go`)
- Orchestrates the indexing workflow
- Manages concurrent block processing with **100 parallel goroutines**
- Implements fork detection and recovery logic
- Two-phase sequencing system

#### 3. Data Model (`model/`)
Key models:
- Transaction hierarchy: `Transaction` (interface) → `BaseTransaction` → `EthTransaction`
- Block model with full transaction tree
- Position tracking for resumable indexing
- Support for arbitrary payment metadata

#### 4. Repository Pattern (`repository/`)
Five repository interfaces providing abstracted data access:
- `BlockRepository` - block CRUD and queries
- `TransactionRepository` - transaction storage
- `TransactionPaymentRepository` - payment details
- `IndexerPositionRepository` - resume points
- `SequencerPositionRepository` - fork recovery state

Custom SQL column mapping for maintainability.

#### 5. Dependency Injection (`di/`)
- Dual approach: Manual DI and Google Wire integration
- Configurable client initialization
- Connection pooling management (3 concurrent connections)

---

## Advanced Technical Features

### A. Concurrent Block Indexing
```go
numRoutines := 100  // 100 goroutines for parallel processing
```
- Producer-consumer pattern with buffered channels
- Up to 100 concurrent goroutines fetching and processing blocks
- Channel-based flow control prevents resource exhaustion

### B. Fork Handling & Recovery Algorithm
```
1. Track sequencer position (last verified block)
2. Detect fork: ParentHash mismatch between DB and blockchain
3. Walk back incrementally until finding matching hash
4. Delete all blocks after fork point
5. Restart re-fetching from fork position
```
- Critical for blockchain reorg handling
- Prevents data corruption from chain forks
- O(n) complexity where n = fork depth

### C. Position-Based Resumable Indexing
- **IndexerPosition**: tracks last processed block per coin type
- **SequencerPosition**: tracks verified block per coin type
- Enables resume from exact position if process crashes
- Multi-blockchain support via `coin_type` enum (BTC=0, ETH=60)

### D. BigInt Arithmetic Support
- Direct `math/big.Int` support for transaction amounts
- Prevents precision loss on large numerical values (fees, gas, amounts)
- Serialized as byte arrays in PostgreSQL
- Handles Wei precision (10^-18 ETH)

---

## Database Design

### Schema Structure (`postgres/db.sql` - 93 lines)

| Table | Purpose | Key Features |
|-------|---------|--------------|
| `block` | Core blockchain blocks | UNIQUE constraints on hash, number, parent_hash |
| `eth_transaction` | Transaction metadata | B-tree index on block_number |
| `transaction_payment` | Payment details | Foreign key with CASCADE delete |
| `indexer_position` | Resume checkpoint | Index on coin_type |
| `sequencer_position` | Fork recovery state | Index on coin_type |

### Performance Optimizations
- Primary keys on all tables (bigint auto-increment)
- Foreign key constraints with CASCADE delete for referential integrity
- Compound indexing on `coin_type` for multi-blockchain support
- Timestamp auditing (created_at, updated_at) on all entities

---

## Technology Stack

### Core Dependencies
- **Language**: Go 1.19
- **Blockchain Client**: `github.com/ethereum/go-ethereum` v1.10.23
- **Database**: PostgreSQL with `lib/pq` driver
- **Configuration**: Koanf (multi-format config with environment override)
- **Testing**: Testify (assertions + suite pattern)
- **DI Framework**: Google Wire (code-gen dependency injection)

### Development Infrastructure
- Docker Compose for PostgreSQL + Adminer
- Support for Ethereum test networks (Goerli via QuickNode RPC)
- Automated schema management

---

## Testing Infrastructure

### Test Coverage
6 test files covering:
- Block repository CRUD and queries
- Transaction persistence
- Transaction payment handling
- Indexer position tracking (resume functionality)
- Sequencer position tracking (fork recovery)

### Test Suite Architecture
```go
BaseTestSuite {
  - Automatic database setup/teardown
  - CASCADE truncation between tests
  - Fixture data for reproducible tests
  - Suite-based organization (Testify)
}
```

Located at [test/base_test_suite.go](test/base_test_suite.go)

---

## Key Metrics

### Code Statistics
- **1,491 lines** of production Go code
- **38+ commits** showing iterative development
- **5 test suites** with comprehensive coverage
- **~250-300 LOC** per core module (well-organized)

### Performance Characteristics
- **100 parallel goroutines** for block processing
- **3 concurrent database connections** with pooling
- **O(n) fork detection** algorithm (n = fork depth)
- **Sub-block precision** recovery via dual position tracking
- **2+ blockchain types** supported (Ethereum implemented, Bitcoin-ready)

---

## Critical File Locations

### Core Implementation
- [indexer/indexer.go](indexer/indexer.go) - Main indexing engine
- [blockchain/eth/ethereum.go](blockchain/eth/ethereum.go) - Ethereum client
- [blockchain/blockchain.go](blockchain/blockchain.go) - Blockchain interface

### Repository Layer
- [repository/block_repository.go](repository/block_repository.go)
- [repository/transaction_repository.go](repository/transaction_repository.go)
- [repository/transaction_payment_repository.go](repository/transaction_payment_repository.go)
- [repository/indexer_position_repository.go](repository/indexer_position_repository.go)
- [repository/sequencer_position_repository.go](repository/sequencer_position_repository.go)

### Data Models
- [model/block.go](model/block.go)
- [model/transaction.go](model/transaction.go)
- [model/eth_transaction.go](model/eth_transaction.go)
- [model/transaction_payment.go](model/transaction_payment.go)
- [model/indexer_position.go](model/indexer_position.go)
- [model/sequencer_position.go](model/sequencer_position.go)

### Infrastructure
- [di/di.go](di/di.go) - Dependency injection setup
- [postgres/db.sql](postgres/db.sql) - Database schema
- [config/config.yaml](config/config.yaml) - Configuration
- [docker-compose.yml](docker-compose.yml) - Local dev environment

### Testing
- [test/base_test_suite.go](test/base_test_suite.go) - Test infrastructure
- [repository/block_repository_test.go](repository/block_repository_test.go)
- [repository/transaction_repository_test.go](repository/transaction_repository_test.go)
- [repository/transaction_payment_repository_test.go](repository/transaction_payment_repository_test.go)
- [repository/indexer_position_repository_test.go](repository/indexer_position_repository_test.go)
- [repository/sequencer_position_repository_test.go](repository/sequencer_position_repository_test.go)

---

## Complex Problems Solved

1. **The Search Problem**: Raw blockchain data is immutable and sequential - this system makes it relational and queryable

2. **Concurrent Indexing**: Processing hundreds of blocks in parallel without data corruption using goroutines and channels

3. **Fork Recovery**: Detecting and recovering from blockchain reorganizations (critical for correctness)

4. **Position Tracking**: Enabling restartable indexing that doesn't re-process blocks

5. **Multi-Blockchain Support**: Abstracted architecture supporting different chains with same infrastructure

6. **Precision Arithmetic**: Handling Ethereum's large numerical values (Wei units ~10^18)

---

## Architectural Strengths

- **Pluggable Blockchain Implementations**: Easy to add Bitcoin, Litecoin, etc.
- **Interface-Driven Design**: Clear separation of concerns
- **Testable**: All layers independently testable with test suites
- **Configurable**: YAML + environment variables for flexibility
- **Production-Ready**: Error handling, logging, connection pooling
- **Resumable**: Can restart from exact position without duplication

---

## Git History Context

### Recent Commits
- `83ab0b8` - Initial support for fork handling during indexing
- `cab374e` - Initial support for sequencing in indexer
- `0495c88` - Updates to use coin types in tests
- `5ffdb62` - Initial updates for sequencing blocks and handling forks
- `96e79ca` - Support for storing indexer position and prevent docker files from build path

### Current Branch
- **Branch**: main
- **Status**: Clean working directory

---

## Development Commands

### Running the Indexer
```bash
go run cmd/main.go
```

### Database Setup
```bash
docker-compose up -d  # Start PostgreSQL + Adminer
```

### Running Tests
```bash
go test ./...
```

### Configuration
- Config file: [config/config.yaml](config/config.yaml)
- Environment variables override config values
- Database connection pooling configurable

---

## Common Development Patterns

### Adding a New Blockchain
1. Implement `Blockchain` interface in `blockchain/` directory
2. Add new coin type to `model/coin_type.go`
3. Update DI configuration in `di/di.go`
4. No changes needed to indexer, repository, or database layer

### Adding a New Repository
1. Define interface in `repository/` directory
2. Implement SQL queries with custom column mapping
3. Add to DI wire setup
4. Create test suite extending `BaseTestSuite`

### Extending Transaction Model
1. Create new transaction type implementing `Transaction` interface
2. Extend database schema in `postgres/db.sql`
3. Add repository methods for new fields
4. Update tests with new fixtures

---

## Future Enhancement Ideas

- [ ] Add Bitcoin blockchain support
- [ ] Implement GraphQL API layer for queries
- [ ] Add metrics/monitoring (Prometheus)
- [ ] Implement WebSocket subscriptions for real-time updates
- [ ] Add indexing for smart contract events/logs
- [ ] Support for ERC-20 token transfer tracking
- [ ] Add caching layer (Redis) for frequently accessed data
- [ ] Implement backfilling strategy for historical blocks
- [ ] Add admin API for managing indexer state

---

## Resume/Portfolio Highlights

### Key Achievements
- Engineered distributed blockchain indexer with 100 concurrent goroutines
- Implemented fork detection algorithm with automatic recovery
- Designed resumable indexing system with dual position tracking
- Built multi-blockchain architecture with pluggable implementations
- Comprehensive test coverage with suite-based testing

### Technical Skills Demonstrated
- Go concurrency patterns (goroutines, channels)
- Distributed systems design
- Database design and optimization
- Interface-driven architecture
- Dependency injection
- Test-driven development
- Docker containerization

---

## Notes for Future Development

### Known Considerations
- Currently uses `go-ethereum` v1.10.23 - may need updates for newer Ethereum features
- Connection pool size (3) may need tuning for production loads
- Fork detection walks backward linearly - could be optimized with binary search for deep forks
- No retry logic for transient RPC failures - should add exponential backoff

### Design Decisions
- Chose repository pattern over ORM for explicit SQL control
- Dual position tracking (indexer + sequencer) provides clean separation between fetch and verify phases
- BigInt stored as bytes rather than numeric for maximum precision
- Goroutine count (100) chosen for balance between throughput and resource usage

---

*Last Updated: 2025-10-28*
*Codebase Version: Commit 83ab0b8*
