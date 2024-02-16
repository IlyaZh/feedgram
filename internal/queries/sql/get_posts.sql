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
FROM feedgram.posts
WHERE
    ($1::BIGINT IS NULL OR id = $1) AND
    ($2::BOOLEAN IS NULL OR has_read = $2)
ORDER BY created_at DESC
LIMIT $3