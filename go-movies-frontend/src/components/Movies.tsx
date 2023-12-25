import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";

export interface Movie {
    id: number;
    title: string;
    release_date: string;
    runtime: number;
    mpaa_rating: string;
    description: string;
}

function Movies() {
    const [movies, setMovies] = useState<Movie[]>([]);

    useEffect(() => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("Accept", "application/json");
       const requestOptions = {
            method: "GET",
            headers: headers,
        };
        fetch("http://localhost:8080/movies", requestOptions)
            .then((response) => response.json())
            .then((data) => setMovies(data))
            .catch((error) => console.log(error));
    }, []);

    return (
        <div>
            <h2>Movies</h2>
            <hr />
        <table className="table table-striped table-hover">
            <thead>
                <tr>
                    <th>Movie</th>
                    <th>Release Date</th>
                    <th>Rating</th>
                </tr>
            </thead>
            <tbody>
                {movies.map((movie) => (
                    <tr key={movie.id}>
                       <Link to={`/movies/${movie.id}`}> <td>{movie.title}</td></Link>
                        <td>{movie.release_date}</td>
                        <td>{movie.mpaa_rating}</td>
                    </tr>
                ))}
            </tbody>
        </table>
        </div>

    );
}

export default Movies;
