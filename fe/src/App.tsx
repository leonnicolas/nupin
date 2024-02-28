import React, { useEffect, useState } from "react";
import "./App.css";
import {
  Configuration,
  DefaultApi,
  User,
  ResponseError,
  ApiErrorFromJSON,
  ApiError,
} from "./client";
import {
  Alert,
  Button,
  Container,
  Form,
  Navbar,
  Nav,
  Spinner,
} from "react-bootstrap";

const client = new DefaultApi(new Configuration({ basePath: "/api" }));

function App() {
  const [user, setUser] = useState<User>();
  const [code, setCode] = useState<string | undefined>(undefined);
  const [error, setError] = useState<ApiError | undefined>(undefined);
  const [message, setMessage] = useState<string | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    client.getUser().then((user) => {
      setUser(user);
      setCode("");
    });
  }, []);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();
    if (!isValid(code ?? "")) {
      // Should never happen with our validation
      alert("Invalid pin code");
      return;
    }
    setLoading(true);
    setError(undefined);
    setMessage(undefined);

    client
      .updatePin({
        setCodePayload: { code: parseInt(code ?? "0") },
      })
      .then(() => {
        setError(undefined);
        setMessage("Pin code updated");
      })
      .finally(() => {
        setLoading(false);
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
    <div>
      <Navbar>
        <Container>
          <Navbar.Brand href="/">Door Pin</Navbar.Brand>
          <Nav>
            <Button href="/logout">Logout</Button>
          </Nav>
        </Container>
      </Navbar>
      <Container className="d-flex flex-column justify-content-center align-items-center my-5">
        <Form onSubmit={handleSubmit}>
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
                setCode(removeInvalid(e.target.value));
              }}
              value={code ?? ""}
              isValid={isValid(code ?? "")}
              isInvalid={!isValid(code ?? "")}
            />
            <Form.Control.Feedback type="invalid">
              Must have exactly 6 digits
            </Form.Control.Feedback>
          </Form.Group>
          <Button
            className="mb-3"
            type="submit"
            disabled={!isValid(code ?? "") || loading}
          >
            <span>Update Pin Code</span>
            {loading && <Spinner className="ms-2" size="sm" />}
          </Button>
          {message && <Alert variant="success">{message}</Alert>}
          {error && (
            <Alert title={error.error} variant="danger">
              {error.displayMessage}
            </Alert>
          )}
        </Form>
      </Container>
    </div>
  );
}

export default App;
