# a<sup>n</sup>b<sup>n</sup>c<sup>n</sup> Validator

> An app for validating if a string conforms to the language L = {a<sup>n</sup>b<sup>n</sup>c<sup>n</sup> / n ≥ 0}.

## Features

- Validates input strings against the a<sup>n</sup>b<sup>n</sup>c<sup>n</sup> pattern.
- Supports the empty string case (where n=0).
- Provides immediate "YES" or "NO" feedback.
- Clean, modern UI with responsive design.

## How to Use

1. Enter your string in the input field (e.g., `aaabbbccc`, `abc`, or leave empty for n=0).
2. Click the "Validate" button.
3. See the result ("YES" or "NO") in the result box below.

For example:
- Input: `aaabbbccc`
- Output: `YES`

- Input: `aabbc`
- Output: `NO`

- Input: `` (empty string)
- Output: `YES`

## Tech Stack

This app is built with:
- [Wails](https://wails.io/) - Go framework for desktop apps (provides the application shell).
- [Svelte](https://svelte.dev/) + TypeScript - Frontend UI and core validation logic.
- Modern CSS - Styling.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes.

## Building

To build a redistributable, production mode package, use `wails build`.

## Antivirus Detection

Some antivirus programs, including Windows Defender, may flag this application as suspicious. This is a false positive that occurs with many Wails applications. If you encounter this:

1. You can submit a false positive report to Microsoft: https://www.microsoft.com/en-us/wdsi/filesubmission
2. Temporarily disable your antivirus to install the application
3. Add an exception for this application in your antivirus software
