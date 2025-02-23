# Attendance Tracker

A web-based application that helps students track their attendance and determine if they can afford to skip a class while maintaining the required attendance percentage.

## Features

- **Real-time Attendance Calculation**: Automatically updates attendance percentage based on inputs.
- **Customizable Timetable**: Users can set their daily class schedule.
- **Class Skipping Predictor**: Calculates whether a student can skip a class without falling below 75%.
- **Persistent Data Storage**: User data is stored in a database for future reference.
- **User Authentication**: Secure login to track individual progress.

## Tech Stack

### Backend

- **Golang** (with repository pattern & dependency injection)
- **Postgres SQL** (for local data storage)

### Frontend

- **React** (for an interactive UI)
- **TailwindCSS** (for styling)

## Installation

### Prerequisites

Ensure you have the following installed:

- Go
- Node.js & npm
- Postgres SQL

### Steps

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/attendance-tracker.git
   cd attendance-tracker
   ```

2. **Backend Setup:**

   ```sh
   cd backend
   go mod tidy
   go run main.go
   ```

3. **Frontend Setup:**
   ```sh
   cd frontend
   npm install
   npm run dev
   ```

## Usage

1. Log in or create an account.
2. Enter your class timetable and total working days.
3. Mark attendance for each class attended.
4. Check real-time attendance stats and skip class predictions.

## License

This project is licensed under the MIT License.
