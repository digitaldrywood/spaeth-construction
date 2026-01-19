# Cloverbelt Construction Website

## Project Overview

A marketing website for Cloverbelt Construction LLC, a Wisconsin-based construction company serving the Chippewa Valley region since 2001.

## Tech Stack

- **Backend**: Go 1.25 with Echo v4 web framework
- **Templating**: Templ for type-safe HTML templates
- **Interactivity**: HTMX for server-driven UI updates + Alpine.js for client-side state
- **Styling**: Tailwind CSS v4
- **Email**: Brevo API for contact form notifications
- **Deployment**: Vercel (serverless Go functions)

## Project Structure

```
.
├── api/                    # Vercel serverless function entry point
│   └── index.go
├── cmd/server/             # Main server application
│   ├── main.go
│   ├── slog.go            # Structured logging setup
│   └── generate.go        # go generate directives
├── internal/
│   ├── config/            # Environment configuration
│   ├── ctxkeys/           # Context keys for request-scoped values
│   ├── email/             # Brevo email integration
│   ├── handler/           # HTTP handlers and routes
│   ├── meta/              # Page metadata helpers
│   └── middleware/        # Echo middleware setup
├── templates/
│   ├── components/        # Reusable UI components
│   │   ├── about.templ
│   │   ├── contact.templ
│   │   ├── gallery.templ
│   │   ├── hero.templ
│   │   ├── projects.templ
│   │   └── services.templ
│   ├── layouts/           # Base layouts and common elements
│   │   ├── base.templ
│   │   ├── footer.templ
│   │   ├── header.templ
│   │   └── meta.templ
│   └── pages/             # Full page templates
│       ├── about.templ
│       ├── contact.templ
│       ├── home.templ
│       ├── projects.templ
│       └── services.templ
├── static/
│   └── css/
│       ├── input.css      # Tailwind source with custom styles
│       └── output.css     # Generated CSS (gitignored)
├── public/
│   └── images/            # Static images
└── vercel.json            # Vercel deployment config
```

## Development Commands

```bash
# Install development tools
make setup

# Run with hot reload (Air)
make dev

# Build production binary
make build

# Run tests
make test

# Lint code
make lint

# Generate templ files
make generate

# Build Tailwind CSS
make css

# Watch Tailwind CSS
make css-watch
```

## Environment Variables

Copy `.envrc.example` to `.envrc` and configure:

```bash
PORT=3000
ENV=development                    # or "production"
SITE_NAME="Cloverbelt Construction"
SITE_URL="http://localhost:3000"   # Production URL for meta tags
BREVO_API_KEY=""                   # Required for contact form emails
CONTACT_EMAIL="info@cloverbeltconstructionwi.com"
LOG_LEVEL=INFO                     # DEBUG, INFO, WARN, ERROR
```

## Routes

| Path       | Description           |
|------------|-----------------------|
| `/`        | Home page             |
| `/about`   | About us page         |
| `/services`| Services overview     |
| `/projects`| Project portfolio     |
| `/contact` | Contact form          |
| `/health`  | Health check endpoint |

## Key Patterns

### Page Metadata

Use the `meta.New()` builder for SEO metadata:

```go
meta.New("Page Title", "Description text").
    WithOGImage("/images/og-image.jpg").
    WithKeywords("keyword1", "keyword2")
```

### HTMX Contact Form

The contact form uses HTMX for seamless submissions:

```html
<form hx-post="/contact" hx-swap="outerHTML">
```

### Alpine.js Components

Client-side state is managed with Alpine.js:

- Header scroll state: `x-data="{ scrolled: false }"`
- Mobile menu toggle: `x-data="{ mobileOpen: false }"`
- Hero carousel: `x-data="heroCarousel()"`

### Tailwind Custom Utilities

Custom CSS classes defined in `static/css/input.css`:

- `.blueprint-grid` - Blueprint-style grid background
- `.construction-grid` - Construction grid pattern
- `.construction-stripes` - Diagonal stripe pattern
- `.btn-primary` - Primary green gradient button
- `.btn-secondary` - Secondary outline button
- `.card-construction` - Card with hover effects
- `.input-construction` - Styled form inputs

## Deployment

### Vercel

The project is configured for Vercel serverless deployment:

1. Connect repository to Vercel
2. Set environment variables in Vercel dashboard
3. Deploy automatically on push to main

Build commands are defined in `vercel.json`.

### Manual Build

```bash
# Build everything
make build

# Run production binary
./spaeth-construction
```

## Brand Colors

Primary green palette (Cloverbelt brand):

- `#16a34a` - Primary green (600)
- `#22c55e` - Light green (500)
- `#15803d` - Dark green (700)
- `#f0fdf4` - Lightest green (50)

## Adding New Pages

1. Create handler method in `internal/handler/pages.go`
2. Create template in `templates/pages/`
3. Register route in `internal/handler/handler.go`
4. Update navigation in `templates/layouts/header.templ`

## Contact Form Flow

1. User submits form on `/contact`
2. Handler validates required fields
3. If Brevo API key is set, sends email notification
4. Returns success state via HTMX swap
5. URL updated to `/contact?success=true`
