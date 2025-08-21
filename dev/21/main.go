package main

import "fmt"

// ProductProvider - целевой интерфейс
type ProductProvider interface {
	GetProduct() Product
}

// Product - современный формат продукта
type Product struct {
	Name  string
	Price float64
}

// LegacyProduct - устаревший формат продукта
type LegacyProduct struct {
	ProductName string
	ProductCost float64
}

// LegacyProductSource - легаси-система
type LegacyProductSource struct{}

func (l *LegacyProductSource) FetchProduct() LegacyProduct {
	return LegacyProduct{
		ProductName: "Vintage Clock",
		ProductCost: 99.99,
	}
}

// ProductAdapter - адаптер для легаси-системы
type ProductAdapter struct {
	legacySource *LegacyProductSource
}

// NewProductAdapter - конструктор адаптера
func NewProductAdapter(legacySource *LegacyProductSource) *ProductAdapter {
	return &ProductAdapter{legacySource: legacySource}
}

// GetProduct - реализация целевого интерфейса
func (a *ProductAdapter) GetProduct() Product {
	legacyProduct := a.legacySource.FetchProduct()
	return a.adaptLegacyProduct(legacyProduct)
}

// adaptLegacyProduct - выделенная функция для преобразования данных
func (a *ProductAdapter) adaptLegacyProduct(legacyProduct LegacyProduct) Product {
	return Product{
		Name:  legacyProduct.ProductName,
		Price: legacyProduct.ProductCost,
	}
}

// ModernProductSource - современная система, уже реализующая целевой интерфейс
type ModernProductSource struct{}

func (m *ModernProductSource) GetProduct() Product {
	return Product{
		Name:  "Modern Gadget",
		Price: 149.99,
	}
}

func main() {
	fmt.Println("Adapter Pattern Example")

	// Использование современной системы
	modernSource := &ModernProductSource{}
	fmt.Println("Modern profuct:")
	printProduct(modernSource)

	// Использование легаси-системы через адаптер
	legacySource := &LegacyProductSource{}
	adapter := NewProductAdapter(legacySource)
	fmt.Println("Adapted legacy-product:")
	printProduct(adapter)
}

// printProduct - клиентская функция, работающая с любым ProductProvider
func printProduct(provider ProductProvider) {
	product := provider.GetProduct()
	fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
}
