package storage

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/hacdan/cinema-api/types"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	store *sql.DB
}

func New(dbURL string) Storage {
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	return Storage{
		store: db,
	}
}

func (s *Storage) InitDB() error {
	createMoviesTableSQL := "CREATE TABLE IF NOT EXISTS movies (id TEXT, title TEXT, description TEXT, duration NUMERIC, genre TEXT);"
	createTheaterTableSQL := "CREATE TABLE IF NOT EXISTS theaters (id TEXT, name TEXT, location TEXT);"
	_, err := s.store.Exec(createMoviesTableSQL)
	if err != nil {
		return err
	}
	_, err = s.store.Exec(createTheaterTableSQL)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) AddMovie(m types.Movie) (types.Movie, error) {
	m.Id = uuid.NewString()
	_, err := s.store.Exec(
		"INSERT INTO movies (id, title, description, duration, genre) VALUES(?, ?, ?, ?, ?);",
		m.Id,
		m.Title,
		m.Description,
		m.Duration,
		m.Genre,
	)
	if err != nil {
		return types.Movie{}, err
	}
	return m, err
}

func (s *Storage) GetMovie(id string) (types.Movie, error) {
	row := s.store.QueryRow("SELECT id, title, description, duration, genre FROM movies WHERE id = ?;", id)
	movie := types.Movie{}

	err := row.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Duration)
	if err == sql.ErrNoRows {
		return types.Movie{}, err //TODO: Change to error about id not found
	}
	if err != nil {
		return types.Movie{}, err
	}
	return movie, nil
}

func (s *Storage) GetMovies() ([]types.Movie, error) {
	rows, err := s.store.Query("SELECT id, title, description, duration, genre FROM movies;")
	movies := []types.Movie{}
	if err == sql.ErrNoRows {
		return movies, err
	}
	if err != nil {
		return movies, err
	}
	for rows.Next() {
		movie := types.Movie{}
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Duration, &movie.Genre)
		if err != nil {
			return movies, err
		}
		movies = append(movies, movie)
	}

	return movies, err
}

func (s *Storage) UpdateMovie(m types.Movie) (types.Movie, error) {
	row, err := s.store.Query(
		"UPDATE movies SET title = ?, description = ?, duration = ?, genre = ? WHERE id = ?;",
		m.Title,
		m.Description,
		m.Duration,
		m.Genre,
		m.Id,
	)
	if err != nil {
		return types.Movie{}, err
	}
	updatedMovie := types.Movie{}
	err = row.Scan(&updatedMovie.Id, &updatedMovie.Title, &updatedMovie.Description, &updatedMovie.Duration, &updatedMovie.Genre)
	if err != nil {
		return updatedMovie, err
	}
	return updatedMovie, err
}

func (s *Storage) DeleteMovie(id string) error {
	_, err := s.store.Exec("DELETE FROM movies WHERE id = ?;", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) AddTheater(theater types.Theater) (types.Theater, error) {
	theater.Id = uuid.NewString()
	_, err := s.store.Exec(
		"INSERT INTO theaters (id, name, location) VALUES(?, ?, ?);",
		theater.Id,
		theater.Name,
		theater.Location,
	)
	if err != nil {
		return types.Theater{}, err
	}
	return theater, err
}

func (s *Storage) GetTheater(id string) (types.Theater, error) {
	row := s.store.QueryRow("SELECT name, location FROM theaters WHERE id = ?;", id)
	theater := types.Theater{}

	err := row.Scan(&theater.Id, &theater.Name, &theater.Location)
	if err == sql.ErrNoRows {
		return types.Theater{}, err //TODO: Change to error about id not found
	}
	if err != nil {
		return types.Theater{}, err
	}
	return theater, nil
}

func (s *Storage) GetTheaters() ([]types.Theater, error) {
	rows, err := s.store.Query("SELECT id, name, location FROM theaters;")
	theaters := []types.Theater{}
	if err == sql.ErrNoRows {
		return theaters, err
	}
	if err != nil {
		return theaters, err
	}
	for rows.Next() {
		theater := types.Theater{}
		err = rows.Scan(&theater.Id, &theater.Name, &theater.Location)
		if err != nil {
			return theaters, err
		}
		theaters = append(theaters, theater)
	}

	return theaters, err
}

/*
func (s *Storage) GetShowTimes(id string) ([]types.Showtime, error) {
}

func (s *Storage) AddReservation(r types.MovieReservation) (types.MovieReservation, error) {
}

func (s *Storage) ConfirmReservation(id string) (types.MovieReservation, error) {
}

func (s *Storage) GetReservation(id string) (types.MovieReservation, error) {
}*/
