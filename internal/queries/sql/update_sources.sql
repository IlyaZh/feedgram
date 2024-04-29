UPDATE sources
SET 
    last_post_link = :last_post_link,
    last_posted_at = COALESCE(:last_posted_at, last_posted_at),
    updated_at = now()
WHERE id = :id
;
