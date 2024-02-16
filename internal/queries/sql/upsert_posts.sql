INSERT INTO feedgram.posts(
    source_id,
    title,
    body,
    link,
    author,
    additional_info,
    posted_at
)VALUES(
        :source_id
        :title,
        :body,
        :link,
        :author,
        :has_read,
        :additional_info,
        :posted_at
) ON CONFLICT (link) DO NOTHING