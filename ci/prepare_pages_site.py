#!/usr/bin/env python3
# SPDX-license-identifier: Apache-2.0
##############################################################################
# Copyright (c) 2026
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Apache License, Version 2.0
# which accompanies this distribution, and is available at
# http://www.apache.org/licenses/LICENSE-2.0
##############################################################################

"""Stage documentation for GitHub Pages.

This keeps the source Markdown compatible with Runme while removing
Runme-specific code fence attributes from the staged copy consumed by Jekyll.
"""

from __future__ import annotations

import argparse
import re
import shutil
from pathlib import Path

runmeFencePattern = re.compile(
    r"^(?P<indent>\s*)(?P<fence>`{3,}|~{3,})(?P<language>[^\s`~{]+)\s+\{[^}]+\}(?P<eol>\r?\n?)$"
)


def strip_runme_fence_attributes(content: str) -> str:
    """Remove Runme cell attributes from fenced code block opening lines."""

    return "".join(
        runmeFencePattern.sub(
            r"\g<indent>\g<fence>\g<language>\g<eol>",
            line,
        )
        for line in content.splitlines(keepends=True)
    )


def stage_docs(source_dir: Path, destination_dir: Path) -> None:
    """Copy docs to a staging directory and sanitize Markdown for Pages."""

    if destination_dir.exists():
        shutil.rmtree(destination_dir)

    shutil.copytree(source_dir, destination_dir)

    for markdown_file in destination_dir.rglob("*.md"):
        sanitized = strip_runme_fence_attributes(markdown_file.read_text())
        markdown_file.write_text(sanitized)


def main() -> None:
    """Parse arguments and stage the documentation tree."""

    parser = argparse.ArgumentParser(
        description="Copy docs into a Pages staging directory and strip Runme fence attributes."
    )
    parser.add_argument("source", type=Path, help="source documentation directory")
    parser.add_argument("destination", type=Path, help="destination staging directory")
    args = parser.parse_args()

    stage_docs(args.source, args.destination)


if __name__ == "__main__":
    main()
