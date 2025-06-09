# Golang Terminal Simple AI chat
This Go program creates a terminal-based AI chat application using the Google Generative AI API and the Pterm library for interactive terminal UI. Here's a simple description:
The program loads a Gemini API key from a .env file and initializes a chat client. It enters a loop where it:

- Prompts the user for input using Pterm's interactive text input (displayed in magenta).
- Sends the input to the Gemini AI model (gemini-1.5-flash) to generate a response.
- Displays a spinner for 2 seconds to simulate processing.
- Prints the AI's response in light blue using Pterm.
- Stores the response in a local chat history struct, assigning it an ID.
- Continues the loop for ongoing chat interaction.

  The program uses godotenv to manage environment variables, pterm for styled terminal output, and the genai package to interact with Google's generative AI. The chat history is stored in memory but not persisted between sessions.

  ![image](https://github.com/user-attachments/assets/49b1bf76-f069-4ac3-aa8e-407c189d35a9)
