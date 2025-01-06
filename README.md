# random-name-picker
Build a simple CLI tool to input names, randomly select one, and display the result.

## Core Features
Input: Accept a list of names via command-line input or a text file.

Output: Randomly select and display a name.

## Tasks
1. Write a Go program to:
    - Accept names as command-line arguments.
    - Print the list of names back to the user.

2. Randomly pick a name from the list:
    - Import the math/rand package for randomization.
    - Implement random selection logic.

3. Enhance the User Experience
    - Add prompts for better clarity:
        - Display how many names were entered.
    - Handle errors (e.g., no names provided).

4. Read names from a file
    - Allow names to be read from a file.
        - Use the os and bufio packages.
    - Save the randomly selected name to a results file.

5. Ensure the program works correctly in different scenarios.
    - Test edge cases:
        - No input.
        - Single name.
        - Long lists of names.
    - Debug any issues.