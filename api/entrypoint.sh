#!/bin/sh

rustup override set nightly
cargo install diesel_cli --no-default-features --features postgres
diesel migration run

exec "$@"
