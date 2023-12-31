# Changelog

## [0.0.6](https://github.com/jimeh/macos-battery-exporter/compare/v0.0.5...v0.0.6) (2023-12-21)


### Bug Fixes

* **output:** write log output to stderr by default instead of stdout ([7b2e98e](https://github.com/jimeh/macos-battery-exporter/commit/7b2e98e150f9ee6d630023879395df0912047667))

## [0.0.5](https://github.com/jimeh/macos-battery-exporter/compare/v0.0.4...v0.0.5) (2023-12-16)


### Features

* **go:** rename prom package to suitable prombat ([436e4a4](https://github.com/jimeh/macos-battery-exporter/commit/436e4a4b01d96654b7012f795f2d305ca4084681))

## [0.0.4](https://github.com/jimeh/macos-battery-exporter/compare/v0.0.3...v0.0.4) (2023-12-16)


### Continuous Integration

* **lint:** setup golangci-lint ([2ac3ecb](https://github.com/jimeh/macos-battery-exporter/commit/2ac3ecb555e0f6eea369516328f1f03da7d61251))

## [0.0.3](https://github.com/jimeh/macos-battery-exporter/compare/v0.0.2...v0.0.3) (2023-12-16)


### Bug Fixes

* **package:** resolve issue with running as a homebrew service ([9a168f9](https://github.com/jimeh/macos-battery-exporter/commit/9a168f9ff918f6539ca85d43202759197ed952b3))

## [0.0.2](https://github.com/jimeh/macos-battery-exporter/compare/v0.0.1...v0.0.2) (2023-12-16)


### Bug Fixes

* **battery:** find ioreg executable more reliably ([993b036](https://github.com/jimeh/macos-battery-exporter/commit/993b036d99362b6bebd36545fc34d325863421d5))

## 0.0.1 (2023-12-16)


### Features

* **init:** initial working implementation ([204938f](https://github.com/jimeh/macos-battery-exporter/commit/204938f5b18712e5314cb47c96ee1fbc04fbe70d))


### Bug Fixes

* **battery:** handle no batteries being present ([a3dc259](https://github.com/jimeh/macos-battery-exporter/commit/a3dc259e3b57dd386fd05ce6b57ad14d7940238d))
* **makefile:** correctly set VERSION ([1d464fb](https://github.com/jimeh/macos-battery-exporter/commit/1d464fbd3af42a55b04b7b1468ffef2711f8d7cf))
* **metrics:** adhere to units in metric names ([471ba43](https://github.com/jimeh/macos-battery-exporter/commit/471ba437c43db4eefe228fa6a007833a40b55af1))
