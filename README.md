# Stockyard Grubstake

**Personal finance tracker — import bank CSV, categorize transactions, visualize spending**

Part of the [Stockyard](https://stockyard.dev) family of self-hosted developer tools.

## Quick Start

```bash
docker run -p 9320:9320 -v grubstake_data:/data ghcr.io/stockyard-dev/stockyard-grubstake
```

Or with docker-compose:

```bash
docker-compose up -d
```

Open `http://localhost:9320` in your browser.

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `9320` | HTTP port |
| `DATA_DIR` | `./data` | SQLite database directory |
| `GRUBSTAKE_LICENSE_KEY` | *(empty)* | Pro license key |

## Free vs Pro

| | Free | Pro |
|-|------|-----|
| Limits | 2 accounts, 500 transactions | Unlimited accounts and transactions |
| Price | Free | $2.99/mo |

Get a Pro license at [stockyard.dev/tools/](https://stockyard.dev/tools/).

## Category

Creator & Small Business

## License

Apache 2.0
