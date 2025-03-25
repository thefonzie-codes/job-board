## 🚀 Setup and Run Instructions

To get the job board up and running on your local machine, follow these steps:

### Prerequisites
- **Go**: Ensure you have Go (version 1.24.0 or later) installed. [Download here](https://go.dev/dl/).
- **SQLite**: The project uses SQLite via `go-sqlite3`. No separate installation is needed as it’s included in the dependencies.

### Steps
1. **Clone the Repository**
   ```bash
   git clone https://github.com/thefonzie-codes/job-board.git
   cd job-board
   ```

2. **Navigate to the Backend Directory**
   ```bash
   cd backend
   ```

3. **Install Dependencies**
   Run the following command to download the required Go modules (e.g., `go-sqlite3`):
   ```bash
   go mod tidy
   ```

4. **Set Up the Database**
   - The app uses a SQLite database (`jobboard.db`).
   - The schema and seed data are applied automatically from:
     - `internal/db/schema.sql` (table structure)
     - `internal/db/seed.sql` (sample data)
   - Ensure these files exist in the `internal/db/` directory.

5. **Run the Application**
   ```bash
   go run cmd/main.go
   ```
   - This will:
     - Initialize the SQLite database (`jobboard.db`) in the `backend/` directory.
     - Apply the schema and seed data.
     - Print a list of available jobs to the console.

6. **Verify Output**
   - You should see:
     ```
     DB initialized w/ schema and seeded.
     Available Jobs:
     [List of jobs from seed.sql]
     ```
   - If there’s an error, check the logs for details (e.g., missing files or database issues).

### Notes
- The database file (`jobboard.db`) will be created in the `backend/` directory.
- To reset the database, delete `jobboard.db` and rerun the app.
- Future enhancements (e.g., web UI, scraping) will build on this foundation.



# 🛠️ Customizable Job Board – Project Requirements

## 🎯 Goal

Create a personalized job board for a small group of users (starting with a friend seeking Event Coordinator roles). The board will collect jobs from external sources via scraping, allows filtering, and links directly to job applications (bypassing LinkedIn, Indeed, etc when possible to allow them to apply on the company website).

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
  - Job Application URL (direct)
  - Source name
  - Posted date

### 3. Filtering UI (eventually)
- Filter by:
  - Keyword search
  - Location & If Remote/Hybrid/Onsite
  - Employment type
  - Tags/Industries

### 4. Scraper / Data Collector
- Scrapes jobs from:
  - Company sites
  - Smaller boards (e.g. WorkInCulture, ArtsWork, etc.)
- Stores jobs in DB or structured file
- (eventually) Can check if posting is still available

### 5. Personalized Feeds
- Users see a filtered feed based on their preferences (saved)
- Could be based on a URL param or config at first

---

## Nice-to-Haves (Post-MVP)
- Daily/weekly updates - depends on user
- Bookmark/save jobs
- Application tracking (e.g. “Applied”, “Interviewed”) - Kanban style
- Admin panel for scraping jobs manually or scheduling updates
- Mobile responsive or PWA
- Integration(Not sure about priority here) with Airtable, Notion, Google Sheets

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

- **Backend:** Go
- **Frontend (later):** Next.js + TypeScript
- **Database:** SQLite for MVP, Postgres later
- **Scraping:** Python (BeautifulSoup/Scrapy) or Go (Colly)
- **Deployment:** Render, Railway, or Fly.io - maybe AWS?

---

## 🔜 Working on

1. [] Define target job sites to scrape
2. [] Set up schema and DB
3. [] Write scraper for 1–2 sources
4. [] Expose a `/api/jobs` endpoint
5. [] Create simple frontend to display jobs