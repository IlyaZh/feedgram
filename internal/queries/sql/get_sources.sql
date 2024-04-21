SELETC
    s.id,
    s.url,
    s.title,
    s.link,
    s.description,
    s.is_active,
    s.created_at,
    s.updated_at,
    s.deleted_at,
    p.last_posted_at
FROM feedgram.soures AS s
LEFT JOIN feedgram.posts as p
ON s.id = p.source_id
WHERE
    (:id::BIGINT IS NULL OR s.id > :id) AND
    (:is_active::TINYINT IS NULL OR s.is_active = :is_active) AND
    s.deleted_at IS NULL
ORDER BY s.created_at ASC
LIMIT :limit