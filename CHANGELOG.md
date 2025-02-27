# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Changed
* Updated license to reflect current year (2025)

### Fixed
* Fix writer failure not being returned from `Patch` methods.

## [1.0.2] - 2025-02-27
### Changed
* Update kingpin v2 import path.

### Security
* bsdiff: Fix CVE-2014-9862 (https://www.x41-dsec.de/lab/advisories/x41-2020-006-bspatch/)
* bsdiff: Fix compile error because of missing limits.h

## [1.0.1] - 2022-07-11
### Changed
* Code is now properly licensed under 2-clause BSD license. #7
* Updated github.com/dsnet/compress to 0.0.1.

## [1.0.0] - 2018-11-30
### Added
* Add bsdiff diffing and patching functionality.

[Unreleased]: https://github.com/icedream/go-bsdiff/compare/v1.0.2...HEAD
[1.0.2]: https://github.com/icedream/go-bsdiff/releases/tag/v1.0.2
[1.0.1]: https://github.com/icedream/go-bsdiff/releases/tag/v1.0.1
[1.0.0]: https://github.com/icedream/go-bsdiff/releases/tag/v1.0.0



[_breaking_change_token]: BREAKING
