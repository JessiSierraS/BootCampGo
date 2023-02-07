package hunt

import (
	"errors"
)

var (
	ErrTired = errors.New("cannot hunt, i am really tired")
	ErrNotHungry = errors.New("cannot hunt, i am not hungry")
	ErrSlowShark = errors.New("could not catch it")
	ErrPreyNil = errors.New("cannot hunt, prey nil")
)

// Shark, Estructura del Tiburón
type Shark struct {
	hungry bool		// Hambriento
	tired  bool		// Cansado
	speed  int		// Velocidad		
}

// Prey, Estructura de la Presa
type Prey struct {
	name  string	// Nombre
	speed int		// Velocidad
}

// Método Hunt(Caza) de la estructura Shark(Tiburón) con parámetro de entrada una estructura llamada Prey(Presa)
func (s *Shark) Hunt(p *Prey) (err error) {
	// Si la "Prey"(Presa) es Nil, retornará un mensaje con respuesta "Presa nil"
	if p == nil {
		return ErrPreyNil
	}
	//Si la presa es nula y todos los parámetros del tiburon se cumplen
	if s.tired && !s.hungry && s.speed >= p.speed {
		return nil
	}
	// Si el atributo "tired"(cansado) contiene un valor bool true, retornará un mensaje con respuesta "No puedo cazar, estoy muy cansado"
	if s.tired{
		// return fmt.Errorf("cannot hunt, i am really tired")
		return ErrTired
	}
	// Si el atributo "hungry"(hambriento) contiene un valor diferente bool true, es decir, es un bool false, retornará un mensaje con respuesta "No puedo cazar, no tengo hambre"
	if !s.hungry{
		// return fmt.Errorf("cannot hunt, i am not hungry")
		return ErrNotHungry
	}
	// Si la velocidad de "Prey"(Presa) es mayor o igual a la velocidad del "Shark"(Tiburón), retornará un mensaje con respuesta "No pude atraparlo"
	if p.speed >= s.speed {
		// El atributo "tired"(cansado) se le asignará un valor bool true
		s.tired = true
		// return fmt.Errorf("could not catch it")
		return ErrSlowShark
	}
	// Caso feliz, donde el Tiburón caza a la Presa. Se debe validar que una vez que cazó a la presa, el tiburón se llenó y se cansó.
	s.hungry = false
	s.tired = true
	return 
}

