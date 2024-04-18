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
    (:id::BIGINT IS NULL OR id > :id) AND
    (:is_active::TINYINT IS NULL OR is_active = :is_active) AND
    deleted_at IS NULL
ORDER BY created_at ASC
LIMIT :limit