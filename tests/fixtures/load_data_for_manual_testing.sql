INSERT INTO
    tags_spaces (name, id)
VALUES
   ('first-tags-space', '2df72cd2-5c77-48bd-8f24-985a97c75923'),
   ('second-tags-space', 'bc7c9821-9bee-43ca-8135-3827ec97b44e');

INSERT INTO
    tags (name, id, tags_space_id)
VALUES
    ('ts1-tag1', 'e0e556ea-f624-4a7c-b1a8-9de07b495a0f', '2df72cd2-5c77-48bd-8f24-985a97c75923'),
    ('ts1-tag2', 'bc25fd31-3246-4b25-a6fb-7a3e148eadf5', '2df72cd2-5c77-48bd-8f24-985a97c75923'),
    ('ts1-tag3', '723b2aca-c5bc-43ea-9b79-a5e7ffa04398', '2df72cd2-5c77-48bd-8f24-985a97c75923');


INSERT INTO
    tags (name, id, tags_space_id)
VALUES
    ('ts2-tag1', 'd7c661ea-f94b-4a42-8ea2-a59ca48331b9', 'bc7c9821-9bee-43ca-8135-3827ec97b44e'),
    ('ts2-tag2', '75741b9a-0d3a-4311-be37-04d8f67540d0', 'bc7c9821-9bee-43ca-8135-3827ec97b44e'),
    ('ts2-tag3', 'fc8dd0f5-4dfc-445d-9875-a6503be814cd', 'bc7c9821-9bee-43ca-8135-3827ec97b44e');

