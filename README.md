# 🚀 Ulsidor: Technical Overview 🌌

## 📌 Introduction
Ulsidor, developed by Blackwater Industries, represents the next generation in TTRPG management. This toolset, leveraging a patented entity management system, facilitates efficient oversight of quests, NPCs, locations, and factions. The aim is to provide a comprehensive solution without compromising on quality.

---

## 🌐 Blackwater Industries System Architecture

### 🖥️ Core (`internal/core/core.go`)

The Core acts as the central processing unit of the Ulsidor toolset. It's designed to ensure seamless orchestration between all application sub-systems. The upcoming "Modules" are set to further enhance the capabilities of the toolset.

#### 🧪 Codex-o-Matic™

The Codices are specialized "Entity Handlers". Each entity type is paired with its dedicated codex to ensure adaptability and precision in management.

#### 💾 Repository Vaults

Repository Vaults are integral to data management. They're designed to safeguard entities, ensuring data consistency and durability. The "lazy" loading mechanism optimizes resource utilization.

---

## 🖱️ CLI (`cmd/ulsidor/`)

The CLI serves as the primary interface for Ulsidor. It's designed for intuitive command input and interactive prompts, facilitating efficient game management.

### 🛸 Survey-o-Tron™

The Survey component is an essential part of the CLI. It's responsible for collecting and validating entity data. Each entity type has been provided with a dedicated sub-package to ensure precision.

---

## 💽 Data Dynamics™

The Data Dynamics™ suite, a product of Blackwater Industries, is dedicated to external data and configuration management. It comprises:

- **config/**: 🛠️ Designed for application-specific configuration.
- **game/**: 🌌 Manages global game or world variables.
- **storage/**: 📦 Handles the loading and saving of entity repositories, with an integrated changelog mechanism for tracking modifications.

---

## 🛣️ Development Roadmap

### 🎯 Immediate Objectives:
- Refine CRUD commands for quests.
- Implement comprehensive unit testing.
- Enhance quest command functionalities.

### 🌠 Future Vision:
- Develop a state-of-the-art revision handler.
- Launch the Faction Tracker module.

---

📖 This document serves as a technical overview of Ulsidor, developed by Blackwater Industries. The focus remains on delivering a top-tier game management solution.
