import React from "react";
import Input from "./form/Input";
import { useNavigate, useOutletContext } from "react-router-dom";
interface ContextType {
  setJwtToken: (token: string) => void;
  setAlertClassName: (className: string) => void;
  setAlertMessage: (message: string) => void;
  // include other properties of the context here
}
function Login() {
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");

  const navigate = useNavigate();

  const { setJwtToken, setAlertClassName, setAlertMessage } =
    useOutletContext<ContextType>();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (email === "admin@example.com") {
      console.log("Login successful");
      setJwtToken("token");
      setAlertClassName("d-none");
      setAlertMessage("");
      navigate("/");
    } else {
      setAlertClassName("alert-danger");
      setAlertMessage("Invalid credentials");
    }
  };
  return (
    <div className="col-md-6 offset-md-3">
      <h2>Login</h2>
      <hr />
      <form onSubmit={handleSubmit}>
        <Input
          name="email"
          title="Email Address"
          type="email"
          value={email}
          autoComplete="email-new"
          handleChange={(e) => setEmail(e.currentTarget.value)}
        />
        <Input
          name="password"
          title="Password"
          type="text"
          value={password}
          autoComplete="password-new"
          handleChange={(e) => setPassword(e.currentTarget.value)}
        />
        <input type="submit" className="btn btn-primary" value="Login" />
      </form>
    </div>
  );
}

export default Login;
