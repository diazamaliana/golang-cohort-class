package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Product struct {
    ID        int
    Name      string
    CreatedAt string
    UpdatedAt string
	Variants  []Variant
}

type Variant struct {
    ID         int
    VariantName string
    Quantity    int
    ProductID   int
    CreatedAt   string
    UpdatedAt   string
}


var (
	db   *sql.DB
	err  error
	rows *sql.Rows
)

func main() {
    db, err := NewDB()
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
    }
    defer db.Close()

	fmt.Println("Successfully connected to database")

    // Penggunaan fungsi-fungsi CRUD

    // 1. Membuat produk baru
    // productID, err := createProduct("Product 5")
    // if err != nil {
    //     fmt.Println("Error creating product:", err)
    // }
    // fmt.Printf("Produk dengan ID: %d telah dibuat\n", productID)

    // 2. Mendapatkan produk berdasarkan ID
    // product, err := getProductById(20)
    // if err != nil {
    //     fmt.Println("Error retrieving product:", err)
    // }
    // fmt.Printf("Produk: %+v\n", product)

    // 3. Mengupdate produk yang sudah ada
    // err = updateProduct(18, "Product 5 Updated")
    // if err != nil {
    //     fmt.Println("Error updating product:", err)
    // }
    // fmt.Printf("Produk dengan ID: %d telah diperbarui\n", 18)

    // 4. Membuat varian baru untuk produk
    // variantID, err := createVariant(21, "Variant A", 20)
    // if err != nil {
    //     fmt.Println("Error creating variant:", err)
    // }
    // fmt.Printf("Variasi dengan ID: %d telah dibuat\n", variantID)

    // 5. Mengupdate varian berdasarkan ID
    // err = updateVariantById(1, "Variant A Updated", 60)
    // if err != nil {
    //     fmt.Println("Error updating variant:", err)
    // }
    // fmt.Printf("Variasi dengan ID: %d telah diperbarui\n", 1)

    // 6. Mendapatkan produk dengan semua varian yang terkait
    // productWithVariants, err := getProductWithVariant(20)
    // if err != nil {
    //     fmt.Println("Error retrieving product with variants:", err)
    // }
    // fmt.Printf("Produk dengan Variasi: %+v\n", productWithVariants)

	// 7. Menghapus varian berdasarkan ID
    // err = deleteVariantById(3)
    // if err != nil {
    //     fmt.Println("Error deleting variant:", err)
    // }
    // fmt.Printf("Variasi dengan ID: %d telah dihapus\n", 3)
}


func NewDB() (*sql.DB, error) {
    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Access environment variables - using MYSQL
    // Rename .env.example to .env. Enter your configuration to your .env file
    user:= os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	db, err = sql.Open("mysql", mysqlInfo)
    if err != nil {
        return nil, err
    }
    return db, nil
}

func createProduct(name string) (int, error) {
    query := "INSERT INTO products (name, created_at, updated_at) VALUES (?, NOW(), NOW())"
    result, err := db.Exec(query, name)
    if err != nil {
        return 0, err
    }

    id, _ := result.LastInsertId()
    return int(id), nil
}

func updateProduct(id int, name string) error {
    query := "UPDATE products SET name = ?, updated_at = NOW() WHERE id = ?"
    _, err := db.Exec(query, name, id)
    return err
}

func getProductById(id int) (Product, error) {
    var product Product
    query := "SELECT id, name, created_at, updated_at FROM products WHERE id = ?"
    err := db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
    if err != nil {
        return Product{}, err
    }

    return product, nil
}

func updateVariantById(id int, name string, quantity int) error {
    _, err := db.Exec("UPDATE variants SET variant_name = ?, quantity = ?, updated_at = NOW() WHERE id = ?", name, quantity, id)
    return err
}

func createVariant(productID int, name string, quantity int) (int, error) {
    result, err := db.Exec("INSERT INTO variants (product_id, variant_name, quantity, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())", productID, name, quantity)
    if err != nil {
        return 0, err
    }

    id, _ := result.LastInsertId()
    return int(id), nil
}

func getProductWithVariant(productID int) (Product, error) {
    product, err := getProductById(productID)
    if err != nil {
        return Product{}, err
    }

    query := "SELECT id, variant_name, quantity, product_id, created_at, updated_at FROM variants WHERE product_id = ?"
    rows, err := db.Query(query, productID)
    if err != nil {
        return Product{}, err
    }
    defer rows.Close()

    var variants []Variant
    for rows.Next() {
        var variant Variant
        err := rows.Scan(&variant.ID, &variant.VariantName, &variant.Quantity, &variant.ProductID, &variant.CreatedAt, &variant.UpdatedAt)
        if err != nil {
            return Product{}, err
        }
        variants = append(variants, variant)
    }
    product.Variants = variants

    return product, nil
}

func deleteVariantById(id int) error {
    _, err := db.Exec("DELETE FROM variants WHERE id = ?", id)
    return err
}
