import { useCallback, useEffect, useState } from "react";
import { Link, Outlet, useNavigate } from "react-router-dom";
import Alert from "./components/Alert";

function App() {
  const [jwtToken, setJwtToken] = useState<string>("");
  const [alertMessage, setAlertMessage] = useState<string>("");
  const [alertClassName, setAlertClassName] = useState<string>("d-none");
  const [tickInterval, setTickInterval] = useState<NodeJS.Timeout>();

  const navigate = useNavigate();



  const toggleRefresh = useCallback((status: boolean) => {
    console.log("toggling refresh");

    if (status) {

      let i = setInterval(() => {
        const requestOptions: RequestInit = {
          method: "GET",
          headers: { "Content-Type": "application/json" },
          credentials: "include" as RequestCredentials,
        };
        console.log("turning on ticking");
        fetch("/refresh", requestOptions)
          .then((response) => response.json())
          .then((data) => {
            if (data.access_token) {
              setJwtToken(data.access_token);

            }
          })
          .catch((error) => console.log("user not logged in"));
        console.log("ticking");
      }, 600000);
      setTickInterval(i);

      console.log("setting tick interval to ", i);
    } else {
      console.log("turning off ticking", tickInterval);
      setTickInterval(undefined);
      clearInterval(tickInterval as NodeJS.Timeout);
    }
  }, [tickInterval]);

  const logOut = () => {
    const requestOptions: RequestInit = {
      method: "GET",
      headers: { "Content-Type": "application/json" },
      credentials: "include" as RequestCredentials,
    };
    fetch("/logout", requestOptions)
      .catch((error) => console.log("error logging out", error))
      .finally(() => setJwtToken(""));
      toggleRefresh(false);

    navigate("/login");
  };

  useEffect(() => {
    if (jwtToken === "") {
      const requestOptions: RequestInit = {
        method: "GET",
        headers: { "Content-Type": "application/json" },
        credentials: "include" as RequestCredentials,
      };

      fetch("/refresh", requestOptions)
        .then((response) => response.json())
        .then((data) => {
          if (data.access_token) {
            setJwtToken(data.access_token);
            toggleRefresh(true);
          }
        })
        .catch((error) => console.log("user not logged in", error));
    }
  }, [jwtToken, toggleRefresh]);

  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <h1 className="mt-3">Go watch a movie</h1>
        </div>
        <div className="col text-end">
          {jwtToken === "" ? (
            <Link to="/login">
              <span className="badge bg-success">Login</span>
            </Link>
          ) : (
            <a href="#!" onClick={logOut}>
              <span className="badge bg-danger">Logout</span>
            </a>
          )}
        </div>
        <hr className="mb-3" />
      </div>
      <div className="row">
        <div className="col-md-2">
          <nav>
            <div className="list-group">
              <Link to="/" className="list-group-item list-group-item-action">
                Home
              </Link>
              <Link
                to="/movies"
                className="list-group-item list-group-item-action"
              >
                Movies
              </Link>
              <Link
                to="/genres"
                className="list-group-item list-group-item-action"
              >
                Genres
              </Link>
              {jwtToken !== "" && (
                <>
                  <Link
                    to="/admin/movie/0"
                    className="list-group-item list-group-item-action"
                  >
                    Add Movie
                  </Link>
                  <Link
                    to="/manage-catalogue"
                    className="list-group-item list-group-item-action"
                  >
                    Manage Catalogue
                  </Link>
                  <Link
                    to="graphql"
                    className="list-group-item list-group-item-action"
                  >
                    GraphQl
                  </Link>
                </>
              )}
            </div>
          </nav>
        </div>
        <div className="col-md-10">

          <Alert message={alertMessage} className={alertClassName} />
          <Outlet
            context={{
              jwtToken: jwtToken,
              setJwtToken: setJwtToken,
              setAlertClassName,
              setAlertMessage,
              toggleRefresh,
            }}
          />
        </div>
      </div>
    </div>
  );
}

export default App;
