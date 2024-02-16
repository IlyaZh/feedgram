INSERT INTO feedgram.soures(
    url,
    title,
    link,
    description,
    is_active
) VALUES(
         :url,
         :title,
         :link,
         :description,
         :is_active
) ON CONFLICT (url) DO UPDATE SET
    title = excluded.title,
    link = excluded.link,
    description = excluded.description,
    updated_at = now()
;
