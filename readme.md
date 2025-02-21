
# PassGen API

The PassGen API is a simple API to generate random passwords based on user-provided parameters. In the future I plan to add more complex patterns for generating passwords, such as passwords based on keywords.

## Endpoints

### `POST /randomPass`

Generates a random password based on the provided parameters.

#### Request

- **Method**: `POST`
- **URL**: `/randomPass`
- **Headers**: 
  - `Content-Type: application/json`
- **Body**:
  ```json
  {
    "size": 12,
    "useUpper": true,
    "useNum": true,
    "useSpecial": true
  }
  ```

  - `size` (int): Desired password length. Default value is `8`.
  - `useUpper` (bool): Include uppercase letters. Default value is `true`.
  - `useNum` (bool): Include numbers. Default value is `true`.
  - `useSpecial` (bool): Include special characters. Default value is `true`.

#### Response

- **Status Code**: `200 OK`
- **Headers**: 
  - `Content-Type: application/json`
- **Body**:
  ```json
  {
    "password": "aB3$dEfGhIjK"
  }
  ```
    -  `Password` max size is capped to 128

#### Errors

- **Status Code**: `405 Method Not Allowed`
  - **Description**: HTTP method not allowed. Only `POST` is accepted.
- **Status Code**: `400 Bad Request`
  - **Description**: Password size too short. Minimum size is `4`.
- **Status Code**: `500 Internal Server Error`
  - **Description**: Error encoding the response.

## How to Start the Server

To start the server, run the following command:

```sh
go run cmd/main.go
```

The server will be available at `http://localhost:8000`.


## Example Usage

To generate a random password, send a `POST` request to `http://localhost:8000/randomPass` with the following body:

```json
{
  "size": 12,
  "useUpper": true,
  "useNum": true,
  "useSpecial": true
}
```

The response will be a randomly generated password based on the provided parameters.
