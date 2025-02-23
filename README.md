# UNCLE CLI 💫👴

A CLI integrated with the API of the most diverse LLM models, made for developers to obtain quick answers and analyses via terminal.

## Technologies
- [Golang](https://go.dev/#)
- [Cobra](https://github.com/spf13/cobra)
- [Charm](https://github.com/charmbracelet)

## LLMs
- Gemini
- ChatGPT | Soon...
- Deepseek | Soon ...
- Qwen | Soon ...

## Build app
1. Clone repository:
```bash
    git clone https://github.com/AykoSousa/uncle-cli.git
```
2. Install dependencies:
```bash
    go mod tidy
```
3. Build the application:
```bash
   go build
```

## Features

### Config
1. When running the application, pass configuration parameters to choose the llm and model.
```bash
    uncle -l "llm name" -m "llm model" -k "API_KEY"
```

2. You can configure an `uncle-config.json` file in the same directory where the executable is to use a default model without having to always run the application passing parameters such as API-KEY, MODEL and LLM Name.
```json
{
    "llm": "gemini | chatgpt, ...",
    "model": "gemini-1.5-flash",
    "api-key": "YOUR_API_KEY"
}
```

3. You can also customize the default configuration by creating an `uncle-config.json` file in any directory and passing it as a parameter to the application. You can also use the file as a base, but choose a different model.
```bash
    uncle --config "path/to/config/file" -m "llm model"
```

### Execution
To run the application you must pass a prompt with a configuration file already created or manually pass the parameters as described above.
```bash
    uncle -m "gemini-2.0-flash" -p "What is DevOps culture?"
```

Output:
```text
DevOps culture is a set of practices, philosophies, and tools that aims to improve collaboration and communication between development and operations teams, ultimately accelerating the software development lifecycle and delivering higher-quality software faster. It's about breaking down silos and fostering a culture of shared responsibility, automation, and continuous improvement.
...
```

## Contributions
This project aims to improve my Golang skills using the API integrations that I frequently do in my work. Feel free to suggest improvements and contribute to the implementation of features that solve your specific problems.