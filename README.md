# 🌌 NASA Space Project

A simple, fast space-tracking API built with Go and PostgreSQL. It automatically fetches live asteroid data from NASA, stores it permanently on your device, and lets you add or delete items in real time.

## ✨ Features
* **Live NASA Data Sync:** Reaches across the internet to parse, normalize, and inject real-time cosmic payload data structures.
* **Full CRUD Capabilities:** Implements dynamic routing for real-time creation (`POST`) and precise parameter-validated record removals (`DELETE`).
* **API Pagination Matrix:** Utilizes smart `limit` and `offset` algorithms to stream large astronomical data sets smoothly over the network.
* **Persistent Disk Storage:** Uses full relational database schemas to ensure custom space tracking data survives complete system restarts.

## 🛠️ Tech Stack
* **Backend:** Go (Golang)
* **Routing Engine:** Go Standard Library Multiplexer (`net/http`)
* **Database:** PostgreSQL (Relational RDBMS storage)
* **DB Driver:** `pgx/v5` (High-performance connection pooling)
