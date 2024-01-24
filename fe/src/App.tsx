import React, { useEffect, useState } from "react";
import "./App.css";
import {
  Configuration,
  DefaultApi,
  User,
  ResponseError,
  ApiErrorFromJSON,
  ApiError
} from "./client";
import { Alert, Button, Container, Form } from "react-bootstrap";

function App() {
  const [user, setUser] = useState<User>();
  const [client] = useState<DefaultApi>(
    new DefaultApi(new Configuration({ basePath: "/api" })),
  );
  const [code, setCode] = useState<string | undefined>(undefined);
  const [error, setError] = useState<ApiError | undefined>(undefined);
  const [message, setMessage] = useState<string | undefined>(undefined);

  useEffect(() => {
    client.getUser().then((user) => {
      setUser(user);
    });
  }, [client]);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();
    if (!isValid(code ?? "")) {
      // Should never happen with our validation
      alert("Invalid pin code");
      return;
    }

    client
      .updatePin({
        setCodePayload: { code: parseInt(code ?? "0") },
      })
      .then(() => {
        setError(undefined);
        setMessage("Pin code updated");
      })
      .catch((e: ResponseError) => {
        e.response.json().then((json) => {
          const apiError = ApiErrorFromJSON(json);
          setError(apiError);
        });
      });
  };

  const removeInvalid = (code: string): string => {
    return code.replace(/[^1-9]/g, "");
  };

  const isValid = (code: string): boolean => {
    if (/[0-9]{6}/.test(code)) {
      return true;
    }
    return false;
  };

  return (
    <Container className="d-flex flex-column min-vh-100 justify-content-center align-items-center">
      <Form className="my-3" onSubmit={handleSubmit}>
        <Form.Group className="mb-3">
          <Form.Label>Email Address</Form.Label>
          <Form.Control type="email" disabled defaultValue={user?.email} />
        </Form.Group>
        <Form.Group className="mb-3">
          <Form.Label>Name</Form.Label>
          <Form.Control type="text" disabled defaultValue={user?.name} />
        </Form.Group>
        <Form.Group className="mb-3">
          <Form.Label>Pin Code</Form.Label>
          <Form.Control
            maxLength={6}
            minLength={6}
            type="password"
            onChange={(e) => {
              if (code === undefined) {
                setCode("");
              } else {
                setCode(removeInvalid(e.target.value));
              }
            }}
            value={code ?? ""}
            isValid={isValid(code ?? "")}
            isInvalid={!isValid(code ?? "")}
          />
          <Form.Control.Feedback type="invalid">
            Must have exactly 6 digits
          </Form.Control.Feedback>
        </Form.Group>
        <Button className="mb-3" type="submit" disabled={!isValid(code?? "")}>Update Pin Code</Button>
        {message && <Alert variant="success">{message}</Alert>}
        {error && <Alert title={error.error} variant="danger">{error.displayMessage}</Alert>}
      </Form>
    </Container>
  );
}

export default App;
