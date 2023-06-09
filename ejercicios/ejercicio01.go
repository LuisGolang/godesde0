package ejercicios

import "strconv"

func ConviertoANumero(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}

//cuando quiero hacer una funcion publica debo escribir la primera letra en mayuscula
// numero, _ se utiliza para darle tratamiento al error del return de la funcion

/*Manejo de errores:
if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
Podemos llamar el error como queramos, en este caso es "err", err.Error() me mostrara el nombre del error
*/
