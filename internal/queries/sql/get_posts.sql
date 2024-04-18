SELECT
    id,
    source_id,
    title,
    body,
    link,
    author,
    has_read,
    additional_info,
    created_at,
    posted_at,
    updated_at
FROM posts
WHERE
    (:id::BIGINT IS NULL OR id = :id) AND
    (:has_read::TINYINT IS NULL OR has_read = :has_read)
ORDER BY created_at DESC
LIMIT :limit