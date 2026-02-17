---
name: cargo
description: "Rust development: build, test, clippy, publish"
---
# Cargo (Rust)

Use the **bash** tool for Rust development with Cargo.

## Setup

1. **Check if installed:**
   ```bash
   command -v cargo && cargo --version
   ```

2. **Install (rustup provides cargo):**
   ```bash
   # macOS
   brew install rustup-init && rustup-init

   # Official installer (cross-platform)
   curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
   source $HOME/.cargo/env
   ```

## Build & Run
```bash
cargo build
cargo build --release
cargo run
cargo run --release
```

## Test
```bash
cargo test
cargo test -- --nocapture
cargo test test_name
cargo test --doc
cargo bench
```

## Lint & Format
```bash
cargo clippy -- -W clippy::pedantic
cargo fmt
cargo fmt -- --check
```

## Dependencies
```bash
cargo add <crate>
cargo update
cargo tree
cargo audit
```

## Publish
```bash
cargo publish --dry-run
cargo publish
```

## Tips
- Use cargo clippy before committing
- Use cargo audit for security vulnerabilities
- Use cargo tree to visualize dependency graph
