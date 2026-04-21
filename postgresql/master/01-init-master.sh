#!/bin/bash
set -e

echo ">>> Setting up replication user and pgvector on MASTER..."

# Expand variables dulu ke local vars yang aman
REPL_USER_VAL="${POSTGRES_REPLICATION_USER}"
REPL_PASS_VAL="${POSTGRES_REPLICATION_PASSWORD}"

psql -v ON_ERROR_STOP=1 \
     --username "$POSTGRES_USER" \
     --dbname "$POSTGRES_DB" \
     <<-EOSQL
    -- Buat replication user
    DO \$\$
    BEGIN
        IF NOT EXISTS (
            SELECT FROM pg_catalog.pg_roles WHERE rolname = '$REPL_USER_VAL'
        ) THEN
            CREATE ROLE "$REPL_USER_VAL"
                WITH REPLICATION LOGIN
                ENCRYPTED PASSWORD '$REPL_PASS_VAL';
        END IF;
    END
    \$\$;

    -- Enable pgvector extension
    CREATE EXTENSION IF NOT EXISTS vector;

    -- Grant schema usage ke replication user
    GRANT USAGE ON SCHEMA public TO "$REPL_USER_VAL";
EOSQL

echo ">>> Checking pg_hba.conf for existing replication entry..."

# Hanya tambahkan jika belum ada, hindari duplikat
if ! grep -q "replication.*${REPL_USER_VAL}" "$PGDATA/pg_hba.conf"; then
    cat >> "$PGDATA/pg_hba.conf" <<EOF

# Allow replication connections from slaves (auto-added)
host    replication     ${REPL_USER_VAL}    0.0.0.0/0    md5
host    all             ${REPL_USER_VAL}    0.0.0.0/0    md5
EOF
    echo ">>> pg_hba.conf updated."
else
    echo ">>> pg_hba.conf entry already exists, skipping."
fi

# Reload config agar pg_hba.conf langsung aktif
psql -v ON_ERROR_STOP=1 \
     --username "$POSTGRES_USER" \
     --dbname "$POSTGRES_DB" \
     -c "SELECT pg_reload_conf();"

echo ">>> Master replication setup complete."