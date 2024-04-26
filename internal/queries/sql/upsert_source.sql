INSERT INTO sources(
    url,
    title,
    link,
    description,
    is_active,
    last_post_link,
    last_posted_at
) VALUES(
         :url,
         :title,
         :link,
         :description,
         :is_active,
         :last_post_link,
         :last_posted_at
) ON DUPLICATE KEY UPDATE
    title = VALUE(title),
    link = VALUE(link),
    description = VALUE(description),
    last_post_link = VALUE(last_post_link),
    last_posted_at = VALUE(last_posted_at),
    updated_at = now()
;
