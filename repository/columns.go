package repository

const INDEXER_POSITION_INSERT_COLS = "coin_type, position"
const INDEXER_POSITION_SELECT_COLS = "id, coin_type, position, created_at, updated_at"

const BLOCK_INSERT_COLS = "hash, parent_hash, number"
const BLOCK_SELECT_COLS = "id, hash, parent_hash, number, created_at, updated_at"

const TRANSACTION_INSERT_COLS = "hash, block_number, fee, gas, gas_price, is_contract_creation"
const TRANSACTION_SELECT_COLS = "id, hash, block_number, fee, gas, gas_price, is_contract_creation, created_at, updated_at"

const TRANSACTION_PAYMENT_INSERT_COLS = "transaction_id, from, to, index, amount"
const TRANSACTION_PAYMENT_SELECT_COLS = "id, transaction_id, from, to, index, amount, created_at, updated_at"
