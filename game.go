package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	counter int
	image   *ebiten.Image
}

func (g *Game) Init() {
	err := g.loadImage("images/adventur.png") // Укажите путь к изображению
	if err != nil {
		log.Fatal(err) // Обработка ошибки
	}
}

func (g *Game) loadImage(path string) error {
	file, err := os.Open(path) // Чтение файла изображения
	if err != nil {
		return err // Если произошла ошибка, верните её
	}
	defer file.Close() // Обязательно закрываем файл после использования

	img, err := png.Decode(file) // Декодирование изображения из файла
	if err != nil {
		return err
	}
	g.image = ebiten.NewImageFromImage(img) // Создание ebiten.Image
	return nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) { // Увеличиваем счетчик на 1
		g.counter++
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) { // Уменьшаем счетчик на 1
		g.counter--
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()                                                       // Чистим экран
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Счетчик: %d", g.counter)) // Отображение счетчика
	screen.DrawImage(g.image, nil)                                       // Отображение изображения
}

func (g *Game) Layout(int, int) (int, int) {
	return 1080, 920 // Устанавливаем размеры окна
}

//test 123
