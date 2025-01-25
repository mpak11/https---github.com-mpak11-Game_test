package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Входная точка программы
func main() {
	game := &Game{} // Создаем новый экземпляр Game
	game.Init()     // Инициализируем игру

	ebiten.SetWindowSize(1080, 920)                    // Устанавливаем размеры окна
	ebiten.SetWindowTitle("Моя первая игра на Ebiten") // Устанавливаем заголовок окна
	if err := ebiten.RunGame(game); err != nil {       // Запускаем игру
		log.Fatal(err) // Если произошла ошибка, выводим сообщение и завершаем
	}
}
