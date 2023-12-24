import React from "react";
import Ticket from "../images/movie-tickets.jpg";
import { Link } from "react-router-dom";

function Home() {
  return (
    <div className="text-center">
      <h2>Find a movie to watch</h2>
      <hr />
      <Link to="/movies">
        {" "}
        <img src={Ticket} alt="movie tickets" />{" "}
      </Link>
    </div>
  );
}

export default Home;
