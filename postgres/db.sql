-- Adminer 4.8.1 PostgreSQL 14.5 (Debian 14.5-1.pgdg110+1) dump

DROP TABLE IF EXISTS "block";
DROP SEQUENCE IF EXISTS "block_Id_seq";
CREATE SEQUENCE "block_Id_seq" INCREMENT  MINVALUE  MAXVALUE  CACHE ;

CREATE TABLE "public"."block" (
    "id" bigint DEFAULT nextval('"block_Id_seq"') NOT NULL,
    "hash" text NOT NULL,
    "parent_hash" text NOT NULL,
    "number" bigint NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "block_Hash" UNIQUE ("hash"),
    CONSTRAINT "block_Number" UNIQUE ("number"),
    CONSTRAINT "block_ParentHash" UNIQUE ("parent_hash"),
    CONSTRAINT "block_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "eth_transaction";
DROP SEQUENCE IF EXISTS "transaction_Id_seq";
CREATE SEQUENCE "transaction_Id_seq" INCREMENT  MINVALUE  MAXVALUE  CACHE ;

CREATE TABLE "public"."eth_transaction" (
    "id" bigint DEFAULT nextval('"transaction_Id_seq"') NOT NULL,
    "hash" text NOT NULL,
    "block_number" bigint NOT NULL,
    "fee" numeric NOT NULL,
    "gas" numeric NOT NULL,
    "gas_price" numeric NOT NULL,
    "is_contract_creation" boolean NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "transaction_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "eth_transaction_block_number" ON "public"."eth_transaction" USING btree ("block_number");


DROP TABLE IF EXISTS "indexer_position";
DROP SEQUENCE IF EXISTS indexer_position_id_seq;
CREATE SEQUENCE indexer_position_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 2 CACHE 1;

CREATE TABLE "public"."indexer_position" (
    "id" bigint DEFAULT nextval('indexer_position_id_seq') NOT NULL,
    "coin_type" smallint NOT NULL,
    "position" bigint NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "indexer_position_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "indexer_position_coin_type" ON "public"."indexer_position" USING btree ("coin_type");


DROP TABLE IF EXISTS "sequencer_position";
DROP SEQUENCE IF EXISTS block_position_id_seq;
CREATE SEQUENCE block_position_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."sequencer_position" (
    "id" bigint DEFAULT nextval('block_position_id_seq') NOT NULL,
    "coin_type" smallint NOT NULL,
    "position" bigint NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "block_position_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "transaction_payment";
DROP SEQUENCE IF EXISTS transaction_payment_id_seq;
CREATE SEQUENCE transaction_payment_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 9 CACHE 1;

CREATE TABLE "public"."transaction_payment" (
    "id" bigint DEFAULT nextval('transaction_payment_id_seq') NOT NULL,
    "transaction_id" bigint NOT NULL,
    "from" text NOT NULL,
    "to" text NOT NULL,
    "index" smallint NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "amount" numeric,
    CONSTRAINT "transaction_payment_id" PRIMARY KEY ("id"),
    CONSTRAINT "transaction_payment_transaction_id" UNIQUE ("transaction_id")
) WITH (oids = false);


ALTER TABLE ONLY "public"."eth_transaction" ADD CONSTRAINT "eth_transaction_block_number_fkey" FOREIGN KEY (block_number) REFERENCES block(number) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;

ALTER TABLE ONLY "public"."transaction_payment" ADD CONSTRAINT "transaction_payment_transaction_id_fkey" FOREIGN KEY (transaction_id) REFERENCES eth_transaction(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;

-- 2023-03-04 23:27:02.173384+00
