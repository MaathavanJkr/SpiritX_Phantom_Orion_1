# SecureConnect
## Usage
To start using SecureConnect, follow these steps:

### Backend
1. **Set up environment variables:**
    Create a `.env` file in the root directory and add the necessary environment variables. Refer to `.env.example` for the required variables.

2. **Install Go:**
    Ensure Go is installed on your system. If not, download and install it from [golang.org](https://golang.org/).

3. **Install the required dependencies:**
    ```bash
    go mod download
    ```

4. **Run the backend application:**
    ```bash
    go run .
    ```
5. **Access the API:**
    The API will be running on `http://localhost:8080`.

### Frontend User


1. **Navigate to the fron4. **Access the application:**
    Open your web browser and navigate to `http://localhost:8080`.tend directory:**
    ```bash
    cd frontend
    ```
2. **Set up environment variables:**
    Create a `.env` file in the root directory and add the necessary environment variables. 

    ```plaintext
    NEXT_PUBLIC_BACKEND_URL=http://localhost:8080
    ```

3. **Install React:**
    ```bash
    npm install
    ```

4. **Run the frontend application:**
    ```bash
    npm run dev
    ```

5. **Access the application:**
    Open your web browser and navigate to `http://localhost:3000` to access the user interface.
