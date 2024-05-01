SELECT
    id,
    url,
    title,
    link,
    description,
    is_active,
    last_post_link,
    last_posted_at,
    created_at,
    updated_at,
    deleted_at,
    last_posted_at
FROM sources
WHERE
    COALESCE(?, id) = id AND
    COALESCE(?, is_active) = is_active AND
    deleted_at IS NULL
ORDER BY created_at ASC
LIMIT ?