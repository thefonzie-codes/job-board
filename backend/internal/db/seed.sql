INSERT INTO jobs (
    title, 
    company, 
    location, 
    type, 
    tags, 
    description, 
    apply_url, 
    source_name,
    posted_at
)
VALUES 
(
    'Event Coordinator',
    'Local Arts Org',
    'Toronto, ON',
    'Contract',
    'Events,Arts',
    'Coordinate large-scale events...',
    'https://artsorg.ca/jobs/123',
    'artsorg.ca',
    DATETIME('now', '-2 days')
),
(
    'Community Manager',
    'Creative Collective',
    'Remote',
    'Full-time',
    'Community,Remote,Creative',
    'Grow and support the community...',
    'https://creativecollective.org/jobs/456',
    'creativecollective.org',
    DATETIME('now', '-1 days')
);

