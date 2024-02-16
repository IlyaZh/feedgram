SELETC
    id,
    url,
    title,
    link,
    description,
    is_active,
    created_at,
    updated_at,
    deleted_at
FROM feedgram.soures
WHERE
    ($1::BIGINT IS NULL OR id > $1) AND
    ($2::BOOLEAN IS NULL OR is_active = $2) AND
    deleted_at IS NULL
ORDER BY created_at ASC
LIMIT $3