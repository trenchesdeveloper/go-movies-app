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
        let moviesList: Movie[] = [
            { id: 1, title: "The Shawshank Redemption", release_date: '1986-03-07', runtime: 116, mpaa_rating: 'R', description: 'Arandom description' },
            { id: 2, title: "The Godfather", release_date: '1981-06-12', runtime: 116, mpaa_rating: 'R', description: 'Arandom description' },
        ];
        setMovies(moviesList);
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
