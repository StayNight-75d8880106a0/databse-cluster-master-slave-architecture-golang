#!/bin/bash
set -e

PGDATA="/var/lib/postgresql/data"
MASTER_HOST="${MASTER_HOST:-crime_master_db}"
MASTER_PORT="5432"

echo ">>> [SLAVE] Waiting for master at $MASTER_HOST:$MASTER_PORT..."

until PGPASSWORD="$REPL_PASS" pg_isready \
        -h "$MASTER_HOST" \
        -p "$MASTER_PORT" \
        -U "$REPL_USER"; do
    echo "    Master not ready yet, retrying in 3s..."
    sleep 3
done

echo ">>> [SLAVE] Master is ready."

if [ -s "$PGDATA/PG_VERSION" ]; then
    echo ">>> [SLAVE] Data directory exists, skipping pg_basebackup."
else
    echo ">>> [SLAVE] Cleaning data directory contents..."

    # Hapus ISI folder, bukan folder-nya sendiri
    # (folder adalah mount point Docker, tidak bisa dihapus)
    find "$PGDATA" -mindepth 1 -delete

    echo ">>> [SLAVE] Running pg_basebackup from master..."

    PGPASSWORD="$REPL_PASS" pg_basebackup \
        -h "$MASTER_HOST" \
        -p "$MASTER_PORT" \
        -U "$REPL_USER" \
        -D "$PGDATA" \
        --wal-method=stream \
        --write-recovery-conf \
        --progress \
        --verbose

    echo ">>> [SLAVE] pg_basebackup complete."

    APP_NAME="$(hostname)"
    cat >> "$PGDATA/postgresql.auto.conf" <<EOF

# Primary connection string (set by slave init script)
primary_conninfo = 'host=$MASTER_HOST port=$MASTER_PORT user=$REPL_USER password=$REPL_PASS application_name=$APP_NAME'
EOF

    touch "$PGDATA/standby.signal"

    chown -R postgres:postgres "$PGDATA"
    chmod 700 "$PGDATA"

    echo ">>> [SLAVE] Standby configuration written."
fi

echo ">>> [SLAVE] Starting PostgreSQL in hot standby mode..."

exec gosu postgres postgres \
    -D "$PGDATA" \
    -c hot_standby=on \
    -c hot_standby_feedback=on