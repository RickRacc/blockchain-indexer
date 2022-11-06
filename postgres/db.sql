-- Adminer 4.8.1 PostgreSQL 14.5 (Debian 14.5-1.pgdg110+1) dump

-- connect "bonotans";

DROP TABLE IF EXISTS "block";
DROP SEQUENCE IF EXISTS "block_Id_seq";
CREATE SEQUENCE "block_Id_seq";

CREATE TABLE "public"."block" (
                                  "Id" bigint DEFAULT nextval('"block_Id_seq"') NOT NULL,
                                  "Hash" text NOT NULL,
                                  "ParentHash" text NOT NULL,
                                  "Number" bigint NOT NULL,
                                  CONSTRAINT "block_Hash" UNIQUE ("Hash"),
                                  CONSTRAINT "block_Number" UNIQUE ("Number"),
                                  CONSTRAINT "block_ParentHash" UNIQUE ("ParentHash"),
                                  CONSTRAINT "block_pkey" PRIMARY KEY ("Id")
) WITH (oids = false);


DROP TABLE IF EXISTS "eth_transaction";
DROP SEQUENCE IF EXISTS "transaction_Id_seq";
CREATE SEQUENCE "transaction_Id_seq" ;

CREATE TABLE "public"."eth_transaction" (
                                            "Id" bigint DEFAULT nextval('"transaction_Id_seq"') NOT NULL,
                                            "Hash" text NOT NULL,
                                            "BlockNumber" bigint NOT NULL,
                                            "Fee" integer NOT NULL,
                                            "Gas" integer NOT NULL,
                                            "GasPrice" integer NOT NULL,
                                            "IsContractCreation" boolean NOT NULL,
                                            "CreatedAt" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                            "UpdatedAt" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                            CONSTRAINT "transaction_pkey" PRIMARY KEY ("Id")
) WITH (oids = false);


DROP TABLE IF EXISTS "indexer_position";
DROP SEQUENCE IF EXISTS indexer_position_id_seq;
CREATE SEQUENCE indexer_position_id_seq;

CREATE TABLE "public"."indexer_position" (
                                             "id" bigint DEFAULT nextval('indexer_position_id_seq') NOT NULL,
                                             "coin_type" smallint NOT NULL,
                                             "position" bigint NOT NULL,
                                             CONSTRAINT "indexer_position_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "indexer_position_coin_type" ON "public"."indexer_position" USING btree ("coin_type");


-- 2022-10-13 19:48:09.785653+00