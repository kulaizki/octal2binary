<script lang="ts">
  let inputText: string = "";
  let resultText: string = "Enter text above and click 'Reverse Words' to see the result";

  function reverseWords(): void {
    if (!inputText) {
      resultText = "Please enter some text first";
      return;
    }

    const lines = inputText.split('\n');
    
    const processedLines = lines.map(line => {
      const words = line.split(' ');
      const reversedWords = words.map(word => {
        return word.split('').reverse().join('');
      });
      return reversedWords.join(' ');
    });

    resultText = processedLines.join('\n');
  }
</script>

<main>
  <div class="container">
    <div class="app-header">
      <h1>Word Reverser</h1>
      <p class="subtitle">Reverse the letters in each word while maintaining word order</p>
    </div>
    
    <div class="card">
      <div class="input-section">
        <label for="input-text">Input Text</label>
        <textarea 
          id="input-text"
          bind:value={inputText} 
          placeholder="Type your text here..."
          rows="4"
        ></textarea>
      </div>
      
      <button class="hover:cursor-pointer reverse-btn" on:click={reverseWords}>
        <span>Reverse Words</span>
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 3L21 7L17 11"></path><path d="M21 7H13C10.2386 7 8 9.23858 8 12V12C8 14.7614 5.76142 17 3 17H3"></path></svg>
      </button>
      
      <div class="result-section">
        <label for="result-text">Result</label>
        <div id="result-text" class="result-box">{resultText}</div>
      </div>
    </div>
    
    <div class="attribution">
      <p>Built with Svelte + Wails by <span class="name-gradient">Fitzsixto</span></p>
    </div>
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, sans-serif;
    background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
    background-repeat: no-repeat;
    background-attachment: fixed;
    min-height: 100vh;
    color: #f1f1f1;
    display: flex;
    justify-content: center;
    align-items: center;
    overflow: auto;
  }
  
  main {
    width: 100%;
    max-width: 100vw;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1rem;
    box-sizing: border-box;
  }
  
  .container {
    width: 90%;
    max-width: 600px;
    padding: 2rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    box-sizing: border-box;
  }

  .app-header {
    text-align: center;
    margin-bottom: 2rem;
    width: 100%;
  }

  h1 {
    font-size: 2.5rem;
    margin: 0;
    background: linear-gradient(to right, #4facfe 0%, #00f2fe 100%);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    font-weight: 700;
  }

  .subtitle {
    color: #b8c1ec;
    margin-top: 0.5rem;
    font-size: 1rem;
  }

  .card {
    width: 100%;
    background: rgba(255, 255, 255, 0.05);
    backdrop-filter: blur(10px);
    border-radius: 16px;
    padding: 1.5rem;
    box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
    border: 1px solid rgba(255, 255, 255, 0.1);
  }

  label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
    color: #4facfe;
  }

  .input-section, .result-section {
    margin-bottom: 1.5rem;
    width: 100%;
  }

  textarea {
    width: 100%;
    padding: 1rem;
    border-radius: 8px;
    border: 1px solid rgba(255, 255, 255, 0.2);
    font-size: 1rem;
    font-family: inherit;
    resize: vertical;
    background-color: rgba(0, 0, 0, 0.2);
    color: #fff;
    transition: all 0.3s ease;
    box-sizing: border-box;
    max-height: 200px;
  }

  textarea:focus {
    outline: none;
    border-color: #4facfe;
    box-shadow: 0 0 0 2px rgba(79, 172, 254, 0.3);
  }

  textarea::placeholder {
    color: rgba(255, 255, 255, 0.5);
  }

  .reverse-btn {
    background: linear-gradient(to right, #4facfe 0%, #00f2fe 100%);
    color: #1a1a2e;
    border: none;
    border-radius: 8px;
    padding: 0.75rem 1.5rem;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    margin: 1.5rem 0;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    width: auto;
    align-self: center;
  }

  .reverse-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 7px 14px rgba(0, 0, 0, 0.2);
  }

  .reverse-btn:active {
    transform: translateY(0);
  }

  .result-box {
    width: 100%;
    min-height: 80px;
    max-height: 200px;
    padding: 1rem;
    border-radius: 8px;
    background-color: rgba(0, 0, 0, 0.2);
    border: 1px solid rgba(255, 255, 255, 0.2);
    color: #fff;
    white-space: pre-wrap;
    word-break: break-word;
    box-sizing: border-box;
    overflow-y: auto;
  }

  .attribution {
    margin-top: 2rem;
    font-size: 0.85rem;
    color: rgba(255, 255, 255, 0.5);
    text-align: center;
  }

  .name-gradient {
    background: linear-gradient(to bottom, #87CEEB 0%, #1E90FF 100%);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    font-weight: 600;
  }

  @media (max-width: 768px) {
    .container {
      width: 95%;
      padding: 1rem;
    }
    
    .card {
      padding: 1rem;
    }
    
    h1 {
      font-size: 2rem;
    }
  }
</style>
