# 🛠️ Customizable Job Board – Product Requirements

## 🎯 Goal

Create a personalized job board for a small group of users (starting with a friend seeking Event Coordinator roles). The board aggregates jobs from external sources via scraping, allows filtering, and links directly to job applications (bypassing LinkedIn when possible).

---

## 🧩 Core Features (MVP)

### 1. User Profiles / Preferences
- Basic user identity or config
- Preference options:
  - Job titles (e.g. "Event Coordinator")
  - Locations (e.g. Remote, Toronto)
  - Employment type (e.g. Full-time, Contract)
  - Industries or tags

### 2. Job Listing Feed
- Aggregated from scraped sources
- Fields per job:
  - Title
  - Company
  - Location
  - Type (Full-time, Contract, etc.)
  - Tags
  - Description preview
  - Apply URL (direct)
  - Source name
  - Posted date

### 3. Filtering UI (eventually)
- Filter by:
  - Keyword search
  - Location
  - Employment type
  - Tags/Industries

### 4. Scraper / Data Collector
- Scrapes jobs from:
  - Company sites
  - Smaller boards (e.g. WorkInCulture, ArtsWork, etc.)
- Stores jobs in DB or structured file

### 5. Personalized Feeds
- Each user sees a filtered feed based on saved preferences
- Could be based on a URL param or config at first

---

## 💅 Nice-to-Haves (Post-MVP)
- Email digests (e.g. daily/weekly updates)
- Bookmark/save jobs
- Application tracking (e.g. “Applied”, “Interviewed”)
- Admin panel for scraping jobs manually or scheduling updates
- Mobile responsive or PWA
- Integration with Airtable, Notion, Google Sheets

---

## 👥 User Roles
### Job Seeker
- View personalized job feed
- Filter jobs
- Save/bookmark jobs (later)

### Admin (You)
- Run scraper
- Manage source list
- Review/remove jobs (optional)

---

## 🗃️ Suggested Data Schema

### `users`
| Field | Type | Notes |
|-------|------|-------|
| id | UUID | Primary key |
| name | String | Optional |
| email | String | Optional |
| preferences | JSONB | Titles, locations, tags, etc. |

### `jobs`
| Field | Type | Notes |
|-------|------|-------|
| id | UUID | Primary key |
| title | String | Required |
| company | String | Required |
| location | String | City / remote |
| type | String | Full-time / Contract |
| tags | String[] | Skills, industries |
| description | Text | Optional preview |
| apply_url | String | Link to external site |
| source_name | String | Where it was scraped from |
| posted_at | Date | Scraped or original date |

### `user_saved_jobs`
| Field | Type | Notes |
|-------|------|-------|
| id | UUID | Primary key |
| user_id | UUID | Foreign key to users |
| job_id | UUID | Foreign key to jobs |
| status | String | e.g. "saved", "applied" |
| notes | Text | Optional |

---

## 🧪 Stack Recommendations

- **Backend:** Go or Node.js (your call!)
- **Frontend (later):** Next.js + TypeScript
- **Database:** SQLite for MVP, Postgres for scaling
- **Scraping:** Python (BeautifulSoup/Scrapy) or Go (Colly)
- **Deployment:** Render, Railway, or Fly.io

---

## 🔜 Next Steps

1. Define target job sites to scrape
2. Set up schema and DB
3. Write scraper for 1–2 sources
4. Expose a `/api/jobs` endpoint
5. Create simple frontend to display jobs
# job-board
