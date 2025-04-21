# Word Reverser

> A modern desktop app for reversing the letters in each word while maintaining word order.

## Features

- Reverses each word in a text string letter by letter while keeping word order intact
- Clean, modern UI with responsive design
- Real-time text transformation
- Beautiful gradient styling with smooth animations

## How to Use

1. Enter your text in the input field
2. Click the "Reverse Words" button
3. See the reversed text in the result box below

For example:
- Input: `Hello world`
- Output: `olleH dlrow`

## Tech Stack

This app is built with:
- [Wails](https://wails.io/) - Go framework for desktop apps
- [Svelte](https://svelte.dev/) + TypeScript - Frontend UI
- Modern CSS - Styling with gradients and animations

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Antivirus Detection

Some antivirus programs, including Windows Defender, may flag this application as suspicious. This is a false positive that occurs with many Wails applications. If you encounter this:

1. You can submit a false positive report to Microsoft: https://www.microsoft.com/en-us/wdsi/filesubmission
2. Temporarily disable your antivirus to install the application
3. Add an exception for this application in your antivirus software