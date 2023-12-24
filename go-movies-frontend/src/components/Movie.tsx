import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Movie as IMovie } from "./Movies";

function Movie() {
  const [movie, setMovie] = useState<IMovie>({} as IMovie);
  // get the movie id from the url
  let { id } = useParams();
  useEffect(() => {
    let myMovie = {
      id: 1,
      title: "The Shawshank Redemption",
      release_date: "1986-03-07",
      runtime: 116,
      mpaa_rating: "R",
      description: "Arandom description",
    };
    setMovie(myMovie);
  }, [id]);
  return (
    <div>
      <h2>Movie: {movie.title}</h2>
      <small>
        <em>
          {movie.release_date}, {movie.runtime} minutes, Rated{" "}
          {movie.mpaa_rating}
        </em>
      </small>
      <hr />
      <p>{movie.description}</p>
    </div>
  );
}

export default Movie;
