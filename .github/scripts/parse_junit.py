#!/usr/bin/env python3
"""
parse_junit.py
Parses every *-junit.xml in --results-dir, writes a Markdown job summary
to --summary-file (GITHUB_STEP_SUMMARY), and dumps totals to --output-file.
"""

import argparse
import glob
import json
import os
import sys
import xml.etree.ElementTree as ET
from dataclasses import dataclass, field
from pathlib import Path
from typing import List


@dataclass
class TestCase:
    classname: str
    name: str
    time: float
    status: str          # PASS | FAIL | ERROR | SKIP
    message: str = ""
    stdout: str = ""


@dataclass
class Suite:
    name: str
    tests: int = 0
    failures: int = 0
    errors: int = 0
    skipped: int = 0
    time: float = 0.0
    cases: List[TestCase] = field(default_factory=list)


def parse_file(xml_path: str) -> Suite:
    tree = ET.parse(xml_path)
    root = tree.getroot()

    # Handle both <testsuites> wrappers and bare <testsuite> roots
    suite_elem = root if root.tag == "testsuite" else root.find("testsuite")
    if suite_elem is None:
        suite_elem = root  # fall back to root if no testsuite child

    suite = Suite(
        name=suite_elem.get("name") or Path(xml_path).stem,
        tests=int(suite_elem.get("tests", 0)),
        failures=int(suite_elem.get("failures", 0)),
        errors=int(suite_elem.get("errors", 0)),
        skipped=int(suite_elem.get("skipped", 0) or suite_elem.get("skip", 0)),
        time=float(suite_elem.get("time", 0) or 0),
    )

    for tc in suite_elem.findall(".//testcase"):
        failure = tc.find("failure")
        error   = tc.find("error")
        skipped = tc.find("skipped")
        sysout  = tc.find("system-out")

        if failure is not None:
            status  = "FAIL"
            # Prefer the full CDATA body over the short summary in the message attribute
            message = (failure.text or "").strip() or failure.get("message", "")
        elif error is not None:
            status  = "ERROR"
            message = (error.text or "").strip() or error.get("message", "")
        elif skipped is not None:
            status  = "SKIP"
            message = skipped.get("message", "")
        else:
            status  = "PASS"
            message = ""

        suite.cases.append(TestCase(
            classname=tc.get("classname", ""),
            name=tc.get("name", "?"),
            time=float(tc.get("time", 0) or 0),
            status=status,
            message=message.strip()[:2000],
            stdout=(sysout.text or "").strip()[:1000] if sysout is not None else "",
        ))

    return suite


def status_icon(s: str) -> str:
    return {"PASS": "✅", "FAIL": "❌", "ERROR": "💥", "SKIP": "⏭️"}.get(s, "❓")


def write_summary(suites: List[Suite], out_fh) -> dict:
    total_tests    = sum(s.tests    for s in suites)
    total_failures = sum(s.failures for s in suites)
    total_errors   = sum(s.errors   for s in suites)
    total_skipped  = sum(s.skipped  for s in suites)
    total_passed   = total_tests - total_failures - total_errors - total_skipped

    overall_ok = total_failures == 0 and total_errors == 0

    headline = "✅  All tests passed" if overall_ok else "❌  Tests failed"

    lines = [
        f"## {headline}",
        "",
        "| | Count |",
        "|---|---|",
        f"| ✅ Passed  | **{total_passed}** |",
        f"| ❌ Failed  | **{total_failures}** |",
        f"| 💥 Errors  | **{total_errors}** |",
        f"| ⏭️ Skipped | **{total_skipped}** |",
        f"| 📋 Total   | **{total_tests}** |",
        "",
    ]

    for suite in suites:
        suite_ok = suite.failures == 0 and suite.errors == 0
        icon     = "✅" if suite_ok else "❌"

        lines += [
            f"### {icon} {suite.name}",
            "",
            f"**{suite.tests}** tests · "
            f"**{suite.failures}** failed · "
            f"**{suite.errors}** errors · "
            f"**{suite.skipped}** skipped · "
            f"_{suite.time:.2f}s_",
            "",
        ]

        bad_cases = [c for c in suite.cases if c.status in ("FAIL", "ERROR")]
        if bad_cases:
            lines += [
                "<details open>",
                f"<summary>Show {len(bad_cases)} failing test(s)</summary>",
                "",
                "| Test | Status | Message |",
                "|---|---|---|",
            ]
            for tc in bad_cases:
                full_name = f"{tc.classname}.{tc.name}" if tc.classname else tc.name
                msg = tc.message.replace("\n", " ").replace("|", "\\|")
                lines.append(
                    f"| `{full_name}` | {status_icon(tc.status)} {tc.status} | {msg} |"
                )
                if tc.stdout:
                    stdout_escaped = tc.stdout.replace("\r", "")
                    lines += [
                        "",
                        f"<details><summary><code>{tc.name}</code> — captured output</summary>",
                        "",
                        "```",
                        stdout_escaped,
                        "```",
                        "",
                        "</details>",
                    ]
            lines += ["", "</details>", ""]

    out_fh.write("\n".join(lines))
    out_fh.flush()

    return {
        "total":    total_tests,
        "passed":   total_passed,
        "failures": total_failures,
        "errors":   total_errors,
        "skipped":  total_skipped,
        "ok":       overall_ok,
    }


def emit_annotations(suites: List[Suite]) -> None:
    """Print a ::error:: workflow command for every failing test case.

    These appear as inline annotations in the GitHub Actions UI and
    in the PR checks panel — even before the user opens the job summary.
    """
    for suite in suites:
        for tc in suite.cases:
            if tc.status not in ("FAIL", "ERROR"):
                continue
            full_name = f"{tc.classname}/{tc.name}" if tc.classname else tc.name
            body = (tc.message or tc.stdout or "(no message)")[:500]
            # Encode special characters required by the workflow command format
            encoded = (
                body
                .replace("%", "%25")
                .replace("\r", "")
                .replace("\n", "%0A")
            )
            print(f"::error title={full_name}::{encoded}")


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--results-dir",  required=True)
    parser.add_argument("--summary-file", required=True)
    parser.add_argument("--output-file",  required=True)
    args = parser.parse_args()

    xml_files = sorted(glob.glob(os.path.join(args.results_dir, "**", "*-junit.xml"),
                                  recursive=True))

    if not xml_files:
        print("::warning::No *-junit.xml files found in", args.results_dir)
        with open(args.summary_file, "w") as f:
            f.write("## ⚠️  No test results found\n\n"
                    "Jenkins returned no JUnit XML artifacts.\n")
        with open(args.output_file, "w") as f:
            json.dump({"total": 0, "passed": 0, "failures": 0,
                       "errors": 0, "skipped": 0, "ok": True}, f)
        sys.exit(0)

    suites: List[Suite] = []
    for path in xml_files:
        try:
            suites.append(parse_file(path))
            print(f"Parsed: {path}")
        except Exception as exc:
            print(f"::warning::Could not parse {path}: {exc}")

    with open(args.summary_file, "a") as summary_fh:
        totals = write_summary(suites, summary_fh)

    emit_annotations(suites)

    with open(args.output_file, "w") as f:
        json.dump(totals, f, indent=2)

    print(json.dumps(totals, indent=2))

    if not totals["ok"]:
        print(f"\n❌ {totals['failures']} failure(s), {totals['errors']} error(s)")
        # Don't sys.exit(1) here — the workflow's final step handles the exit code
        # so that artifact upload still runs.


if __name__ == "__main__":
    main()
