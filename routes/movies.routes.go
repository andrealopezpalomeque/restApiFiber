package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Movie struct {
	Title string `json:"title"`
	Id int `json:"id"`
	

}

func UseMoviesRoutes(router fiber.Router) {
	//slice de peliculas
	movies :=[]*Movie{
		{Title: "Movie 1", Id: 1},
		{Title: "Movie 2", Id: 2},
	}
	//todas las peliculas
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})

	//pelicula por id
	router.Get("/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		
		// for _, movie := range movies {
		// 	if movie.Id == id {
		// 		return c.JSON(movie)
		// 	}
		// }

		//return c.Status(404).SendString("No se encontro la pelicula")

		var movieFound Movie

		for _, movie := range movies {
			if movie.Id == id {
				movieFound = *movie
			}
		}

		return c.JSON(fiber.Map{
			"movie": movieFound,
		})

		
	})	


	//crear pelicula
	router.Post("/", func(c *fiber.Ctx) error {
	
		type Request struct {
			Title string `json:"title"`
			Id int `json:"id"`
		}

		var body Request

		c.BodyParser(&body) //parsea el body y lo guarda en body

		newMovie := Movie{
			Title: body.Title,
			Id: len(movies) + 1,
		}

		movies = append(movies, &newMovie)

		return c.JSON(fiber.Map{
			"message": "Pelicula creada correctamente",

	})

	})


}