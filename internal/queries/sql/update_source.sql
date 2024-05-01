UPDATE sources
SET 
    last_sync_at = COALESCE(:last_sync_at, last_sync_at),
    last_post_link = :last_post_link,
    last_posted_at = COALESCE(:last_posted_at, last_posted_at),
    updated_at = now()
WHERE id = :id
;
