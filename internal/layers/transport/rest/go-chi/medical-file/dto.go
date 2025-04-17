package medical_file

import (
	"github.com/go-andiamo/chioas"
)

type UploadedMedicalFile struct {
	Name string `json:"name" schema:"name"`
	Data []byte `json:"data" schema:"data"`
}

var UploadedMedicalFileOpenApiDefinition = chioas.Schema{
	Name:               "uploaded-medical-file",
	RequiredProperties: []string{"name", "data"},
	Properties: chioas.Properties{
		{
			Name:    "name",
			Type:    "string",
			Example: "Medical File Name",
		},
		{
			Name:   "data",
			Type:   "string",
			Format: "binary",
			Example: "iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAApgAAAKYB3X3/OAA" +
				"AABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAANCSURBVEiJtZZPbBtFFMZ/M7ubXdtdb1xSFyeilBapySVU8h8" +
				"OoFaooFSqiihIVIpQBKci6KEg9Q6H9kovIHoCIVQJJCKE1ENFjnAgcaSGC6rEnxBwA04Tx43t2FnvDAfjkNibxgHxnWb2e/u992b" +
				"ee7tCa00YFsffekFY+nUzFtjW0LrvjRXrCDIAaPLlW0nHL0SsZtVoaF98mLrx3pdhOqLtYPHChahZcYYO7KvPFxvRl5XPp1sN3ad" +
				"WiD1ZAqD6XYK1b/dvE5IWryTt2udLFedwc1+9kLp+vbbpoDh+6TklxBeAi9TL0taeWpdmZzQDry0AcO+jQ12RyohqqoYoo8RDwJr" +
				"U+qXkjWtfi8Xxt58BdQuwQs9qC/afLwCw8tnQbqYAPsgxE1S6F3EAIXux2oQFKm0ihMsOF71dHYx+f3NND68ghCu1YIoePPQN1pG" +
				"RABkJ6Bus96CutRZMydTl+TvuiRW1m3n0eDl0vRPcEysqdXn+jsQPsrHMquGeXEaY4Yk4wxWcY5V/9scqOMOVUFthatyTy8QyqwZ" +
				"+kDURKoMWxNKr2EeqVKcTNOajqKoBgOE28U4tdQl5p5bwCw7BWquaZSzAPlwjlithJtp3pTImSqQRrb2Z8PHGigD4RZuNX6JYj6w" +
				"j7O4TFLbCO/Mn/m8R+h6rYSUb3ekokRY6f/YukArN979jcW+V/S8g0eT/N3VN3kTqWbQ428m9/8k0P/1aIhF36PccEl6EhOcAUCr" +
				"XKZXXWS3XKd2vc/TRBG9O5ELC17MmWubD2nKhUKZa26Ba2+D3P+4/MNCFwg59oWVeYhkzgN/JDR8deKBoD7Y+ljEjGZ0sosXVTvb" +
				"c6RHirr2reNy1OXd6pJsQ+gqjk8VWFYmHrwBzW/n+uMPFiRwHB2I7ih8ciHFxIkd/3Omk5tCDV1t+2nNu5sxxpDFNx+huNhVT3/z" +
				"MDz8usXC3ddaHBj1GHj/As08fwTS7Kt1HBTmyN29vdwAw+/wbwLVOJ3uAD1wi/dUH7Qei66PfyuRj4Ik9is+hglfbkbfR3cnZm7c" +
				"hlUWLdwmprtCohX4HUtlOcQjLYCu+fzGJH2QRKvP3UNz8bWk1qMxjGTOMThZ3kvgLI5AzFfo379UAAAAASUVORK5CYII=",
		},
	},
}

type MedicalFileInfo struct {
	ID        string `json:"id"`
	Extension string `json:"extension"`
	Name      string `json:"name"`
}

var MedicalFileInfoOpenApiDefinition = chioas.Schema{
	Name:               "medical-file-info",
	RequiredProperties: []string{"id", "extension", "name"},
	Properties: chioas.Properties{
		{
			Name:    "id",
			Type:    "string",
			Example: "00000000-0000-0000-0000-000000000001",
		},
		{
			Name:    "extension",
			Type:    "string",
			Example: ".png",
		},
		{
			Name:    "name",
			Type:    "string",
			Example: "Medical File Name",
		},
	},
}
